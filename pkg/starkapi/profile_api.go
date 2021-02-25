package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type ProfileApi struct{
	client *Client
}

func (profileApi *ProfileApi) host() string {
	return profileApi.client.host
}

func (profileApi *ProfileApi) url() string {
	return fmt.Sprintf("%s/core/profiles", profileApi.client.host)
}


func (profileApi *ProfileApi) GetOne(id uint32) (domain.Profile, error) {
	var profile domain.Profile

	resp, err := profileApi.client.get(fmt.Sprintf("%s/%d", profileApi.url(), id))

	if err != nil { return profile, err }

	err = json.Unmarshal(resp, &profile)
	if err != nil{
		return profile, err
	}

	return profile, nil
}

func (profileApi *ProfileApi) GetAll() (domain.Profiles, error) {
	var profiles domain.Profiles
	url := profileApi.url() + "/"
	resp, err := profileApi.client.get(url)

	if err != nil { return profiles, err }

	err = json.Unmarshal(resp, &profiles)

	if err != nil{
		return profiles, err
	}

	return profiles, nil
}
