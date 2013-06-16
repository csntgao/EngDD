// word
package wordlib

import (
	"fmt"
	//"strconv"

	"strings"
	//"io/ioutil"
	//"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type WordMeaning struct {
	Property string
	Meaning  []string
}
type WordSentenceObject struct {
	EnglishSentence string
	ChineseSentence string
	Source          string
}
type PhraseMeaningObject struct {
	Meaning string
	ESample string
	CSample string
}
type WordPhraseObject struct {
	Phrase   string
	Meanings []PhraseMeaningObject
}

type LikeItem struct {
	Item        string
	Similiarity float32
}
type WordObject struct {
	Id             bson.ObjectId        "_id"
	WordName       string               "wordName"
	WordValue      string               "wordValue"
	APronounce     string               "aPronounce"
	BPronounce     string               "bPronounce"
	Catalogs       []string             "catalogs"
	MeaningBlocks  []string             "meaningBlocks"
	Meanings       []WordMeaning        "meanings"
	PhrasesBlock   string               "phrasesBlock"
	IdiomsBlock    string               "iDiomsBlock"
	SentencesBlock string               "sentencesBlock"
	Sentences      []WordSentenceObject "sentences"
	Phrases        []WordPhraseObject   "phrases"
	Idioms         []WordPhraseObject   "idioms"
	SpellLikes     []LikeItem           "spellLikes"
	PronounceLikes []LikeItem           "pronounceLikes"
	MeaningLikes   []LikeItem           "meaningLikes"

	SpellSimilarWord     string
	PronounceSimilarWord string
	MeaningSimilarWord   string
	Usage                string
	EPhrase              string
	CPhrase              string
	ESentence            string
	CSentence            string
}

type RelationObject struct {
	Id       bson.ObjectId "_id"
	Word1    string        "word1"
	Word2    string        "word2"
	Relation string        "relation"
	Value    float32       "value"
}

func min(a int, b int) int {
	result := a
	if a > b {
		result = b
	}
	return result
}

func EditDistance(Word1 string, Word2 string) int {

	runeWord1 := []rune(strings.ToLower(Word1))
	runeWord2 := []rune(strings.ToLower(Word2))

	if len(runeWord1) == 0 {
		return len(runeWord2)
	}
	if len(runeWord2) == 0 {
		return len(runeWord1)
	}

	result := 0
	alen := len(runeWord1) + 1
	dist := make([]int, (len(runeWord1)+1)*(len(runeWord2)+1))
	for i := 0; i <= len(runeWord1); i++ {
		dist[i] = i

	}
	for j := 0; j <= len(runeWord2); j++ {
		dist[j*alen] = j
	}
	for i := 1; i <= len(runeWord1); i++ {
		for j := 1; j <= len(runeWord2); j++ {
			if runeWord1[i-1] == runeWord2[j-1] {

				dist[j*alen+i] = dist[(j-1)*alen+i-1]
			} else {
				minDist := min(dist[j*alen+(i-1)], dist[(j-1)*alen+i])
				dist[j*alen+i] = min(minDist, dist[(j-1)*alen+i-1]) + 1
			}
		}
	}
	result = dist[len(dist)-1]
	//fmt.Println(dist)
	return result
}

func AddCatalog(i int, w *WordObject) bool {
	w.Catalogs = append([]string{}, "Kaoyan2013")
	fmt.Printf("%d : %s Catalog  %s saved\n", i, w.WordName, "Kaoyan2013")
	return true
}
func BuildSpellRelations(w1 *WordObject, w2 *WordObject) (bool, *RelationObject) {
	suc := false
	var pR *RelationObject = nil

	i := EditDistance(w1.WordName, w2.WordName)
	minLen := min(len(w1.WordName+w2.WordName)/4, 2)
	if i <= minLen {
		R := RelationObject{bson.NewObjectId(), w1.WordName, w2.WordName, "SpellLike", float32(-i)}
		suc = true
		pR = &R
		//fmt.Println(R)

	}
	return suc, pR
}

func BuildPronounceRelations(w1 *WordObject, w2 *WordObject) (bool, *RelationObject) {
	suc := false
	var pR *RelationObject = nil

	i := EditDistance(w1.APronounce, w2.APronounce)
	minLen := min(len(w1.WordName+w2.WordName)/4, 2)
	if i <= minLen {
		R := RelationObject{bson.NewObjectId(), w1.WordName, w2.WordName, "PronounceLike", float32(-i)}
		suc = true
		pR = &R
		//fmt.Println(R)

	}
	return suc, pR
}

func MeaningsCloseness(w1 *WordObject, w2 *WordObject) int {
	result := 0
	arr1 := []string{}
	arr2 := []string{}
	for _, m := range w1.Meanings {
		for _, s := range m.Meaning {
			arr1 = append(arr1, m.Property+s)
		}
	}

	for _, m := range w2.Meanings {
		for _, s := range m.Meaning {
			arr2 = append(arr2, m.Property+s)
		}
	}

	for _, s := range arr1 {
		for _, t := range arr2 {
			if s == t {
				result++
				break
			}
		}
	}
	return result

}

func BuildMeaningRelations(w1 *WordObject, w2 *WordObject) (bool, *RelationObject) {
	suc := false
	var pR *RelationObject = nil

	i := MeaningsCloseness(w1, w2)
	if i > 0 {
		R := RelationObject{bson.NewObjectId(), w1.WordName, w2.WordName, "MeaningLike", float32(i)}
		suc = true
		pR = &R
		fmt.Println(R)

	}
	return suc, pR
}

func (this *WordObject) UpdateBaseInfo() {
	this.SpellSimilarWord = ""
	var s float32 = -100.0
	for _, li := range this.SpellLikes {
		if li.Similiarity > s {
			s = li.Similiarity
			this.SpellSimilarWord = li.Item
		}
	}
	this.PronounceSimilarWord = ""
	s = -100.0
	for _, li := range this.PronounceLikes {
		if li.Similiarity > s {
			s = li.Similiarity
			this.PronounceSimilarWord = li.Item
		}
	}
	this.MeaningSimilarWord = ""
	s = -100.0
	for _, li := range this.MeaningLikes {
		if li.Similiarity > s {
			s = li.Similiarity
			this.MeaningSimilarWord = li.Item
		}
	}
	this.Usage = ""
	for _, m := range this.Meanings {
		//abw.Usage 
		s := m.Property
		for i, n := range m.Meaning {
			if i == 0 {
				s = s + " " + n
			} else {
				s += ";" + n
			}

		}
		this.Usage += s + "<BR/>"
	}
	for _, p := range this.Phrases {
		this.EPhrase = p.Phrase
		for _, m := range p.Meanings {
			this.CPhrase += m.Meaning + ";"
		}
		break

	}
	for _, s := range this.Sentences {
		this.ESentence = s.EnglishSentence
		this.CSentence = s.ChineseSentence
		break
	}

}
