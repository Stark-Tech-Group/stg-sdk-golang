package starkapi

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testFormsApiURL =  "/core/forms"
const testFormsApiURLWithRef =  "/core/forms/e.test/controls"

func TestFormsApi_host(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, testHost, formsApi.host())
}

func TestFormsApi_baseUrl(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, fmt.Sprintf("%s%s", testHost, testFormsApiURL), formsApi.baseUrl())
}

func TestFormsApi_controlsPrefix(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	formsApi := api.FormsApi
	assert.Equal(t, "/controls", formsApi.controlsPrefix())
}

func TestFormsApi_GetAllControls(t *testing.T) {
	api := Client{}
	host := testHost
	api.Init(host)
	assetsApi := api.AssetsApi
	assert.Equal(t, testHost, assetsApi.host())
}

func TestFormsApi_GetAllControlsForAsset(t *testing.T) {

}