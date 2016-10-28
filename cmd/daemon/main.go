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

	logs := make(chan mgotail.Oplog, 1)
	done := make(chan bool, 1)
	go q.Tail(logs, done)
	go func(done chan bool) {
		for log := range logs {
			for _, watcher := range oplord.Watchers {
				if watcher.Matcher(log) {
					if err := watcher.Action(log); err != nil {
						done <- true
					}
					if err := oplord.SaveTime(session, log.Timestamp); err != nil {
						done <- true
					}
				}
			}
		}
	}(done)
	<-done
}
