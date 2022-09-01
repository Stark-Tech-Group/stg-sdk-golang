package starkapi

import (
	"testing"
)

func TestAddTagToAsset(t *testing.T) {
	// call needs to be mocked
	/*
		un := os.Getenv(env.STG_SDK_API_UN)
		pw := os.Getenv(env.STG_SDK_API_PW)
		host := os.Getenv(env.STG_SDK_API_HOST)

		api := Client{}
		api.Init(host)
		_, err := api.Login(un, pw)

		if err != nil {
			log.Fatalf("auth err: %s", err)
		}

		//Test valid asset
		asset := domain.Asset{
			Id:   1,
			Ref:  "e.test",
			Url:  "test/url",
			Name: "Test",
			Type: "Equip",
		}

		assetsApi := api.AssetsApi
		assetErr := assetsApi.AddNewTag(asset, "Test", "1")
		assert.Equal(t, nil, assetErr)


		//Test invalid tag
		badTagErr := assetsApi.AddNewTag(asset, "", "1")
		assert.NotEqual(t, nil, badTagErr)

		//Test asset with no Ref
		badAsset := domain.Asset{
			Id:   1,
			Ref:  "",
			Url:  "test/url",
			Name: "",
			Type: "Equip",
		}


		badAssetErr := assetsApi.AddNewTag(badAsset, "Test", "1")

		assert.NotEqual(t, nil, badAssetErr)

	*/
}
