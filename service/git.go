package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func getPublicKey() *ssh.PublicKeys {
	// Get the home directory
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	// Construct the path to the private key
	privateKeyPath := path.Join(homeDir, ".ssh", "id_rsa")

	// Read the private key
	privateKey, err := os.ReadFile(privateKeyPath)
	if err != nil {
		panic(err)
	}

	// Create the ssh auth method
	publicKeys, err := ssh.NewPublicKeys("git", []byte(privateKey), "")
	if err != nil {
		panic(err)
	}

	return publicKeys

}

func makeChange(repoStorage string, name string) {
	pathToApi := path.Join(repoStorage, "src", "model", "api.ts")

	flyApi := fmt.Sprintf("export const Routes = 'https://%s.fly.dev/';  // Hello, World!", name)
	os.WriteFile(pathToApi, []byte(flyApi), 0777)
}

func FetchFrontend(nameOfTheme string, nameOfWebsite string) {

	publicKey := getPublicKey()

	repoStorage := path.Join("tmp", nameOfTheme)

	var repo *git.Repository

	repo, err := git.PlainClone(repoStorage, false, &git.CloneOptions{
		URL:  "git@github.com:serranoio/blog-1.git",
		Auth: publicKey,
	})

	if err != nil {
		repo, err = git.PlainOpen(repoStorage)
	}

	w, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}

	makeChange(repoStorage, nameOfWebsite)

	_, err = w.Add(".")
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Commit(fmt.Sprintf("Website %s added", nameOfWebsite), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "David Serrano",
			Email: "davidserranodev@outlook.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = repo.Push(&git.PushOptions{
		Auth: publicKey,
	})
	if err != nil {
		log.Fatal(err)
	}

}
