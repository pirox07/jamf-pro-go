package jamf_pro_go

import (
	"net/http"
)

const (
	APIVersionPolicies	= "classic"
	APIPathPolices		= "policies"
)

//type GetPoliciesResult struct {
type Policies struct {
	Size uint32									`xml:"size"`
	Policy []GetPoliciesResultPolicyOverview 	`xml:"policy"`
}

type GetPoliciesResultPolicyOverview struct {
	ID 		uint32 `xml:"id"`
	Name	string `xml:"name"`
}

func (c *Client) GetPolicies() (*Policies, error) {
	var result Policies

	err := c.call(APIPathPolices, http.MethodGet,
		APIVersionPolicies, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}