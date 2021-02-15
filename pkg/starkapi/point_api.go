package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type PointApi struct{
	client *Client
}

func (pointApi *PointApi) delete(id int) (*response.DeleteResponse, error){

	resp, err := pointApi.client.delete(pointUrl(pointApi.client.host, id))

	if err != nil{
		return nil, err
	}

	deleteResp := response.DeleteResponse{}

	err = json.Unmarshal(resp, &deleteResp)

	if err != nil{
		panic(err)
	}

	return &deleteResp, nil
}

func pointsUrl(host string) string {
	return fmt.Sprintf("%s/core/points",host)
}

func pointUrl(host string , id int) string {
	return fmt.Sprintf("%s/%d", pointsUrl(host), id)
}

func (pointApi *PointApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/points", pointApi.client.host)
}

func (pointApi *PointApi) CreateOne(ask domain.Point) (domain.Point, error) {
	url := pointApi.BaseUrl()
	body, err := json.Marshal(ask)

	var point domain.Point
	if err != nil {
		return point, err
	}

	resp, err := pointApi.client.post(url, body)
	if err != nil {
		return point, err
	}
	err = json.Unmarshal(resp, &point)
	if err != nil {
		return point, err
	}

	return point, nil
}

func (pointApi *PointApi) UpdateOne(id int, jsonBody []byte)(domain.Point, error) {
	url := fmt.Sprintf("%s/%v", pointApi.BaseUrl(), id)

	var point domain.Point
	resp, err := pointApi.client.put(url, jsonBody)
	if err != nil { return point, err }

	err = json.Unmarshal(resp, &point)
	if err != nil { return point, err }

	return point, nil
}

func (pointApi *PointApi) AddNewTag(point domain.Point, name string, value string) error {
	url := fmt.Sprintf("%s/%v/tags", pointApi.BaseUrl(), point.Id)
	ask := domain.Tag{Name: name, Value: value}

	body, err := json.Marshal(ask)
	if err != nil {
		return err
	}

	_, err = pointApi.client.post(url, body)
	if err != nil {
		return err
	}

	return  nil
}