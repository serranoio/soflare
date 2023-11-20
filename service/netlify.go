package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Repo struct {
	Branch      string `json:"branch"`
	Cmd         string `json:"cmd"`
	DeployKeyID string `json:"deploy_key_id"`
	Dir         string `json:"dir"`
	Private     bool   `json:"private"`
	Provider    string `json:"provider"`
	Repo        string `json:"repo"`
	RepoID      string `json:"repo_id"`
}

type SiteBody struct {
	Name string `json:"name"`
	Repo Repo   `json:"repo"`
}

type DeployKeyPayload struct {
	ID        string `json:"id"`
	PublicKey string `json:"public_key"`
	CreatedAt string `json:"created_at"`
}

const netlifyToken = "nfp_LH1n3sxRV49m2y18URdzsiMUxnSUBE1N41da"
const githubToken = "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"

func createDeployKey(themeName string) *DeployKeyPayload {
	url := "https://api.netlify.com/api/v1/deploy_keys"
	payload := []byte("")
	authorization := netlifyToken

	deployKeyPayload := &DeployKeyPayload{}

	body := sendPostRequest(url, authorization, payload, "POST")
	defer body.Close()

	err := json.NewDecoder(body).Decode(&deployKeyPayload)

	if err != nil {
		log.Fatal("asda")
	}

	return deployKeyPayload
}

type InsertKeyBody struct {
	Key       string `json:"key"`
	Read_Only bool   `json:"read_only"`
}

func insertKeyIntoRepo(deployKeyPayload *DeployKeyPayload) {
	owner := "serranoio"
	repo := "blog-1"
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/keys", owner, repo)
	insertKeyBody := InsertKeyBody{
		Key:       deployKeyPayload.PublicKey,
		Read_Only: false,
	}

	payload, _ := json.Marshal(insertKeyBody)

	authorization := githubToken

	body := sendPostRequest(url, authorization, payload, "POST")
	defer body.Close()
}

func createSitePayload(themeName string, siteName string) {
	deployKeyPayload := createDeployKey(themeName)
	insertKeyIntoRepo(deployKeyPayload)

	url := "https://api.netlify.com/api/v1/sites"
	authorization := netlifyToken
	siteBody := SiteBody{
		Name: siteName,
		Repo: Repo{
			Branch:      "main",
			Cmd:         "npm run build",
			DeployKeyID: deployKeyPayload.ID,
			Dir:         "build",
			Private:     false,
			Provider:    "github",
			Repo:        "serranoio/blog-1",
			RepoID:      "R_kgDOKokjiA",
		},
	}

	bytes, _ := json.Marshal(siteBody)

	body := sendPostRequest(url, authorization, bytes, "POST")
	defer body.Close()
}

func DeployToNetlify(themeName string, siteName string) error {

	createSitePayload(themeName, siteName)

	return nil
}

func updateFrontend(siteName string, themeName string) {

	url := fmt.Sprintf("https://api.netlify.com/api/v1/sites")
	payload := []byte("")
	authorization := netlifyToken

	body := sendPostRequest(url, authorization, payload, "GET")
	bytes, _ := io.ReadAll(body)
	fmt.Println(bytes)
	defer body.Close()

	url = fmt.Sprintf("https://api.netlify.com/api/v1/sites/%s/deploys", siteName)
	payload = []byte("")
	authorization = netlifyToken

	body = sendPostRequest(url, authorization, payload, "POST")
	defer body.Close()

}
