package starkapi

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"

	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//err := godotenv.Load("../../.env.test")
	//if err != nil {
	//		log.Fatal("Error loading .env.test file", err)/
	//}
	code := m.Run()

	os.Exit(code)
}

func TestHisReadPoint(t *testing.T) {

	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := Client{}
	api.Init(host)
	api.Login(un, pw)

	assetTreeApi := api.AssetTreeApi

	assetTree, err := assetTreeApi.Get()
	if err != nil { t.Error("failed", err) }
	fmt.Printf("count: %x\n", assetTree.AssetTreeMeta.Size)

	for _, branch := range assetTree.AssetTree {
		fmt.Printf("val: %s\n", branch.Name)
	}

}
