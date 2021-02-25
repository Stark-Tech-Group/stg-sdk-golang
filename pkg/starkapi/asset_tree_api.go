package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type AssetTreeApi struct{
	client *Client
}

func (assetTreeApi *AssetTreeApi) baseUrl() string {
	return fmt.Sprintf("%s/core/branches", assetTreeApi.client.host)
}

func (assetTreeApi *AssetTreeApi) adminUrl() string {
	return fmt.Sprintf("%s/admin/branches", assetTreeApi.client.host)
}

func (assetTreeApi *AssetTreeApi) Get() (domain.AssetTree, error) {
	url := fmt.Sprintf("%s/", assetTreeApi.baseUrl())

	var assetTree domain.AssetTree

	resp, err := assetTreeApi.client.get(url)
	if err != nil { return assetTree, err }

	err = json.Unmarshal(resp, &assetTree)
	if err != nil { return assetTree, err }

	return assetTree, nil
}

func (assetTreeApi *AssetTreeApi) GetBranch(id uint32) (domain.Branch, error) {
	url := fmt.Sprintf("%s/%v", assetTreeApi.baseUrl(), id)

	var branch domain.Branch

	resp, err := assetTreeApi.client.get(url)
	if err != nil { return branch, err }

	err = json.Unmarshal(resp, &branch)
	if err != nil { return branch, err }

	return branch, nil
}


func (assetTreeApi *AssetTreeApi) GetChildren(id uint32) (domain.AssetTree, error) {
	url := fmt.Sprintf("%s/%v/children", assetTreeApi.baseUrl(), id)

	var assetTree domain.AssetTree

	resp, err := assetTreeApi.client.get(url)
	if err != nil { return assetTree, err }

	err = json.Unmarshal(resp, &assetTree)
	if err != nil { return assetTree, err }

	return assetTree, nil
}

func (assetTreeApi *AssetTreeApi) GetParents(id uint32) (domain.AssetTree, error) {
	url := fmt.Sprintf("%s/%v/patents", assetTreeApi.baseUrl(), id)

	var assetTree domain.AssetTree

	resp, err := assetTreeApi.client.get(url)
	if err != nil { return assetTree, err }

	err = json.Unmarshal(resp, &assetTree)
	if err != nil { return assetTree, err }

	return assetTree, nil
}