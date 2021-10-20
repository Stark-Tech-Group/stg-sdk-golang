package starkapi

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"github.com/joho/godotenv"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.secrets/.env")
	if err != nil {

		fmt.Printf("Error loading .env.test file %v", err)
	}
	code := m.Run()

	os.Exit(code)
}

func TestGet(t *testing.T) {

	un := os.Getenv(env.STG_SDK_API_UN)
	pw := os.Getenv(env.STG_SDK_API_PW)
	host := os.Getenv(env.STG_SDK_API_HOST)

	api := Client{}
	api.Init(host)
	api.Login(un, pw)

	assetTreeApi := api.AssetTreeApi

	assetTree, err := assetTreeApi.Get()
	if err != nil {
		t.Error("failed", err)
	}
	fmt.Printf("count: %x\n", assetTree.AssetTreeMeta.Size)

	for _, branch := range assetTree.AssetTree {
		fmt.Printf("val: %s\n", branch.Name)
	}

}
