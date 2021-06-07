package starkapi

import (
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func TestGeoApi_ClimateZone(t *testing.T) {

	un := os.Getenv(env.STG_SDK_API_UN)
	pw := os.Getenv(env.STG_SDK_API_PW)
	host := os.Getenv(env.STG_SDK_API_HOST)

	api := Client{}
	api.Init(host)
	_, err := api.Login(un, pw)

	if err != nil {
		log.Fatalf("auth err: %s", err)
	}

	geoApi := api.GeoApi

	climateZone, err := geoApi.ClimateZone("Erie", "NY")

	if err != nil {
		t.Error("failed", err)
	}

	assert.Equal(t, "5", climateZone.Zone)
	assert.Equal(t, "Cold", climateZone.Ba)
	assert.Equal(t, "A", climateZone.MoistureRegime)
	assert.Equal(t, "NY", climateZone.State)

}

func TestGeoApi_GeoCoding(t *testing.T) {

	un := os.Getenv(env.STG_SDK_API_UN)
	pw := os.Getenv(env.STG_SDK_API_PW)
	host := os.Getenv(env.STG_SDK_API_HOST)

	api := Client{}
	api.Init(host)
	_, err := api.Login(un, pw)

	if err != nil {
		log.Fatalf("auth err: %s", err)
	}

	geoApi := api.GeoApi

	geoCoding, err := geoApi.GeoCoding("95 Stark St", "Tonawanda", "NY", "14215")

	if err != nil {
		t.Error("failed", err)
	}

	assert.Equal(t, "Erie", geoCoding.County)

}
