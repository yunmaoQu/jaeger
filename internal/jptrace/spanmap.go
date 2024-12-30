// Copyright (c) 2024 The Jaeger Authors.
// SPDX-License-Identifier: Apache-2.0

package jptrace

import "go.opentelemetry.io/collector/pdata/ptrace"

// SpanMap iterates over all spans in the provided ptrace.Traces and maps each span
// to a key generated by the provided keyFn function. The resulting map has keys of type K
// and values of type ptrace.Span.
func SpanMap[K comparable](traces ptrace.Traces, keyFn func(ptrace.Span) K) map[K]ptrace.Span {
	spanMap := make(map[K]ptrace.Span)
	for i := 0; i < traces.ResourceSpans().Len(); i++ {
		resource := traces.ResourceSpans().At(i)
		for j := 0; j < resource.ScopeSpans().Len(); j++ {
			scope := resource.ScopeSpans().At(j)
			for k := 0; k < scope.Spans().Len(); k++ {
				span := scope.Spans().At(k)
				spanMap[keyFn(span)] = span
			}
		}
	}
	return spanMap
}