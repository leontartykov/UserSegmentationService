package dbclient

import (
	"fmt"
	"log"
	"main/server/session"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Client struct {
	Db *sqlx.DB
}

func NewDbConnection(dbConfig *session.DB) (*Client, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.User,
		dbConfig.DBName,
		dbConfig.Password,
		dbConfig.SSLMode))

	if err != nil {
		log.Println("Error while connecting to db")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Error ping to db")
		return nil, err
	}

	log.Println("Connection to db is successfull")
	return &Client{
		Db: db,
	}, nil
}
