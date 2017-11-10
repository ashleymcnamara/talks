package main

import (
	"log"

	"github.com/bketelsen/talks/codemotion/gc18/services/profile-srv/db"
	"github.com/bketelsen/talks/codemotion/gc18/services/profile-srv/handler"
	"github.com/bketelsen/talks/codemotion/gc18/services/profile-srv/proto/record"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
)

const ServiceName = "gophercon.srv.profile"

func main() {

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Flags(
			cli.StringFlag{
				Name:   "database_url",
				EnvVar: "DATABASE_URL",
				Usage:  "The database URL e.g root@tcp(127.0.0.1:3306)/profile",
			},
		),

		micro.Action(func(c *cli.Context) {
			if len(c.String("database_url")) > 0 {
				db.Url = c.String("database_url")
			}
		}),
	)

	service.Init()

	db.Init()

	record.RegisterRecordHandler(service.Server(), new(handler.Record))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
