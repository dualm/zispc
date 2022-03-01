package zispc

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	AlphabetCount     = 26
	AsciiCapital      = 65
	NanoToMicrosecond = 1000
)

func getRandStr(l int) string {
	r := rand.New(rand.NewSource(time.Now().UnixMicro()))

	bytes := make([]byte, l)

	for i := 0; i < l; i++ {
		b := r.Intn(AlphabetCount) + AsciiCapital
		bytes[i] = byte(b)
	}

	return string(bytes)
}

func getTransactionId() string {
	t := time.Now()
	micro := t.Nanosecond() / NanoToMicrosecond

	return fmt.Sprintf("%s%06d", t.Format("20060102150405"), micro)
}
