package main

import (
	"github.com/Clever/mgotail"
	"github.com/palicao/oplord"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	q := mgotail.OplogQuery{
		Session: session,
		Query:   bson.M{"ts": bson.M{"$gt": oplord.GetLastTime(session)}},
		Timeout: -1, // See http://godoc.org/gopkg.in/mgo.v2#Query.Tail
	}

	logs := make(chan mgotail.Oplog)
	done := make(chan bool)
	go q.Tail(logs, done)
	go func() {
		for log := range logs {
			for _, watcher := range oplord.Watchers {
				if watcher.Matcher(log) {
					watcher.Action(log)
					go oplord.SaveTime(session, log.Timestamp)
				}
			}
		}
	}()
	<-done
}
