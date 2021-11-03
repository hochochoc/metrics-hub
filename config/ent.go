package config

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"metric-hub/ent"
	"time"
)

var (
	client *ent.Client
)

func GetClient() *ent.Client {
	return client
}

func SetClient(newClient *ent.Client) {
	client = newClient
}

func NewEntClient() (*ent.Client, error) {
	pgClient, err := ent.Open(
		"postgres",
		fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
			viper.Get("PG_HOST"),
			viper.Get("PG_PORT"),
			viper.Get("PG_USER"),
			viper.Get("PG_DATABASE"),
			viper.Get("PG_PASSWORD")), ent.Debug(), ent.Log(func(i ...interface{}) {
			for _,v := range i {
				fmt.Println(time.Now().Format("2006-01-02 15:04:05"),v)
				fmt.Print("\n")
			}
		}))
	if err != nil {
		log.Fatalf("failed opening connection to postgresql: %v", err)
	}
	if err := pgClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return pgClient, err
}
