package zispc

import (
	"fmt"
)

var withS bool

// SetWithS, XML sitename为"S00x"的格式
func SetWithS() {
	withS = true
}

// UnsetWIthS, XML sitename为"00x"的格式
func UnsetWithS() {
	withS = false
}

func OnlySiteName() string {
	if withS {
		return "S01"
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

// GetXMLSiteName, 生成XML sitename
func GetXMLSiteName(n int) string {
	if withS {
		return fmt.Sprintf("S%02d", n)
	}

	return fmt.Sprintf("%03d", n)
}

// MakeXMLSites，由字符串切片转换为字符串map
func MakeXMLSites(sites []string) map[string]string {
	m := make(map[string]string)

	for i := range sites {
		m[GetXMLSiteName(i+1)] = sites[i]
	}

	return m
}
