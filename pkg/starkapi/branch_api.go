package starkapi

import (
	"fmt"
)

type BranchApi struct{
	client *Client
}

func (branchApi *BranchApi) BaseUrl() string {
	return fmt.Sprintf("%s/core/branches", branchApi.client.host)
}

func (branchApi *BranchApi) AdminUrl() string {
	return fmt.Sprintf("%s/admin/branches", branchApi.client.host)
}



