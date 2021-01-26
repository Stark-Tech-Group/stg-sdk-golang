package main

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/client"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	un := os.Getenv(env.STG_SDK_API_UN)
	pw := os.Getenv(env.STG_SDK_API_PW)

	api := client.Client{}
	api.Init()
	apiStatus := api.ApiStatus()

	fmt.Printf("date: %s\n", time.Unix(apiStatus.DateTime/1000, 0))

	api.Login(un, pw)

	equipSearchBody := client.SearchBody{
		Query:       fmt.Sprintf("equip"),
		CurrentPage: 1,
		PageSize:    50,
	}

	equipResp := api.Search(equipSearchBody)

	for _, equip := range equipResp.Assets {
		fmt.Printf(" Equip: %s, id: %d\n", equip.Name, equip.ID)
	}

	testEquip()
}

func testEquip() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	un := os.Getenv(env.STG_SDK_API_UN)
	pw := os.Getenv(env.STG_SDK_API_PW)

	api := client.Client{}
	api.Init()
	api.Login(un, pw)

	siteApi := api.GetSiteEndpoint()
	site, err := siteApi.Get(1)
	if err != nil {
		log.Fatalf("error getting site. %v", err)
	}
	log.Printf( "site %s", site.Name)


}