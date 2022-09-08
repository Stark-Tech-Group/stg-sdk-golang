package starkapi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testAssetTreeApiAdminURL=  "/admin/branches"
const testAssetTreeApiCoreURL =  "/core/branches"

func TestAssetTreeApi_BaseUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetTreeApi := api.AssetTreeApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testAssetTreeApiCoreURL), assetTreeApi.baseUrl())
}

func TestAssetTreeApi_AdminUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetTreeApi := api.AssetTreeApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testAssetTreeApiAdminURL), assetTreeApi.adminUrl())
}