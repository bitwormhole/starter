package util

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCurrentTime(t *testing.T) {
	const step = 17
	for timeout := 2000; timeout > 0; timeout -= step {
		time.Sleep(step * time.Millisecond)
		n1 := CurrentTimestamp()
		t1 := Int64ToTime(n1)
		str := strconv.FormatInt(n1, 10)
		os.Stdout.WriteString(str + ": " + t1.String() + "\n")
	}
}

func TestInt64ToTime(t *testing.T) {

	const step = 17
	t0 := time.Now()
	n0 := TimeToInt64(t0)
	end := n0 + 2000

	for n := n0; n < end; n += step {
		n1 := n
		t1 := Int64ToTime(n1)
		n2 := TimeToInt64(t1)
		if n1 != n2 {
			msg := strings.Builder{}
			msg.WriteString("n1!=n2:")
			msg.WriteString("  n1=" + strconv.FormatInt(n1, 10))
			msg.WriteString(", n2=" + strconv.FormatInt(n2, 10))
			msg.WriteString(", time=" + t1.String())
			t.Error(msg.String())
		}

		// os.Stdout.WriteString("TestInt64ToTime: " + t1.String() + "\n")
	}
}

func TestTimeStamp(t *testing.T) {

	t1 := Now()

	tt := t1.GetTime()
	i1 := t1.Int64()
	s1 := t1.String()

	t.Log("tt = ", tt)
	t.Log("i1 = ", i1)
	t.Log("s1 = ", s1)
}

func BenchmarkNow(b *testing.B) {
	b.Log("begin")
	i := b.N
	list := []Time{}
	for ; i > 0; i-- {
		ts := Now()
		list = append(list, ts)
	}
	b.Log("done")
}
