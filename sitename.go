package zispc

import (
	"fmt"
)

var withS bool

func SetWithS() {
	withS = true
}

func UnsetWithS() {
	withS = false
}

func OnlySiteName() string {
	if withS {
		return "S01"
	}

	return "001"
}

func checkSitename(k string) string {
	if len(k) < 3 {
		setError(fmt.Errorf("wrong key length, Key : %s", k))

		return ""
	}

	if withS && k[0] != 'S' {
		setError(fmt.Errorf(`key not start with "S", Key: %s`, k))

		return ""
	}

	return k
}

func GetSitename(n int) string {
	if withS {
		return fmt.Sprintf("S%02d", n)
	}

	return fmt.Sprintf("%03d", n)
}
