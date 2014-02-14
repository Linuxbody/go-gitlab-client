package gogitlab

import (
	"encoding/json"
	"net/url"
	"strings"
)

const (
	// ID
	project_url_deploy_keys = "/projects/:id/keys" // Get list of project deploy keys
	// PROJECT ID AND KEY ID
	project_url_deploy_key = "/projects/:id/keys/:key_id" // Get single project deploy key
)

/*
Get list of project deploy keys
*/
func (g *Gitlab) ProjectDeployKeys(id string) ([]*DeployKey, error) {

	url := strings.Replace(project_url_deploy_keys, ":id", id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var deployKeys []*DeployKey

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err != nil {
		return deployKeys, err
	}

	err = json.Unmarshal(contents, &deployKeys)

	return deployKeys, err
}

/*
Get single project deploy key
*/
func (g *Gitlab) ProjectDeployKey(id, key_id string) (*DeployKey, error) {

	url := strings.Replace(project_url_deploy_key, ":id", id, -1)
	url = strings.Replace(url, ":key_id", key_id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error
	var deployKey *DeployKey

	contents, err := g.buildAndExecRequest("GET", url, nil)
	if err != nil {
		return deployKey, err
	}

	err = json.Unmarshal(contents, &deployKey)

	return deployKey, err
}

/*
Add deploy key to project
*/
func (g *Gitlab) AddProjectDeployKey(id, title, key string) error {

	path := strings.Replace(project_url_deploy_keys, ":id", id, -1)
	path = g.BaseUrl + g.ApiPath + path + "?private_token=" + g.Token

	var err error

	v := url.Values{}
	v.Set("title", title)
	v.Set("key", key)

	body := v.Encode()

	_, err = g.buildAndExecRequest("POST", path, []byte(body))

	return err
}

/*
Remove deploy key from project
*/
func (g *Gitlab) RemoveProjectDeployKey(id, key_id string) error {

	url := strings.Replace(project_url_deploy_key, ":id", id, -1)
	url = strings.Replace(url, ":key_id", key_id, -1)
	url = g.BaseUrl + g.ApiPath + url + "?private_token=" + g.Token

	var err error

	_, err = g.buildAndExecRequest("DELETE", url, nil)

	return err
}