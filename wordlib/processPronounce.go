// processPronounce
package wordlib

import (
	"fmt"
	"regexp"
)

var wordPronounceABExp = GaoRegexp{regexp.MustCompile(`英[\s\S]+"EN-US">(?P<BPronounce>.+)</strong><strong>]</strong>[\s\S]+美[\s\S]+"EN-US">(?P<APronounce>.+)</strong><strong>]</strong>`)}
var wordPronounceExp = GaoRegexp{regexp.MustCompile(`"EN-US">(?P<Pronounce>.+)</strong><strong>]</strong>`)}

func parserPronounce(text string) (string, string) {
	mmap := wordPronounceABExp.FindStringSubmatchMap(text)
	a, b := "", ""
	for _, ww := range mmap {
		for k, v := range ww {
			if k == "APronounce" {
				a = v
			}
			if k == "BPronounce" {
				b = v
			}

		}

	}
	if (a == "") && (b == "") {
		mmap := wordPronounceExp.FindStringSubmatchMap(text)
		for _, ww := range mmap {
			for k, v := range ww {
				if k == "Pronounce" {
					a = v
					b = v
				}

			}

		}
	}
	return a, b

}

func PeekPronounce(i int, w *WordObject) bool {
	result := false
	if fn := "e:/words/" + w.WordName + ".htm"; FileExists(fn) {
		s := ReadFile(w.WordName)
		//fmt.Printf("%s\n:", w.WordName)

		ap, bp := parserPronounce(s)
		w.APronounce = ap
		w.BPronounce = bp
		result = true
		//fmt.Printf("%s %s %s", w.WordName, w.APronounce, w.BPronounce)
	} else {
		fmt.Printf("%s %s无效\n", w.WordName, fn)
	}
	fmt.Printf("%d : %s %s %s saved\n", i, w.WordName, w.APronounce, w.BPronounce)
	return result

}
