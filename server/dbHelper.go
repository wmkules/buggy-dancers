package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"

	bolt "go.etcd.io/bbolt"
)

var db *bolt.DB

func setupDB() (*bolt.DB, error) {
	var err error
	db, err = bolt.Open("database/test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		root, err := tx.CreateBucketIfNotExists([]byte("DB"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("BALLOT"))
		if err != nil {
			return fmt.Errorf("could not create ballot bucket: %v", err)
		}
		_, err = root.CreateBucketIfNotExists([]byte("PROMPT"))
		if err != nil {
			return fmt.Errorf("could not create prompt bucket: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}

	fmt.Println("DB Setup Done")
	return db, nil
}

func dbAddBallot(db *bolt.DB, bal ballot) error {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(99999)
	fmt.Println("Adding ballot with id: ", id)

	return dbUpdateBallot(db, bal, []byte(strconv.Itoa(id)))
}

func dbUpdateBallot(db *bolt.DB, bal ballot, id []byte) error {
	balBytes, err := json.Marshal(bal)
	if err != nil {
		return fmt.Errorf("could not marshal ballot json: %v", err)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Put(id, balBytes)
		if err != nil {
			return fmt.Errorf("could not set ballot: %v", err)
		}
		return nil
	})
	return err
}

func dbPrintBallots(db *bolt.DB) error {
	fmt.Println("Printing db now")
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT"))
		b.ForEach(func(k, v []byte) error {
			fmt.Println(string(k), string(v))
			return nil
		})
		return nil
	})
	if err != nil {
		return fmt.Errorf("could not print ballots: %v", err)
	}
	fmt.Println("That was the db")
	return nil
}

func populateDB(db *bolt.DB) {
	for _, bal := range ballots {
		dbAddBallot(db, bal)
	}
}

func dbVote(db *bolt.DB, v vote) (ballot, error) {
	var bal ballot
	err := db.View(func(tx *bolt.Tx) error {
		balBytes := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(v.BallotID)))

		if err := json.Unmarshal(balBytes, &bal); err != nil {
			return fmt.Errorf("could not fetch ballot from db: %v", err)
		}

		newPrompts := []prompt{}

		for _, p := range bal.Prompts {
			if p.ID == v.PromptID {
				p.Votes += 1
			}
			newPrompts = append(newPrompts, p)
		}
		bal.Prompts = newPrompts
		return nil
	})
	if err != nil {
		return bal, fmt.Errorf("could not update vote: %v", err)
	}
	dbUpdateBallot(db, bal, []byte(v.BallotID))
	fmt.Println(bal)
	return bal, nil
}

func dbPrintById(db *bolt.DB, id int) {
	err := db.View(func(tx *bolt.Tx) error {
		bal := tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Get([]byte([]byte(strconv.Itoa(id))))
		fmt.Printf("Ballot: %s\n", bal)
		fmt.Println(reflect.TypeOf(bal))
		return nil
	})
	if err != nil {
		fmt.Printf("could not print ballots: %v", err)
	}
}
