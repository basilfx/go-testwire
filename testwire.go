package testwire

type ProbeFunc func() string
type SignalFunc func(string)

type TestWire interface {
	IsEnabled() bool

	AddProbe(name string, callback ProbeFunc)
	AddSignal(name string, callback SignalFunc)

	Serve(port uint)
}

var instance TestWire

func init() {
	instance = New()
}

// IsEnabled returns true if TestWire is enabled.
func IsEnabled() bool {
	return instance.IsEnabled()
}

// AddProbe adds a probe.
func AddProbe(name string, callback ProbeFunc) {
	instance.AddProbe(name, callback)
}

// AddSignal adds a signal.
func AddSignal(name string, callback SignalFunc) {
	instance.AddSignal(name, callback)
}

// Serve starts the test interface.
func Serve(port uint) {
	instance.Serve(port)
}
