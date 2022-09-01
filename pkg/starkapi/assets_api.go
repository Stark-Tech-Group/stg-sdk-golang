package starkapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
	logger "github.com/sirupsen/logrus"
)

type AssetsApi struct {
	client *Client
}

func assetsUrl(host string) string {
	return fmt.Sprintf("%s/core/assets", host)
}

func assetUrl(host string, ref uint32) string {
	return fmt.Sprintf("%s/%d", assetsUrl(host), ref)
}

func (assetsApi *AssetsApi) host() string {
	return assetsApi.client.host
}

func (assetsApi *AssetsApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/assets", assetsApi.client.host)
}

func (assetsApi *AssetsApi) AddNewTag(asset domain.Asset, name string, value string) error {
	if len(name) < 1 {
		logger.Error("invalid tag name in assets/AddNewTag.")
		return errors.New("invalid tag name")
	}

	if len(asset.Ref) < 1 {
		logger.Error("invalid asset in assets/AddNewTag. Asset does not have a ref.")
		return errors.New("invalid asset")
	}

	url := fmt.Sprintf("%s/%v/tags", assetsApi.BaseUrl(), asset.Ref)
	ask := domain.Tag{Name: name, Value: value}

	body, err := json.Marshal(ask)
	if err != nil {
		return err
	}

	_, err = assetsApi.client.post(url, body)
	if err != nil {
		return err
	}

	return nil
}
