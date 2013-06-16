// wordUseness
package wordlib

import (
	//"fmt"
	"fmt"
	"regexp"
	"strings"
)

var wordPhrasesBlockExp = GaoRegexp{regexp.MustCompile(`<dt>常用词组</dt>(?P<PhrasesBlock>[\s\S]+)<dt>习惯用语</dt>`)}
var wordIdiomsBlockExp = GaoRegexp{regexp.MustCompile(`<dt>习惯用语</dt>(?P<iDiomsBlock>[\s\S]+)<p id="cizi_see_more" class="see_more">`)}
var wordSentencesBlockExp = GaoRegexp{regexp.MustCompile(`>双语例句</a>[\s\S]+style="display:block">(?P<SentenceBlock>[\s\S]+)查看更多例句>>`)}
var wordSentenceExp = GaoRegexp{regexp.MustCompile(`<dt>\d*\.(?P<English>[\s\S]+)<a class="ico_sound"[\s\S]+</dt> <dd>(?P<Chinese>[\s\S]+)</dd> <p>来自(?P<Source>[\s\S]+)`)}
var wordPhraseExp = GaoRegexp{regexp.MustCompile(`<dd class=[\s\S]+ sid="cx_\d+">(?P<Phrase>[\s\S]+)</h4>(?P<PhraseMeaningBlock>[\s\S]+)`)}
var PhraseMeaningExp = GaoRegexp{regexp.MustCompile(`[\s\S]+<h5>\d*\.(?P<M>[\s\S]+)</h5>(?P<S>[\s\S]+)`)}

func deleteUselessSpace(text string) string {
	s := []rune(text)
	r := []rune{}
	var lastChar rune = -1
	for _, ch := range s {
		if ch == ' ' {
			if lastChar != ' ' {
				r = append(r, ch)
			}

		} else {
			r = append(r, ch)
		}
		lastChar = ch
	}
	return string(r)
}

func parserItem(exp GaoRegexp, text string) string {
	result := ""
	mmap := exp.FindStringSubmatchMap(text)
	//fmt.Println(mmap)
	for _, ww := range mmap {
		for _, v := range ww {
			result = strings.Replace(v, "\n", "", -1)
			result = strings.Replace(result, "\r", "", -1)
			result = strings.Replace(result, "\t", "", -1)
			result = deleteUselessSpace(result)
			break
		}

	}
	return result
}

func parserPhrasesBlock(text string) string {
	return parserItem(wordPhrasesBlockExp, text)
}

func parserIdiomsBlock(text string) string {
	return parserItem(wordIdiomsBlockExp, text)
}

func parserSentencesBlock(text string) string {
	return parserItem(wordSentencesBlockExp, text)
}

func PeekWordUseness(i int, w *WordObject) bool {
	text := ReadFile(w.WordName)
	w.SentencesBlock = parserSentencesBlock(text)
	w.PhrasesBlock = parserPhrasesBlock(text)
	w.IdiomsBlock = parserIdiomsBlock(text)

	fmt.Printf("%d %s:\n %s\n %s\n %s\n\n", i, w.WordName, w.PhrasesBlock, w.IdiomsBlock, w.SentencesBlock)
	return true
}

func ShowWord(i int, w *WordObject) bool {
	fmt.Printf("%d %s:\n %s\n %s\n %s\n", i, w.WordName, w.IdiomsBlock, w.PhrasesBlock, "")
	return false

}
func delHtmlTags(e string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return re.ReplaceAllString(e, "")

}

func SentenceBlockToSentences(i int, w *WordObject) bool {

	rawSentences := strings.Split(w.SentencesBlock, "</p>")
	sentences := []WordSentenceObject{}
	for _, rs := range rawSentences {
		mmap := wordSentenceExp.FindStringSubmatchMap(rs)
		for _, sentenceStructure := range mmap {
			e := sentenceStructure["English"]
			e = delHtmlTags(e)
			c := sentenceStructure["Chinese"]
			s := sentenceStructure["Source"]

			ws := new(WordSentenceObject)
			ws.ChineseSentence = c
			ws.EnglishSentence = e
			ws.Source = s
			sentences = append(sentences, *ws)

			fmt.Printf("%d %s:\n %s\n %s\n %s\n", i, w.WordName, e, c, s)

		}

	}
	result := true
	w.Sentences = sentences
	return result

}

func ParseFromString(phrase string, mb string) WordPhraseObject {
	wph := new(WordPhraseObject)
	wph.Phrase = strings.TrimSpace(phrase)
	rawMeanings := strings.Split(mb, "</ul>")
	for _, s := range rawMeanings {
		mmap := PhraseMeaningExp.FindStringSubmatchMap(s)
		for _, PhraseStructure := range mmap {

			meaning := PhraseStructure["M"]
			Sample := PhraseStructure["S"]
			CSample := ""
			ESample := ""
			arr := strings.Split(Sample, "</li>")
			if len(arr) >= 2 {
				ESample = delHtmlTags(arr[0])
				CSample = delHtmlTags(arr[1])
			} else {
				CSample = ""
				ESample = ""
			}
			pm := new(PhraseMeaningObject)
			pm.Meaning = meaning
			pm.CSample = CSample
			pm.ESample = ESample
			wph.Meanings = append(wph.Meanings, *pm)
		}

	}
	return *wph
}

func PhraseBlockToPhrases(i int, w *WordObject) bool {

	fmt.Printf("%d %s Phrases is being built!  \n", i, w.WordName)

	rawPhrases := strings.Split(w.PhrasesBlock, "</dd>")
	Phrases := []WordPhraseObject{}
	for _, rs := range rawPhrases {
		mmap := wordPhraseExp.FindStringSubmatchMap(rs)
		for _, PhraseStructure := range mmap {

			p := PhraseStructure["Phrase"]
			mb := PhraseStructure["PhraseMeaningBlock"]
			phraseMeanings := ParseFromString(p, mb)

			//fmt.Printf("%d %s %d: %s  \n", i, w.WordName, j, p)
			//fmt.Println(phraseMeanings)
			Phrases = append(Phrases, phraseMeanings)

		}

	}
	result := true
	w.Phrases = Phrases

	return result

}

func IdiomBlockToIdioms(i int, w *WordObject) bool {

	fmt.Printf("%d  the idioms of  %s is being built!  \n", i, w.WordName)

	rawPhrases := strings.Split(w.IdiomsBlock, "</dd>")
	Phrases := []WordPhraseObject{}
	for _, rs := range rawPhrases {
		mmap := wordPhraseExp.FindStringSubmatchMap(rs)
		for _, PhraseStructure := range mmap {

			p := PhraseStructure["Phrase"]
			mb := PhraseStructure["PhraseMeaningBlock"]
			phraseMeanings := ParseFromString(p, mb)

			//fmt.Printf("%d %s %d: %s  \n", i, w.WordName, j, p)
			//fmt.Println(phraseMeanings)
			Phrases = append(Phrases, phraseMeanings)

		}

	}
	result := true
	w.Idioms = Phrases

	return result

}

func DeleteUnusefulBlocks(i int, w *WordObject) bool {

	fmt.Printf("%d  The unuseful  informations of  %s is being deleted!  \n", i, w.WordName)
	w.IdiomsBlock = ""
	w.MeaningBlocks = []string{}
	w.PhrasesBlock = ""
	w.SentencesBlock = ""
	return true

}
