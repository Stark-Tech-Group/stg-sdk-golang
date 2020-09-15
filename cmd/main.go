package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-scripts/pkg/client"
	"log"
	"os"
	"time"
)

const doDelete = false
const siteId = 900000

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	un := os.Getenv("un")
	pw := os.Getenv("pw")

	fmt.Printf("un: %s, pw: %s", un, pw)

	api := client.Client{}
	api.Init()
	apiStatus := api.ApiStatus()

	fmt.Printf("date: %s\n", time.Unix(apiStatus.DateTime/1000, 0))

	api.Login(un, pw)

	equipSearchBody := client.SearchBody{
		Query:       fmt.Sprintf("equip siteId=%d", siteId),
		CurrentPage: 1,
		PageSize:    1000,
		Sort:        "name",
		Order:       "asc",
	}

	equipResp := api.Search(equipSearchBody)

	for _, equip := range equipResp.Assets {
		fmt.Printf(" Equip: %s, id: %d\n", equip.Name, equip.ID)

		pointSearchBody := client.SearchBody{
			Query:       fmt.Sprintf("point equipId=%d", equip.ID),
			CurrentPage: 1,
			PageSize:    1000,
			Sort:        "name",
			Order:       "asc",
		}

		pointResp := api.Search(pointSearchBody)

		for _, point := range pointResp.Assets {
			fmt.Printf("  Point: %s, id: %d\n", point.Name, point.ID)

			if doDelete {
				pointRes := api.DeletePoint(point.ID)
				time.Sleep(250 * time.Nanosecond)
				fmt.Printf("   Deleted: %s\n", pointRes.Message)
			}
		}

		if doDelete {
			equipDel := api.DeleteEquip(equip.ID)
			fmt.Printf(" Equip Deleted: %s\n", equipDel.Message)
		}
	}

	if doDelete {

		siteDel := api.DeleteSite(siteId)
		fmt.Printf(" Site Deleted: %s", siteDel.Message)
	}

}
