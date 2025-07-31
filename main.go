package main

import (
	"github.com/buchgr/bazel-remote/v2/run"
)

func main() {
	run.WithMaybeProxyOverride(nil)
}
