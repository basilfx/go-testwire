// +build !testwire

package testwire

type TestWireDisabled struct {
}

func New() TestWire {
	return &TestWireDisabled{}
}

func (t *TestWireDisabled) IsEnabled() bool {
	return false
}

func (t *TestWireDisabled) AddProbe(name string, callback ProbeFunc) {
	// NOP
}

func (t *TestWireDisabled) AddSignal(name string, callback SignalFunc) {
	// NOP
}

func (t *TestWireDisabled) Serve(port uint) {
	// NOP
}
