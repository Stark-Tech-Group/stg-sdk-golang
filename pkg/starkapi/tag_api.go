package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type TagApi struct {
	client *Client
}

type tagSuggestion struct {
	Context     string `json:"context"`
	Kind        string `json:"type"`
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"desc"`
}

func (tagApi *TagApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/tags", tagApi.client.host)
}

func (tagApi *TagApi) host() string {
	return tagApi.client.host
}

func (tagApi *TagApi) Suggest(query string, context string) ([]domain.Tag, error) {
	tags := make([]domain.Tag, 0)
	url := fmt.Sprintf(tagApi.BaseUrl(), "/suggest")

	data := struct {
		Query   string `json:"query"`
		Context string `json:"context"`
	}{Query: query, Context: context}

	body, err := json.Marshal(data)
	if err != nil {
		return tags, err
	}

	resp, err := tagApi.client.post(url, body)
	if err != nil {
		return tags, err
	}

	var suggestions []tagSuggestion
	err = json.Unmarshal(resp, &suggestions)
	if err != nil {
		return tags, err
	}

	for _, suggestion := range suggestions {
		tags = append(tags, domain.Tag{
			Name:  suggestion.Key,
			Value: suggestion.Value,
		})
	}

	return tags, nil

}
