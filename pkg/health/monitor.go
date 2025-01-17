package health

import (
	"context"
	"sync"

	corev1 "github.com/rancher/opni/pkg/apis/core/v1"
	"github.com/rancher/opni/pkg/health/annotations"
	"github.com/rancher/opni/pkg/util"
	"go.uber.org/zap"
)

type Monitor struct {
	MonitorOptions
	mu              sync.Mutex
	currentHealth   map[string]*corev1.Health
	currentStatus   map[string]*corev1.Status
	healthListeners map[string][]chan *corev1.Health
	statusListeners map[string][]chan *corev1.Status
}

type MonitorOptions struct {
	lg *zap.SugaredLogger
}

type MonitorOption func(*MonitorOptions)

func (o *MonitorOptions) apply(opts ...MonitorOption) {
	for _, op := range opts {
		op(o)
	}
}

func WithLogger(lg *zap.SugaredLogger) MonitorOption {
	return func(o *MonitorOptions) {
		o.lg = lg
	}
}

func NewMonitor(opts ...MonitorOption) *Monitor {
	options := MonitorOptions{
		lg: zap.NewNop().Sugar(),
	}
	options.apply(opts...)

	return &Monitor{
		MonitorOptions:  options,
		currentHealth:   make(map[string]*corev1.Health),
		currentStatus:   make(map[string]*corev1.Status),
		healthListeners: make(map[string][]chan *corev1.Health),
		statusListeners: make(map[string][]chan *corev1.Status),
	}
}

func (m *Monitor) Run(ctx context.Context, updater HealthStatusUpdater) {
	defer func() {
		m.mu.Lock()
		defer m.mu.Unlock()
		m.currentHealth = make(map[string]*corev1.Health)
		m.currentStatus = make(map[string]*corev1.Status)
		for _, listeners := range m.healthListeners {
			for _, c := range listeners {
				close(c)
			}
		}
		for _, listeners := range m.statusListeners {
			for _, c := range listeners {
				close(c)
			}
		}
		m.healthListeners = make(map[string][]chan *corev1.Health)
		m.statusListeners = make(map[string][]chan *corev1.Status)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case update, ok := <-updater.HealthC():
			if !ok {
				m.lg.Debug("health update channel closed")
				return
			}
			m.mu.Lock()
			m.lg.With(
				"id", update.ID,
				"ready", update.Health.Ready,
				"conditions", update.Health.Conditions,
			).With(
				annotations.KeyValuePairs(update.Health.GetAnnotations())...,
			).Info("received health update")
			m.currentHealth[update.ID] = update.Health
			for _, ch := range m.healthListeners[update.ID] {
				ch <- util.ProtoClone(update.Health)
			}
			m.mu.Unlock()
		case update, ok := <-updater.StatusC():
			if !ok {
				m.lg.Debug("status update channel closed")
				return
			}
			m.mu.Lock()
			m.lg.With(
				"id", update.ID,
				"connected", update.Status.Connected,
			).Info("received status update")
			m.currentStatus[update.ID] = update.Status
			for _, ch := range m.statusListeners[update.ID] {
				ch <- util.ProtoClone(update.Status)
			}
			m.mu.Unlock()
		}
	}
}

func (m *Monitor) GetHealthStatus(id string) *corev1.HealthStatus {
	m.mu.Lock()
	defer m.mu.Unlock()
	return &corev1.HealthStatus{
		Health: util.ProtoClone(m.currentHealth[id]),
		Status: util.ProtoClone(m.currentStatus[id]),
	}
}

func (m *Monitor) WatchHealthStatus(ctx context.Context, id string) <-chan *corev1.HealthStatus {
	m.mu.Lock()
	ch := make(chan *corev1.HealthStatus, 10)
	hl := make(chan *corev1.Health, 10)
	sl := make(chan *corev1.Status, 10)
	m.healthListeners[id] = append(m.healthListeners[id], hl)
	m.statusListeners[id] = append(m.statusListeners[id], sl)
	curHealth := util.ProtoClone(m.currentHealth[id])
	curStatus := util.ProtoClone(m.currentStatus[id])

	ch <- &corev1.HealthStatus{
		Health: util.ProtoClone(curHealth),
		Status: util.ProtoClone(curStatus),
	}
	m.mu.Unlock()

	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case health, ok := <-hl:
				if !ok {
					break LOOP
				}
				curHealth = health
			case status, ok := <-sl:
				if !ok {
					break LOOP
				}
				curStatus = status
			}
			hs := &corev1.HealthStatus{
				Health: util.ProtoClone(curHealth),
				Status: util.ProtoClone(curStatus),
			}
			select {
			case <-ctx.Done():
			case ch <- hs:
			}
		}

		m.mu.Lock()
		defer m.mu.Unlock()
		for i, h := range m.healthListeners[id] {
			if h == hl {
				m.healthListeners[id] = append(m.healthListeners[id][:i], m.healthListeners[id][i+1:]...)
				close(hl)
				break
			}
		}
		for i, s := range m.statusListeners[id] {
			if s == sl {
				m.statusListeners[id] = append(m.statusListeners[id][:i], m.statusListeners[id][i+1:]...)
				close(sl)
				break
			}
		}
	}()

	return ch
}
