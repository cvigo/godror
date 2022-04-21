// Copyright 2017, 2021 The Godror Authors
//
//
// SPDX-License-Identifier: UPL-1.0 OR Apache-2.0

package godror

import (
	"context"
	"io"
	"runtime"
	"sync"
	"time"

	"github.com/go-logfmt/logfmt"
	"github.com/go-logr/logr"
)

var globalTracer = &swapTracer{}
var start_txt = "start_time"
var end_txt = "end_time"
var duration_txt = "duration"

func SetTracer(tracer logr.Logger) {
	globalTracer.Swap(tracer)
}
func SetTrace(f func(...interface{}) error) {
	globalTracer.Swap(logr.New(TraceFunc(f)))
}

func SetTraceText(start, end, duration string) {
	if start != "" {
		start_txt = start
	}
	if end != "" {
		end_txt = end
	}
	if duration != "" {
		duration_txt = duration
	}
}

// NewTracefmtTracer returns a Tracer, and that logs using logfmt, to the given io.Writer.
func NewTracefmtTracer(w io.Writer) logr.Logger {
	enc := logfmt.NewEncoder(w)
	return logr.New(TraceFunc(func(keyvals ...interface{}) error {
		firstErr := enc.EncodeKeyvals(keyvals...)
		if err := enc.EndRecord(); err != nil && firstErr == nil {
			return err
		}
		return firstErr
	}))
}

type Tracer interface {
	Trace(start time.Time, end time.Time, keyvals ...interface{}) error
}
type TraceFunc func(keyvals ...interface{}) error

func (f TraceFunc) Trace(start time.Time, end time.Time, keyvals ...interface{}) error {
	return f(append(append(make([]interface{}, 0, 4+len(keyvals)), start, end), keyvals...))
}
func (f TraceFunc) Init(_ logr.RuntimeInfo) {}
func (f TraceFunc) Enabled(_ int) bool      { return f != nil }
func (f TraceFunc) Info(level int, msg string, keyvals ...interface{}) {
	f(append(append(make([]interface{}, 0, 4+len(keyvals)), "lvl", level, "msg", msg), keyvals...)...)
}
func (f TraceFunc) Error(err error, msg string, keyvals ...interface{}) {
	f(append(append(make([]interface{}, 0, 4+len(keyvals)), "msg", msg, "error", err), keyvals...)...)
}
func (f TraceFunc) WithValues(plusKeyVals ...interface{}) logr.LogSink {
	return TraceFunc(func(keyvals ...interface{}) error {
		return f(append(plusKeyVals, keyvals...)...)
	})
}
func (f TraceFunc) WithName(name string) logr.LogSink {
	return f.WithValues("name", name)
}

type swapTracer struct {
	mu     sync.RWMutex
	tracer logr.Logger
	isSet  bool
}

func (sw *swapTracer) IsSet() bool {
	if sw == nil {
		return false
	}
	sw.mu.RLock()
	isset := sw.isSet
	sw.mu.RUnlock()
	return isset
}

func (sw *swapTracer) Swap(tracer logr.Logger) {
	sw.mu.Lock()
	sw.tracer = tracer
	sw.isSet = true
	sw.mu.Unlock()
}

//func (sw *swapTracer) Info(msg string, keyvals ...interface{}) {
//	sw.tracer.Info(msg, keyvals...)
//}
//func (sw *swapTracer) Error(err error, msg string, keyvals ...interface{}) {
//	sw.tracer.Error(err, msg, keyvals...)
//}

func (sw *swapTracer) Trace(start time.Time, end time.Time, keyvals ...interface{}) error {
	sw.mu.RLock()
	tracer := sw.tracer
	sw.mu.RUnlock()
	if !tracer.Enabled() {
		return nil
	}
	runtime.NumCgoCall()
	if len(keyvals) >= 2 {
		if msgKey, ok := keyvals[0].(string); ok && msgKey == "msg" {
			if msg, ok := keyvals[1].(string); ok {
				tracer.Info(msg, append(append(make([]interface{}, 0, 4+len(keyvals)), start_txt, start, end_txt, end, duration_txt, end.Sub(start)), keyvals[2:]...)...)
				return nil
			}
		}
	}
	tracer.Info("", append(append(make([]interface{}, 0, 6+len(keyvals)), start_txt, start, end_txt, end, duration_txt, end.Sub(start)), keyvals...)...)
	return nil
}

type traceCtxKey struct{}

func getTracer() Tracer {
	if globalTracer.IsSet() {
		return globalTracer
	}
	return nil
}
func ctxGetTrace(ctx context.Context) Tracer {
	if ctx != nil {
		if lgr, ok := ctx.Value(traceCtxKey{}).(Tracer); ok {
			return lgr
		}
	}
	return getTracer()
}

// ContextWithTracer returns a context with the given tracer.
func ContextWithTracer(ctx context.Context, tracer Tracer) context.Context {
	return context.WithValue(ctx, traceCtxKey{}, tracer)
}

// ContextWithTrace returns a context with the given log function.
func ContextWithTrace(ctx context.Context, logF func(...interface{}) error) context.Context {
	return context.WithValue(ctx, traceCtxKey{}, TraceFunc(logF))
}

// ExecInTrace runs the passed function and traces the duration.
func ExecInTrace(f TraceRunner, properties ...interface{}) error {
	// Start new span
	tracer := getTracer()
	start := time.Now()
	err := f()
	if tracer != nil {
		_ = tracer.Trace(start, time.Now(), properties...)
	}
	return err
}

type TraceRunner func() error
