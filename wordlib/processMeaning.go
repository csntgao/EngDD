// processMeaning
package wordlib

import (
	"fmt"
	"regexp"
	"strings"
)

var wordMeaningBlockExp = GaoRegexp{regexp.MustCompile(`<strong class="fl">(?P<meaningBlock>[\s\S]+)</label>[\s\S]+\t{10}</span>`)}
var wordPropertyExp = GaoRegexp{regexp.MustCompile(`<strong class="fl">(?P<WordProperty>.+)</strong>`)}
var wordMeaningExp = GaoRegexp{regexp.MustCompile(`<label>(?P<WordMeaning>[^<]+)</label>`)}

func parserMeaningBlocks(text string) []string {
	result := []string{}
	mmap := wordMeaningBlockExp.FindStringSubmatchMap(text)
	//fmt.Println(mmap)
	for _, ww := range mmap {
		for _, v := range ww {
			result = strings.Split(v, "</span>")

		}

	}
	return result

}

func ArraylizeMeanings(i int, w *WordObject) bool {
	fmt.Printf("%d,%s\n", i, w.WordName)
	for i, meaning := range w.Meanings {
		newMeaning := []string{}
		for _, s := range meaning.Meaning {
			if strings.Contains(s, "，") {
				newMeaning = append(newMeaning, strings.Split(s, "，")...)
			} else {
				newMeaning = append(newMeaning, s)
			}
		}
		w.Meanings[i].Meaning = newMeaning
	}
	return true
}
func MeaningBlocksToMeanings(i int, w *WordObject) bool {
	//result := false
	fmt.Printf("%d %s Meanings are building\n", i, w.WordName)
	meanings := []WordMeaning{}

	for _, s := range w.MeaningBlocks {
		if !(strings.HasPrefix(s, "</p><p><strong class=\"fl\">")) {
			s = "</p><p><strong class=\"fl\">" + s
		}
		if !(strings.HasSuffix(s, "</label>")) {
			s = s + "</label>"
		}
		wm := new(WordMeaning)
		mmap := wordPropertyExp.FindStringSubmatchMap(s)
		for _, ww := range mmap {
			for _, v := range ww {
				wm.Property = v
			}

		}
		mmap2 := wordMeaningExp.FindStringSubmatchMap(s)
		for _, ww := range mmap2 {
			for _, v := range ww {
				if strings.HasSuffix(v, "；") {
					v = strings.Replace(v, "；", "", -1)
				}
				wm.Meaning = append(wm.Meaning, v)
			}

		}
		meanings = append(meanings, *wm)

	}
	w.Meanings = meanings
	return true

}

func DeleteMeaningBlockSpace(i int, w *WordObject) bool {
	result := true
	newblock := []string{}
	for _, s := range w.MeaningBlocks {
		s = strings.Replace(s, "\t", "", -1)
		s = strings.Replace(s, "\n", "", -1)
		s = strings.Replace(s, "\r", "", -1)
		newblock = append(newblock, s)

	}
	w.MeaningBlocks = newblock
	return result

}

func PeekMeaningBlocks(i int, w *WordObject) bool {
	result := false
	if fn := "e:/words/" + w.WordName + ".htm"; FileExists(fn) {
		s := ReadFile(w.WordName)

		w.MeaningBlocks = parserMeaningBlocks(s)
		result = true //测试版本，先不写库
		//w.MeaningBlocksw.MeaningBlocks = meaningBlocks
		fmt.Printf("%d %s has been finished\n", i, w.WordName)
		//fmt.Println(meaningBlocks)

	} else {
		fmt.Printf("%s %s无效\n", w.WordName, fn)
	}
	return result

}
