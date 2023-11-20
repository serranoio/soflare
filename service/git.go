package main

import (
	"fmt"
	"os"
	"os/exec"
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

func deleteFolder(nameOfTheme string) {
	cmd := exec.Command("rm", "-rf", nameOfTheme)
	// dir, _ := os.Getwd()
	// cmd.Dir = path.Join(dir, "tmp", nameOfTheme)

	if err := cmd.Run(); err != nil {

	}
}

func FetchFrontend(nameOfTheme string, nameOfWebsite string) error {
	deleteFolder(nameOfTheme)

	publicKey := getPublicKey()

	repoStorage := path.Join("tmp", nameOfTheme)

	var repo *git.Repository

	repo, err := git.PlainClone(repoStorage, false, &git.CloneOptions{
		URL:  "git@github.com:serranoio/blog-1.git",
		Auth: publicKey,
	})

	if err != nil {
		repo, err = git.PlainOpen(repoStorage)

		// repo.Fetch(&git.FetchOptions{})

		// repo.Pull(&git.FetchOptions{})
	}

	w, err := repo.Worktree()
	if err != nil {
		return err
	}

	err = w.Pull(&git.PullOptions{
		Auth:       publicKey,
		RemoteName: "origin",
	})

	if err.Error() == "already up-to-date" {
		err = nil
	}

	if err != nil {
		return err
	}

	makeChange(repoStorage, nameOfWebsite)

	_, err = w.Add(".")
	if err != nil {
		return err
	}

	_, err = w.Commit(fmt.Sprintf("Website %s added", nameOfWebsite), &git.CommitOptions{
		Author: &object.Signature{
			Name:  "David Serrano",
			Email: "davidserranodev@outlook.com",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	err = repo.Push(&git.PushOptions{
		Auth: publicKey,
	})
	if err != nil {
		return err
	}

	return nil
}
