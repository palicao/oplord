package oplord

import (
	"bytes"
	"encoding/json"
	"github.com/Clever/mgotail"
	"net/http"
)

type OplogMatcher func(mgotail.Oplog) bool
type OplogAction func(mgotail.Oplog) error

type Watcher struct {
	Matcher OplogMatcher
	Action  OplogAction
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func SimpleMatcherFactory(collection string, operations []string) OplogMatcher {
	return func(log mgotail.Oplog) bool {
		return collection == log.Namespace && stringInSlice(log.Operation, operations)
	}
}

func SimplePostActionFactory(url string) OplogAction {
	return func(log mgotail.Oplog) error {
		b, err := json.Marshal(log.Object)
		if err != nil {
			return err
		}
		go http.Post(url, "application/json", bytes.NewBuffer(b))
		return nil
	}
}
