package main

import "fmt"

type OptFunc func(*Opts)
type Opts struct {
	maxConn int
	id      string
	tls     bool
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func WithTLS(opts *Opts) {
	opts.tls = true
}

type Server struct {
	Opts
}

func newServer(opts ...OptFunc) *Server {
	o := defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}
	return &Server{
		Opts: o,
	}
}

func withMaxConn(n int) OptFunc {
	return func(o *Opts) {
		o.maxConn = n
	}
}

func withID(s string) OptFunc {
	return func(o *Opts) {
		o.id = s
	}
}

func main() {
	o := newServer(WithTLS, withMaxConn(99),withID("Naveen"))
	fmt.Printf("%+v", o)
}
