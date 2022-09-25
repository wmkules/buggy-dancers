package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
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
	id := rand.Intn(50000)
	fmt.Println("Adding ballot with id: ", id)

	balBytes, err := json.Marshal(bal)
	if err != nil {
		return fmt.Errorf("could not marshal ballot json: %v", err)
	}

	// fmt.Println("Ballot is:", balBytes)

	err = db.Update(func(tx *bolt.Tx) error {
		err = tx.Bucket([]byte("DB")).Bucket([]byte("BALLOT")).Put([]byte(strconv.Itoa(id)), balBytes)
		if err != nil {
			return fmt.Errorf("could not set ballot: %v", err)
		}
		return nil
	})
	fmt.Println("Set Ballot")
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
