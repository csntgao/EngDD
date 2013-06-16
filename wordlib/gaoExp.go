// gaoExp
package wordlib

import (
	"regexp"
)

type GaoRegexp struct {
	*regexp.Regexp
}

// add a new method to our new regular expression type
func (r *GaoRegexp) FindStringSubmatchMap(s string) [](map[string]string) {
	captures := make([](map[string]string), 0)
	matches := r.FindAllStringSubmatch(s, -1)
	if matches == nil {
		return captures
	}
	names := r.SubexpNames()
	for _, match := range matches {
		cmap := make(map[string]string)
		for pos, val := range match {
			name := names[pos]
			if name == "" {
				continue
			}

			cmap[name] = val
		}
		captures = append(captures, cmap)
	}
	return captures
}
