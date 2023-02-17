package zispc

import (
	"fmt"
)

var withS bool

// SetWithS XML site name为"S00x"的格式
func SetWithS() {
	withS = true
}

// UnsetWithS XML site name为"000x"的格式
func UnsetWithS() {
	withS = false
}

// OnlySiteName XML site name为"S001"的格式
func OnlySiteName() string {
	if withS {
		return "S001"
	}

	return "001"
}

func checkSiteName(k string) string {
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

// GetXMLSiteName 生成XML sitename
func GetXMLSiteName(n int) string {
	if withS {
		return fmt.Sprintf("S%03d", n)
	}

	return fmt.Sprintf("%03d", n)
}

// MakeXMLSites 由字符串切片转换为字符串map
func MakeXMLSites(sites []string) map[string]string {
	m := make(map[string]string)

	for i := range sites {
		m[GetXMLSiteName(i+1)] = sites[i]
	}

	return m
}
