package links

import (
	"context"
	"log"

	"github.com/subh1994/graphql-go-server/internal/pkg/db/postgres"
	"github.com/subh1994/graphql-go-server/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	pool := postgres.GetPool()

	queryString := "INSERT INTO Links(Title,Address) VALUES($1,$2) returning id"
	var id int64
	err := pool.QueryRow(context.Background(), queryString, link.Title, link.Address).Scan(&id)
	log.Print("inserted data id--", id)

	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Print("Row inserted!")
	}
	return id
}
