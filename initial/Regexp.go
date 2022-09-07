package initial

import (
	"regexp"
	"strings"
)

func Check_for_banner(s string) bool {
	c, _ := regexp.MatchString(`([a-z]+)=([a-z])`, s)
	if !c {
		return false
	}
	f := strings.SplitAfter(s, "=")
	matched, _ := regexp.MatchString(`--output`, f[0])
	if !matched {
		return false
	}
	e, _ := regexp.MatchString(`[.]`, s)
	if !e {
		return false
	}
	h := strings.SplitAfter(f[1], ".")
	if h[1] == "txt" {
		return matched
	}
	return false
}
