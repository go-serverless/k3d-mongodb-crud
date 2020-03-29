package function

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/openfaas/openfaas-cloud/sdk"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Foo struct {
	Bar string
	Baz string
}

var (
	sess            *mgo.Session
	mongoDatabase   = os.Getenv("mongo_database")
	mongoCollection = os.Getenv("mongo_collection")
)

func init() {
	var err error
	mongoHost := os.Getenv("mongo_host")
	mongoUsername, _ := sdk.ReadSecret("mongo-db-username")
	mongoPassword, _ := sdk.ReadSecret("mongo-db-password")

	if _, err := os.Open("/var/openfaas/secrets/mongo-db-password"); err != nil {
		panic(err.Error())
	}

	info := &mgo.DialInfo{
		Addrs:    []string{mongoHost},
		Timeout:  60 * time.Second,
		Database: mongoDatabase,
		Username: mongoUsername,
		Password: mongoPassword,
	}

	if sess, err = mgo.DialWithInfo(info); err != nil {
		panic(err.Error())
	}
}

func Handle(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		fmt.Println("4 records will be inserted")

		if err := sess.DB(mongoDatabase).C(mongoCollection).Insert(
			&Foo{Bar: "bar", Baz: "baz"},
			&Foo{Bar: "bar1", Baz: "baz1"},
			&Foo{Bar: "bar2", Baz: "baz2"},
			&Foo{Bar: "bar3", Baz: "baz3"},
		); err != nil {
			http.Error(w, fmt.Sprintf("Failed to insert: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":true, "message": "4 records have been inserted"}`))

	} else if r.Method == http.MethodGet {
		fmt.Println("All records will be listed")

		var foo []Foo
		err := sess.DB(mongoDatabase).C(mongoCollection).Find(bson.M{}).All(&foo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to read: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		out, err := json.Marshal(foo)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to marshal: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(out)

	} else if r.Method == http.MethodPut {
		fmt.Println("bar1 will be updated to bar1-updated")

		if err := sess.DB(mongoDatabase).C(mongoCollection).Update(bson.M{"bar": "bar1"}, bson.M{"$set": bson.M{"bar": "bar1-updated"}}); err != nil {
			http.Error(w, fmt.Sprintf("Failed to update: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":true, "message": "bar1 has been updated to bar1-updated"}`))

	} else if r.Method == http.MethodDelete {
		fmt.Println("bar will be deleted")

		if err := sess.DB(mongoDatabase).C(mongoCollection).Remove(bson.M{"bar": "bar"}); err != nil {
			http.Error(w, fmt.Sprintf("Failed to delete: %s", err.Error()), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"success":true, "message": "bar1 has been deleted"}`))
	}
}
