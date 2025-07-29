package main

import (
	_ "net/http/pprof" // Register pprof handlers with DefaultServeMux.

	"github.com/buchgr/bazel-remote/v2/run"
)

func main() {
	run.WithMaybeProxyOverride(nil)
}
