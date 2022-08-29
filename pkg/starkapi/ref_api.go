package starkapi

import "fmt"

type RefApi struct{
	client *Client

}

func (refApi *RefApi) baseUrl() string {
	return fmt.Sprintf("%s/core/branches", refApi.client.host)
}

func (refApi *RefApi) adminUrl() string {
	return fmt.Sprintf("%s/admin/branches", refApi.client.host)
}
