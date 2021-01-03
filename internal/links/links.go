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

func (link Link) Save() string {
	pool := postgres.GetPool()

	queryString := "INSERT INTO Links(Title,Address) VALUES($1,$2) RETURNING id"

	var id string
	err := pool.QueryRow(context.Background(), queryString, link.Title, link.Address).Scan(&id)
	log.Print("inserted data id--", id)

	if err != nil {
		log.Printf(err.Error())
	} else {
		log.Print("Row inserted!")
	}
	return id
}

func GetAll() []Link {
	pool := postgres.GetPool()

	queryString := "SELECT id, title, address FROM Links"
	rows, err := pool.Query(context.Background(), queryString)

	var links []Link

	if err == nil {
		for rows.Next() {
			var link Link
			err := rows.Scan(&link.ID, &link.Title, &link.Address)
			if err != nil {
				log.Println(err)
			} else {
				links = append(links, link)
			}
		}
	}

	return links
}
