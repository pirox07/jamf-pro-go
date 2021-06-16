package jamf_pro_go

import (
	"fmt"
	"io"
	"net/http"
	"path"

	"github.com/google/go-querystring/query"
)

const (
	APIVersionScripts = "v1"
	APIPathScripts    = "scripts"
)

type Scripts struct {
	TotalCount  uint32   `json:"totalCount"`
	Results     []Script `json:"results"`
}

type Script struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Info            string `json:"info"`
	Notes           string `json:"notes"`
	Priority        string `json:"priority"` // [ BEFORE, AFTER, AT_REBOOT ] (default: BEFORE)
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	Parameter4      string `json:"parameter4"`
	Parameter5      string `json:"parameter5"`
	Parameter6      string `json:"parameter6"`
	Parameter7      string `json:"parameter7"`
	Parameter8      string `json:"parameter8"`
	Parameter9      string `json:"parameter9"`
	Parameter10     string `json:"parameter10"`
	Parameter11     string `json:"parameter11"`
	OsRequirements  string `json:"osRequirements"`
	ScriptContents  string `json:"scriptContents"`
}


type GetScriptsOpts struct {
	Page      uint32   `url:"page,omitempty"`
	PageSize  uint32   `url:"page-size,omitempty"`
	Sort      []string `url:"sort,omitempty"`
	Filter    string   `url:"filter,omitempty"`
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

type CreateScriptResult struct {
	ID    string `json:"id"`
	Href  string `json:"href"`
}

type CreateScriptParams struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name"`
	Info            string `json:"info,omitempty"`
	Notes           string `json:"notes,omitempty"`
	Priority        string `json:"priority,omitempty"` // [ BEFORE, AFTER, AT_REBOOT ] (default: BEFORE)
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName,omitempty"`
	Parameter4      string `json:"parameter4,omitempty"`
	Parameter5      string `json:"parameter5,omitempty"`
	Parameter6      string `json:"parameter6,omitempty"`
	Parameter7      string `json:"parameter7,omitempty"`
	Parameter8      string `json:"parameter8,omitempty"`
	Parameter9      string `json:"parameter9,omitempty"`
	Parameter10     string `json:"parameter10,omitempty"`
	Parameter11     string `json:"parameter11,omitempty"`
	OsRequirements  string `json:"osRequirements,omitempty"`
	ScriptContents  string `json:"scriptContents,omitempty"`
}

func (c *Client) CreateScript (params CreateScriptParams) (*CreateScriptResult, error) {
	var result CreateScriptResult

	err := c.call(path.Join(APIVersionScripts, APIPathScripts), http.MethodPost,
		APIVersionScripts, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}


type UpdateScriptParams struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Info            string `json:"info,omitempty"`
	Notes           string `json:"notes,omitempty"`
	Priority        string `json:"priority,omitempty"` // [ BEFORE, AFTER, AT_REBOOT ] (default: BEFORE)
	CategoryID      string `json:"categoryId"`
	CategoryName    string `json:"categoryName,omitempty"`
	Parameter4      string `json:"parameter4,omitempty"`
	Parameter5      string `json:"parameter5,omitempty"`
	Parameter6      string `json:"parameter6,omitempty"`
	Parameter7      string `json:"parameter7,omitempty"`
	Parameter8      string `json:"parameter8,omitempty"`
	Parameter9      string `json:"parameter9,omitempty"`
	Parameter10     string `json:"parameter10,omitempty"`
	Parameter11     string `json:"parameter11,omitempty"`
	OsRequirements  string `json:"osRequirements,omitempty"`
	ScriptContents  string `json:"scriptContents,omitempty"`
}

func (c *Client) UpdateScript (scriptID uint32, params UpdateScriptParams) (*Script, error) {
	var result Script

	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodPut,
		APIVersionScripts, nil, params, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) DeleteScript (scriptID uint32) error {
	err := c.call(path.Join(APIVersionScripts, APIPathScripts, fmt.Sprint(scriptID)), http.MethodDelete,
		APIVersionScripts, nil, nil, nil)
	if err != io.EOF {
		return err
	}
	fmt.Println("[jamf-pro-go] Script (ID: ", scriptID, ") is deleted")

	return nil
}