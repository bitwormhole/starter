package srctestgo

import (
	"fmt"
	"os"
	"testing"
)

func TestEnv(t *testing.T) {

	all := os.Environ()
	fmt.Println("os.Environ():")
	for _, str := range all {
		fmt.Println("\t", str)
	}

}
