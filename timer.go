package oplord

import (
	"github.com/Clever/mgotail"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Timer struct {
	Id string              `bson:"_id"`
	Ts bson.MongoTimestamp `bson:"ts"`
}

func SaveTime(s *mgo.Session, t bson.MongoTimestamp) error {
	_, err := s.DB("oplord").C("times").UpsertId("1", Timer{Id: "1", Ts: t})
	return err
}

func GetLastTime(s *mgo.Session) bson.MongoTimestamp {
	var result Timer
	err := s.DB("oplord").C("times").FindId("1").One(&result)
	if err != nil {
		return mgotail.LastTime(s)
	}
	return result.Ts
}
