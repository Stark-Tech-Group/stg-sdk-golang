package starkapi

import (
	"encoding/json"
	"fmt"
	"github.com/Stark-Tech-Group/stg-sdk-golang/pkg/domain"
)

type ConnApi struct{
	client *Client
}

func (connApi *ConnApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/conns", connApi.client.host)
}

func (connApi *ConnApi) GetOne(id int) (domain.Conn, error) {
	var conn domain.Conn

	resp, err := connApi.client.get(siteUrl(connApi.client.host, id))
	if err != nil { return conn, err }

	err = json.Unmarshal(resp, &conn)
	if err != nil { return conn, err }

	return conn, nil
}

func (connApi *ConnApi) GetAll() (domain.Conns, error) {
	var conns domain.Conns
	url := connApi.BaseUrl() + "/"

	resp, err := connApi.client.get(url)
	if err != nil { return conns, err }

	err = json.Unmarshal(resp, &conns)
	if err != nil { return conns, err }

	return conns, nil
}