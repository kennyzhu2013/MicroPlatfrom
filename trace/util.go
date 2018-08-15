package trace

import (
	"time"
)

func SpanFromHeader(md map[string]string) (*Span, bool) {
	var debug bool
	if md[DebugHeader] == "1" {
		debug = true
	}

	// can we get span header and trace header?
	if len(md[SpanHeader]) == 0 && len(md[TraceHeader]) == 0 {
		return nil, false
	}

	return &Span{
		Id:        md[SpanHeader],
		TraceId:   md[TraceHeader],
		ParentId:  md[ParentHeader],
		Debug:     debug,
		Timestamp: time.Now(),
	}, true
}

func HeaderWithSpan(md map[string]string, s *Span) map[string]string {
	debug := "0"
	if s.Debug {
		debug = "1"
	}

	md[SpanHeader] = s.Id
	md[TraceHeader] = s.TraceId
	md[ParentHeader] = s.ParentId
	md[DebugHeader] = debug
	return md
}
