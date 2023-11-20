package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path"
)

const token = "fo1_fpUXhMak4p853N6DdokPJ3eSqEayKmw6xfDQPOFA17Y"
const endpoint = "https://api.machines.dev"

type AppBody struct {
	AppName string `json:"app_name"`
	OrgSlug string `json:"org_slug"`
}

func createApp(name string) error {
	url := fmt.Sprintf("%s/v1/apps", endpoint)

	appBody := AppBody{
		AppName: name,
		OrgSlug: "personal",
	}

	payload, _ := json.Marshal(appBody)

	body := sendPostRequest(url, token, payload, "POST")
	defer body.Close()

	return nil
}

func writeToml(name string) string {
	wd, _ := os.Getwd()
	backendTomlPath := path.Join(wd, "../", "backend", "fly.toml")

	newConfiguration := fmt.Sprintf(`
	# See https://fly.io/docs/reference/configuration/ for information about how to use this file.
	#
	
	app = "%s"
	primary_region = "ord"
	
	[build]


	[http_service]
	internal_port = 8080
	force_https = true
	auto_stop_machines = true
	auto_start_machines = true
	min_machines_running = 0
	processes = ["app"]
	
	[mounts]
	source="mount"
	destination="/app/pb_data"
	`, name)

	os.WriteFile(backendTomlPath, []byte(newConfiguration), 0777)

	return backendTomlPath

}

func deployMachine(name string) error {
	backendTomlPath := writeToml(name)

	cmd := exec.Command("fly", "deploy")
	cmd.Dir = path.Join(backendTomlPath, "../")
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func DeployFly(name string) error {

	err := createApp(name)
	if err != nil {
		return err
	}
	err = deployMachine(name)

	return err
}
