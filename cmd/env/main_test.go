package cmd

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/client"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env.test")
	if err != nil {
		log.Fatal("Error loading .env.test file", err)
	}
	code := m.Run()

	os.Exit(code)
}

func TestApiStatus(t *testing.T) {
	host := os.Getenv(env.STG_SDK_API_HOST)
	api := client.Client{}
	api.Init(host)

	status, err := api.ApiStatus()
	if err != nil {
		t.Error("Failed to get apis status")
	}

	fmt.Printf(" Api = { Build: %x, Version: %s, Date: %d }",
		status.Build, status.Version, status.DateTime)

}

func TestApiSearch(t *testing.T) {
	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := client.Client{}
	api.Init(host)

	api.Login(un, pw)

	equipSearchBody := client.Query{
		Query:       fmt.Sprintf("equip"),
		CurrentPage: 1,
		PageSize:    50,
	}

	_, err := api.Search(equipSearchBody)
	if err != nil {
		t.Error("Failed to get search", err)
	}
}

func TestGetAllSites(t *testing.T) {
	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := client.Client{}
	api.Init(host)

	api.Login(un, pw)

	fmt.Printf("un: %s\n", un)

	siteApi := api.SiteApi
	sites, err := siteApi.GetAll()

	if err != nil {
		t.Error("Failed to get SiteApi.GetAll", err)
	}
	fmt.Printf("count: %x\n", sites.Count)

	for _, site := range sites.Sites {
		fmt.Printf("Site: %p\n", &site)
	}

}
