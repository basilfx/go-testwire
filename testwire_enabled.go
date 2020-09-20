// +build testwire

package testwire

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type TestWireEnabled struct {
	lock    sync.RWMutex
	probes  map[string]ProbeFunc
	signals map[string]SignalFunc
}

func New() TestWire {
	return &TestWireEnabled{
		probes:  make(map[string]ProbeFunc),
		signals: make(map[string]SignalFunc),
	}
}

func (t *TestWireEnabled) IsEnabled() bool {
	return true
}

func (t *TestWireEnabled) AddProbe(name string, callback ProbeFunc) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.probes[name] = callback
}

func (t *TestWireEnabled) AddSignal(name string, callback SignalFunc) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.signals[name] = callback
}

func (t *TestWireEnabled) handleProbe(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	// Check if probe exists.
	t.lock.RLock()
	probe, ok := t.probes[name]
	t.lock.RUnlock()

	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Invoke probe.
	raw := []byte(probe())

	// Write data to response.
	w.Header().Set("Content-Type", "text/plain")
	w.Write(raw)
}

func (t *TestWireEnabled) handleSignal(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	// Check if probe exists.
	t.lock.RLock()
	signal, ok := t.signals[name]
	t.lock.RUnlock()

	if !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// Read request body.
	raw, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// Signal probe.
	signal(string(raw))
}

func (t *TestWireEnabled) Serve(port uint) {
	r := mux.NewRouter()

	r.HandleFunc("/probes/{name}", t.handleProbe).Methods("GET")
	r.HandleFunc("/signals/{name}", t.handleSignal).Methods("POST")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
