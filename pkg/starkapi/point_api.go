package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/api/response"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type PointApi struct {
	client ApiClient
}

func (pointApi *PointApi) host() string {
	return pointApi.client.getHost()
}

func (pointApi *PointApi) delete(id uint32) (*response.DeleteResponse, error) {

	resp, err := pointApi.client.delete(pointUrl(pointApi.host(), id))

	if err != nil {
		return nil, err
	}

	deleteResp := response.DeleteResponse{}

	err = json.Unmarshal(resp, &deleteResp)

	if err != nil {
		return nil, err
	}

	return &deleteResp, nil
}

func pointsUrl(host string) string {
	return fmt.Sprintf("%s/core/points", host)
}

func pointUrl(host string, id uint32) string {
	return fmt.Sprintf("%s/%d", pointsUrl(host), id)
}

func (pointApi *PointApi) ListPointTypeUrl() string {
	return fmt.Sprintf("%s/core/lists/pointType", pointApi.host())
}

func (pointApi *PointApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/points", pointApi.host())
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

func (pointApi *PointApi) UpdateOne(id uint32, jsonBody []byte) (domain.Point, error) {
	url := fmt.Sprintf("%s/%v", pointApi.BaseUrl(), id)

	var point domain.Point
	resp, err := pointApi.client.put(url, jsonBody)
	if err != nil {
		return point, err
	}

	err = json.Unmarshal(resp, &point)
	if err != nil {
		return point, err
	}

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

	return nil
}

// GetAllTags returns all tags for the provided domain.Point
func (pointApi *PointApi) GetAllTags(point domain.Point) (domain.TagRefs, error) {
	url := fmt.Sprintf("%s/%v/tags", pointApi.BaseUrl(), point.Id)

	var tags domain.TagRefs
	resp, err := pointApi.client.get(url)
	if err != nil {
		return tags, err
	}

	err = json.Unmarshal(resp, &tags)
	if err != nil {
		return tags, err
	}

	return tags, nil
}

// DeleteTag deletes a domain.TagRef from the provided domain.Point
func (pointApi *PointApi) DeleteTag(point domain.Point, tagRef domain.TagRef) error {
	url := fmt.Sprintf("%s/%v/tags/%v", pointApi.BaseUrl(), point.Id, tagRef.Id)

	_, err := pointApi.client.delete(url)
	if err != nil {
		return err
	}

	return nil
}

func (pointApi *PointApi) GetOne(id uint32) (domain.Point, error) {
	url := fmt.Sprintf("%s/%v", pointApi.BaseUrl(), id)

	var point domain.Point

	resp, err := pointApi.client.get(url)
	if err != nil {
		return point, err
	}

	err = json.Unmarshal(resp, &point)
	if err != nil {
		return point, err
	}

	return point, nil
}

// GetAll returns all points within the given limit and offset
func (pointApi *PointApi) GetAll(limit, offset int) (domain.Points, error) {
	url := fmt.Sprintf("%s?limit=%v?offset=%v", pointApi.BaseUrl(), limit, offset)
	var points domain.Points

	resp, err := pointApi.client.get(url)
	if err != nil {
		return points, err
	}
	err = json.Unmarshal(resp, &points)
	if err != nil {
		return points, err
	}
	return points, nil
}

func (pointApi *PointApi) GetAllByRef(ref string) (domain.Points, error) {

	url := fmt.Sprintf("%s?parentRef=%v", pointApi.BaseUrl(), ref)

	var points domain.Points

	resp, err := pointApi.client.get(url)
	if err != nil {
		return points, err
	}
	err = json.Unmarshal(resp, &points)
	if err != nil {
		return points, err
	}
	return points, nil
}

func (pointApi *PointApi) CurVal(id uint32) (domain.CurVal, error) {
	url := fmt.Sprintf("%s/%v/%s", pointApi.BaseUrl(), id, "curVal")

	var curVal domain.CurVal

	resp, err := pointApi.client.get(url)
	if err != nil {
		return curVal, err
	}

	err = json.Unmarshal(resp, &curVal)
	if err != nil {
		return curVal, err
	}

	return curVal, nil
}

func (pointApi *PointApi) HisRead(id uint32, limit uint16, start uint64, end uint64) (domain.HisRead, error) {
	url := fmt.Sprintf("%s/%v/%s?limit=%v&start=%v&end=%v", pointApi.BaseUrl(), id, "hisRead", limit, start, end)

	var hisRead domain.HisRead

	resp, err := pointApi.client.get(url)
	if err != nil {
		return hisRead, err
	}

	err = json.Unmarshal(resp, &hisRead)
	if err != nil {
		return hisRead, err
	}

	return hisRead, nil
}

func (pointApi *PointApi) GetAllPointTypes() (domain.PointTypes, error) {
	url := fmt.Sprintf("%s", pointApi.ListPointTypeUrl())

	var pointTypes domain.PointTypes

	resp, err := pointApi.client.get(url)
	if err != nil {
		return pointTypes, err
	}

	if len(resp) > 0 {
		err = json.Unmarshal(resp, &pointTypes)
		if err != nil {
			return pointTypes, err
		}
	}

	return pointTypes, nil
}
