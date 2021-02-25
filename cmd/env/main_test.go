package cmd

import (
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/env"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/starkapi"
	"log"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//err := godotenv.Load("../../.env.test")
	//if err != nil {
//		log.Fatal("Error loading .env.test file", err)/
	//}
	code := m.Run()

	os.Exit(code)
}

func TestApiStatus(t *testing.T) {
	host := os.Getenv(env.STG_SDK_API_HOST)
	api := starkapi.Client{}
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

	api := starkapi.Client{}
	api.Init(host)

	api.Login(un, pw)

	equipSearchBody := starkapi.Query{
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

	api := starkapi.Client{}
	api.Init(host)

	api.Login(un, pw)

	fmt.Printf("un: %s\n", un)

	siteApi := api.SiteApi
	items, err := siteApi.GetAll()

	if err != nil {
		t.Error("failed", err)
	}
	fmt.Printf("count: %x\n", items.Count)

	for _, item := range items.Sites {
		fmt.Printf("Site: %p\n", &item)
	}

}

func TestGetAllProfiles(t *testing.T) {
	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := starkapi.Client{}
	api.Init(host)

	api.Login(un, pw)

	fmt.Printf("un: %s\n", un)

	profileApi := api.ProfileApi
	items, err := profileApi.GetAll()

	if err != nil {
		t.Error("failed", err)
	}

	fmt.Printf("count: %x\n", items.Count)

	for _, item := range items.Profiles {
		fmt.Printf("profile: %p\n", &item)
	}

}

func TestGetAllConns(t *testing.T) {
	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := starkapi.Client{}
	api.Init(host)
	api.Login(un, pw)

	connApi := api.ConnApi
	items, err := connApi.GetAll()

	if err != nil { t.Error("failed", err) }

	fmt.Printf("count: %x\n", items.Count)

	for _, item := range items.Conns {
		fmt.Printf("conn: %p\n", &item)
	}
}

func TestGetAllEquips(t *testing.T) {
	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := starkapi.Client{}
	api.Init(host)
	api.Login(un, pw)

	equipApi := api.EquipApi
	items, err := equipApi.GetAll()

	if err != nil { t.Error("failed", err) }

	fmt.Printf("count: %x\n", items.Count)

	for _, item := range items.Equips {
		fmt.Printf("equip: %p\n", &item)
	}
}
func TestGetOnePoint(t *testing.T) {

	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	testId := 3857

	api := starkapi.Client{}
	api.Init(host)
	api.Login(un, pw)

	pointApi := api.PointApi
	point, err := pointApi.GetOne(uint32(testId))

	if err != nil { t.Error("failed", err) }

	if testId != int(point.Id) { t.Fail() }
}

func TestCurValPoint(t *testing.T) {

	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	testId := 3857

	api := starkapi.Client{}
	api.Init(host)
	api.Login(un, pw)

	pointApi := api.PointApi

	for i := 0; i < 10; i++ {
		curVal, err := pointApi.CurVal(uint32(testId))
		if err != nil { t.Error("failed", err) }

		log.Printf("curVal: %v", curVal.Read.Val)
		time.Sleep(2 * time.Second)
	}
}

func TestHisReadPoint(t *testing.T) {

	un := 	os.Getenv(env.STG_SDK_API_UN)
	pw := 	os.Getenv(env.STG_SDK_API_PW)
	host :=	os.Getenv(env.STG_SDK_API_HOST)

	api := starkapi.Client{}
	api.Init(host)
	api.Login(un, pw)

	pointApi := api.PointApi

	pointId, limit, start, end := 3857, 1000, 1614024121, 1614110821
	hisRead, err := pointApi.HisRead(uint32(pointId), uint16(limit), uint64(start), uint64(end))
	if err != nil { t.Error("failed", err) }
	fmt.Printf("count: %x\n", hisRead.Size)

	for _, his := range hisRead.His {
		fmt.Printf("val: %v\n", his.Val)
	}

}

/*
func TestAzure(t *testing.T) {


	az :=	os.Getenv(env.STG_WEATHER_EVENT_HUB_CONN)

	fmt.Printf("az: %s\n",az)
	api := azure.EventHubApi{}
	api.Init(az)


	event := map[string]interface{}{
		"ts": time.Now().Unix(),
		"device_id": "123456789",
		"oaTmpF": 1.0,
		"oaRh": 2.0,
		"oaCond": 3.0,
	}


	var data []byte
	data, err := json.Marshal(event)
	if err != nil { t.Error("failed", err) }

	err = api.Send(data)
	if err != nil { t.Error("failed", err) }



	//fmt.Printf("data: %p\n", &data)
}
*/
