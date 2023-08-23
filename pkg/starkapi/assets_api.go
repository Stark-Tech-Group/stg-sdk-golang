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

const invalidTagNameErrorStr = "invalid tag name"
const invalidAssetErrorStr = "invalid asset"

func (assetsApi *AssetsApi) host() string {
	return assetsApi.client.host
}

func (assetsApi *AssetsApi) baseUrl() string {
	return fmt.Sprintf("%s/core/assets", assetsApi.host())
}

func (assetsApi *AssetsApi) tagUrlWithRef(ref string) string {
	return fmt.Sprintf("%s/%s/tags", assetsApi.baseUrl(), ref)
}

func (assetsApi *AssetsApi) AddNewTag(asset domain.Asset, tagName string, tagValue string) error {
	if len(tagName) < 1 {
		logger.Error("invalid tag name in assets/AddNewTag.")
		return errors.New(invalidTagNameErrorStr)
	}

	if len(asset.Ref) < 1 {
		logger.Error("invalid asset in assets/AddNewTag. Asset does not have a ref.")
		return errors.New(invalidAssetErrorStr)
	}

	url := assetsApi.tagUrlWithRef(asset.Ref)
	ask := domain.Tag{Name: tagName, Value: tagValue}

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

func (assetsApi *AssetsApi) AddNewTags(asset domain.Asset, tags []domain.Tag) error {
	if len(tags) < 1 {
		logger.Error("please provide at least one tag/AddNewTags.")
		return errors.New("please provide at least one tag")
	}

	if len(asset.Ref) < 1 {
		logger.Error("invalid asset in assets/AddNewTag. Asset does not have a ref.")
		return errors.New(invalidAssetErrorStr)
	}

	for _, element := range tags {
		if len(element.Name) < 1 {
			logger.Error("invalid tag name in assets/AddNewTags.")
			return errors.New(invalidTagNameErrorStr)
		}
	}

	url := assetsApi.tagUrlWithRef(asset.Ref)

	body, err := json.Marshal(tags)
	if err != nil {
		return err
	}

	_, err = assetsApi.client.post(url, body)
	if err != nil {
		return err
	}

	return nil
}

func (assetsApi *AssetsApi) DeleteTag(asset domain.Asset, tagName string) error {
	if len(tagName) < 1 {
		logger.Error("invalid tag name in assets/DeleteTag.")
		return errors.New(invalidTagNameErrorStr)
	}

	if len(asset.Ref) < 1 {
		logger.Error("invalid asset in assets/DeleteTag. Asset does not have a ref.")
		return errors.New(invalidAssetErrorStr)
	}

	url := fmt.Sprintf("%s/%s", assetsApi.tagUrlWithRef(asset.Ref), tagName)

	_, err := assetsApi.client.delete(url)
	if err != nil {
		return err
	}

	return nil
}

func (assetsApi *AssetsApi) CreateAuditLog(asset domain.Asset, log domain.AuditLog) error {
	if len(asset.Ref) < 1 {
		logger.Error("invalid asset in assets/CreateAuditLog. Asset does not have a ref.")
		return errors.New(invalidAssetErrorStr)
	}

	data, err := json.Marshal(log)
	if err != nil {
		logger.Errorf("failed to marshal log with error : [%s]", err)
		return err
	}

	url := fmt.Sprintf("%s/%s", assetsApi.baseUrl(), asset.Ref)
	_, err = assetsApi.client.post(url, data)
	if err != nil {
		return err
	}

	return nil
}
