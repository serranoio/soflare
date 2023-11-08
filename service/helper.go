package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func sendPostRequest(url string, authorization string, payload []byte) io.ReadCloser {

	r, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		panic(err)
	}

	// r.Header.Add("Accept", "application/vnd.github+json")
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authorization))

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		panic(err)
	}

	return res.Body
}
