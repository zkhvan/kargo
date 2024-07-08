package pprof

import (
	"net/http"
	"net/http/pprof"
	"runtime"
)

// HTTPPrefixPProf is the prefix appended to all Endpoints.
const HTTPPrefixPProf = "/debug/pprof"

// Endpoints defines the debugging endpoints that are added by SetupHandlers.
var Endpoints = map[string]http.Handler{
	HTTPPrefixPProf + "/":             http.HandlerFunc(pprof.Index),
	HTTPPrefixPProf + "/cmdline":      http.HandlerFunc(pprof.Cmdline),
	HTTPPrefixPProf + "/profile":      http.HandlerFunc(pprof.Profile),
	HTTPPrefixPProf + "/symbol":       http.HandlerFunc(pprof.Symbol),
	HTTPPrefixPProf + "/trace":        http.HandlerFunc(pprof.Trace),
	HTTPPrefixPProf + "/heap":         pprof.Handler("heap"),
	HTTPPrefixPProf + "/goroutine":    pprof.Handler("goroutine"),
	HTTPPrefixPProf + "/threadcreate": pprof.Handler("threadcreate"),
	HTTPPrefixPProf + "/block":        pprof.Handler("block"),
	HTTPPrefixPProf + "/mutex":        pprof.Handler("mutex"),
}

// GetHandlers returns the pprof endpoints for the mgr metrics server.
//
// The func can be used in the main.go file of your controller, when initializing the manager:
//
//	func main() {
//		mgrConfig := ctrl.Options{
//			Metrics: metricsserver.Options{
//				BindAddress:   metricsAddr,
//				ExtraHandlers: pprof.GetHandlers(),
//			},
//		}
//	}
func GetHandlers() map[string]http.Handler {
	// Only set the fraction if there is no existing setting
	if runtime.SetMutexProfileFraction(-1) == 0 {
		// Default to report 1 out of 5 mutex events, on average
		runtime.SetMutexProfileFraction(5)
	}

	return Endpoints
}
