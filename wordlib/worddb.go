// worddb
package wordlib

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type WordDatabase struct {
	MongoURL            string
	WordsCollection     *mgo.Collection
	RelationsCollection *mgo.Collection
	session             *mgo.Session
}

type ProcessWordFunc func(i int, w *WordObject) bool
type ProcessRelationFunc func(i int, r *RelationObject, w1 *WordObject, w2 *WordObject) bool
type GetRelationFunc func(w1 *WordObject, w2 *WordObject) (bool, *RelationObject)

func (this *WordDatabase) alert(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		panic(err)
	}
}

func (this *WordDatabase) FindWords(query interface{}, result interface{}) error {
	err := this.WordsCollection.Find(query).All(result)
	this.alert(err, "FindWords错误！")
	return err

}

func (this *WordDatabase) FindRelations(query interface{}, result interface{}) error {
	err := this.RelationsCollection.Find(query).All(&result)
	this.alert(err, "FindWords错误！")
	return err

}
func (this *WordDatabase) Open() {
	var err error
	this.session, err = mgo.Dial("127.0.0.1")
	this.alert(err, "数据库打开错误！")
	this.WordsCollection = this.session.DB("Vocabulary").C("words")
	this.RelationsCollection = this.session.DB("Vocabulary").C("relations")

}

func (this *WordDatabase) Close() {
	this.session.Close()

}

//func (this *WordDatabase) FindWords()
func (this *WordDatabase) BuildAllRelations(p GetRelationFunc) {

	//this.Open()
	//defer this.Close()
	words1 := []WordObject{}
	words2 := []WordObject{}
	this.FindWords(&bson.M{}, &words1)

	for i, w1 := range words1 {
		this.FindWords(&bson.M{"wordName": bson.M{"$gt": w1.WordName}}, &words2)
		for _, w2 := range words2 {

			suc, r := p(&w1, &w2)
			if suc {
				this.RelationsCollection.Insert(r)
				fmt.Printf("%d %s %s的关系建立如下：\n", i, w1.WordName, w2.WordName)
				fmt.Println(r)
			}
		}

	}
	fmt.Println("Finished!")
	//err = c.Find(&bson.M{"wordName": bson.M{"$gte": "a"}}).All(&words)

}
func (this *WordDatabase) GetWordObject(wordName string) *WordObject {
	ws := []WordObject{}
	this.FindWords(bson.M{"wordName": wordName}, &ws)
	for _, r1 := range ws {
		return &r1
	}

	result := (*WordObject)(nil)
	return result

}

func UpdateWordRelation(i int, r *RelationObject, w1 *WordObject, w2 *WordObject) bool {
	fmt.Printf("%d:   ", i)

	if r.Relation == "MeaningLike" {
		li1 := new(LikeItem)
		li1.Item = r.Word2
		li1.Similiarity = r.Value
		w1.MeaningLikes = append(w1.MeaningLikes, *li1)
		li2 := new(LikeItem)
		li2.Item = r.Word1
		li2.Similiarity = r.Value
		w2.MeaningLikes = append(w2.MeaningLikes, *li2)

	}
	if r.Relation == "PronounceLike" {
		li1 := new(LikeItem)
		li1.Item = r.Word2
		li1.Similiarity = r.Value
		w1.PronounceLikes = append(w1.PronounceLikes, *li1)
		li2 := new(LikeItem)
		li2.Item = r.Word1
		li2.Similiarity = r.Value
		w2.PronounceLikes = append(w2.PronounceLikes, *li2)

	}

	if r.Relation == "SpellLike" {
		li1 := new(LikeItem)
		li1.Item = r.Word2
		li1.Similiarity = r.Value
		w1.SpellLikes = append(w1.SpellLikes, *li1)
		li2 := new(LikeItem)
		li2.Item = r.Word1
		li2.Similiarity = r.Value
		w2.SpellLikes = append(w2.SpellLikes, *li2)

	}

	fmt.Println(*r)
	fmt.Println(w1.MeaningLikes, w1.SpellLikes, w1.PronounceLikes)
	fmt.Println(w2.MeaningLikes, w2.SpellLikes, w2.PronounceLikes)

	return true
}

func (this *WordDatabase) GoOverRelations(p ProcessRelationFunc) {

	rs := []RelationObject{}
	this.FindRelations(&bson.M{}, &rs)

	for i, r := range rs {
		w1 := this.GetWordObject(r.Word1)
		w2 := this.GetWordObject(r.Word2)
		if p(i, &r, w1, w2) {
			this.WordsCollection.UpdateId(w1.Id, w1)
			this.WordsCollection.UpdateId(w2.Id, w2)

		}
	}

}
func (this *WordDatabase) GoOverWords(p ProcessWordFunc) {

	// 连接数据库  

	/*
	   // 存储数据  

	   m1 := Mail{bson.NewObjectId(), "user1", "user1@dotcoo.com"}  

	   m2 := Mail{bson.NewObjectId(), "user1", "user2@dotcoo.com"}  

	   m3 := Mail{bson.NewObjectId(), "user3", "user3@dotcoo.com"}  

	   m4 := Mail{bson.NewObjectId(), "user3", "user4@dotcoo.com"}  

	   err = c.Insert(&m1, &m2, &m3, &m4)  

	   if err != nil {  

	       panic(err)  

	   }  

	*/

	// 读取数据  

	words := []WordObject{}

	//err = c.Find(&bson.M{"wordName": bson.M{"$gte": "a"}}).All(&words)
	//err = c.Find(&bson.M{"wordName": "take"}).All(&words)
	this.FindWords(&bson.M{}, &words)

	// 显示数据  

	for i, w := range words {
		if p(i, &w) {
			//fmt.Printf("%d : %s %s %s saved\n", i, w.WordName, w.APronounce, w.BPronounce)
			this.WordsCollection.UpdateId(w.Id, w)
		}
		//fmt.Printf("%d %s %s is not OK\n", i, w.WordName, w.WordValue)

		//s := httprequest(w.WordName)
		//writefile(w.WordName, s)
		/*
			if !fileExists("e:/words/" + w.WordName + ".htm") {
				fmt.Printf("%d %s %s is not OK\n", i, w.WordName, w.WordValue)
				s := httprequest(w.WordName)
				writefile(w.WordName, s)

			}
		*/
	}
	//runtime.Gosched()
	fmt.Println("Finished")

}

var WordDB WordDatabase

func Init() {
	WordDB.MongoURL = "127.0.0.1"
	fmt.Println("Database on 127.0.0.1")
}
