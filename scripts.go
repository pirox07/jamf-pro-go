package jamf_pro_go

import (
	"fmt"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

const (
	APIVersionScripts = "v1"
	APIPathScripts = "scripts"
)

type Scripts struct {
	TotalCount uint32	`json:"totalCount"`
	Results []Script	`json:"results"`
}

type Script struct {
	// Script ID
	ID             string `json:"id"`
	// Display name
	Name           string `json:"name"`
	// Information to be displayed to the administrator when the script is run
	Info           string `json:"info"`
	// Notes to be displayed about the script (e.g. author, creation date)
	Notes          string `json:"notes"`
	// Priority to be used for executing scripts related to other actions
	// [ "AFTER", "BEFORE" ]
	Priority       string `json:"priority"`
	// Category to which the script will be added
	CategoryID     string `json:"categoryId"`
	CategoryName   string `json:"categoryName"`
	// Parameter label Names
	Parameter4     string `json:"parameter4"`
	Parameter5     string `json:"parameter5"`
	Parameter6     string `json:"parameter6"`
	Parameter7     string `json:"parameter7"`
	Parameter8     string `json:"parameter8"`
	Parameter9     string `json:"parameter9"`
	Parameter10    string `json:"parameter10"`
	Parameter11    string `json:"parameter11"`
	// Operating system requirements (e.g., "10.6.8, 10.7.x, 10.8")
	OsRequirements string `json:"osRequirements"`
	// Script code
	ScriptContents string `json:"scriptContents"`
}


type GetScriptsOpts struct {
	Page  		uint32		`url:"page,omitempty"`
	PageSize	uint32		`url:"page-size,omitempty"`
	Sort		[]string		`url:"sort,omitempty"`
	Filter		string		`url:"filter,omitempty"`
}

func (c *Client) GetScripts(opts GetScriptsOpts) (*Scripts, error) {
	var result Scripts

	v, err := query.Values(opts)
	if err != nil {
		return nil, err
	}

	err = c.call(path.Join(APIVersionScripts, APIPathScripts), http.MethodGet,
		APIVersionScripts, v, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) GetScript(scriptID uint32) (*Script, error) {
	var result Script

	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodGet,
		APIVersionScripts, nil, nil, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}