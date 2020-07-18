package config

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
)


const (
	// timeout defines the sleep duration between each connection retry.
	timeout = 5 * time.Second

	// retry defines the number of times to retry.
	retry = 3
)

func InitGORM(dialect, args string) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
		i   int
	)

retry:
	for {
		db, err = gorm.Open(dialect, args)
		switch {
		case err == nil:
			break retry
		case i >= retry:
			return nil, err
		default:
			log.Println(err)
			i++
		}
		time.Sleep(timeout)
	}

	if pingDatabase(db) != nil {
		return db, err
	}

	return db, nil
}


func pingDatabase(db *gorm.DB) error {
	for i := 0; i < retry; i++ {
		err := db.DB().Ping()

		if err == nil {
			return nil
		}

		log.Println("Database ping failed, retry in 1s")
		time.Sleep(timeout)
	}

	return fmt.Errorf("Database ping attempts failed")
}