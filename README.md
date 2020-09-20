# go-testwire
An easy-to-use API to add test automation probes.

## Introduction
This library provides an easy-to-use API to add test probes to your code. These
probes can be used to retrieve information of your application that is normally
not exposed.

This can be used to assert certain conditions in your applications that would
otherwise be hard to assert while performing automated tests (black-box tests).

## Use case
When testing an API, you can perform a request, and maybe another one to assert
a condition. But if you are working with hardware, it might be valueable to
assert that an API request perfoms a certain hardware action on the lower
level.

A test scenario could look like this:

```gherkin
Given a running system
When I start the motor via the API
Then the response is successful
And the hardware turns on the power of the motor
```

This library may be useful to provide an implementation of the last step: it is
information that is typically not exposed via an API. By probing into the
communication layer with the hardware, we can now use this library to observe
the 'internals'.

## Usage
You need to explicitly enable TestWire support by compiling your application
with the `testwire` tag set, e.g. `go build -tags testwire main.go`.

Without this tag, a 'disabled' version of testwire is compiled that provides
stubs.

### Adding probes and signals
Probes retrieve information, signals perform actions. They both assume that
data can be returned as a string.

To add a probe, add the following:

```go
testwire.AddProbe("probe.name", func() string {
    return "Foo"
})
```

To add a signal, add the following:

```go
testwire.AddSignal("signal.name", func(s string) {
    log.Printf("Signal data: %s", s)
})
```

### Serving
To start serving, simply add:

```go
go testwire.Serve(port number)

```

### API
Probes are exposed on `http://localhost:9000/probes/probe.name`. A GET request
returns the probe value. Signals can be invoked by performing a POST request to
`http://localhost:9000/signals/signal.name`.

## License
See the [`LICENSE.md`](LICENSE.md) file (MIT license).