package tracing

var (
	_ Tracer = NopTracer{}
	_ Span   = nopSpan{}
)

type NopTracer struct{}

func (n NopTracer) StartSpan(_ string) Span {
	return nopSpan{}
}

func (n NopTracer) StartSpanFromTraceParent(_, _, _ string) Span {
	return nopSpan{}
}

type nopSpan struct{}

func (n nopSpan) SetBaggageItem(_ string, _ any) {
}

func (n nopSpan) Finish() {
}

func (s nopSpan) TraceID() string {
	return ""
}

func (s nopSpan) SpanID() string {
	return ""
}
