package cli

import "testing"

func TestChainBuilder(t *testing.T) {

	builder := &filterChainBuilder{}

	builder.add(4, &NopFilter{})
	builder.add(2, &ExecutorFilter{})
	builder.add(5, &ContextFilter{})
	builder.add(3, &HandlerFinderFilter{})

	builder.reverse = false
	chain := builder.create()

	t.Log(chain)
}
