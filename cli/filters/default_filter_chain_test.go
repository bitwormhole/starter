package filters

import (
	"testing"

	"github.com/bitwormhole/starter/cli"
)

func TestChainBuilder(t *testing.T) {

	builder := &cli.FilterChainBuilder{}

	builder.Add(4, &NopFilter{})

	builder.Add(2, &ExecutorFilter{})
	builder.Add(5, &ContextFilter{})
	builder.Add(3, &HandlerFinderFilter{})

	const reverse = false
	chain := builder.Create(reverse)

	t.Log(chain)
}
