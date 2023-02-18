package spinnaker

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	goHttp "gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

type MyInfo struct {
	MyUrl        string
	SpinnakerUrl string
	UserName     string
	Password     string
}

func (myInfo *MyInfo) pushRelease(branchName string) {
	fs := memfs.New()
	storer := memory.NewStorage()
	r, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: myInfo.MyUrl,
		Auth: &goHttp.BasicAuth{
			Username: myInfo.UserName,
			Password: myInfo.Password,
		},
	})

	CheckIfError(err)

	_, err = r.CreateRemote(&config.RemoteConfig{
		Name: "upstream",
		URLs: []string{myInfo.SpinnakerUrl},
	})

	CheckIfError(err)

	err = r.Fetch(&git.FetchOptions{
		RemoteName: "upstream",
	})

	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewRemoteReferenceName("upstream", branchName),
	})
	CheckIfError(err)

	headRef, err := r.Head()
	CheckIfError(err)

	ref := plumbing.NewHashReference(plumbing.NewBranchReferenceName(branchName), headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
	CheckIfError(err)

	err = w.Checkout(&git.CheckoutOptions{
		Branch: plumbing.NewBranchReferenceName(branchName),
	})
	CheckIfError(err)

	refHead, err := r.Head()
	CheckIfError(err)
	fmt.Println(refHead)

	err = r.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: &goHttp.BasicAuth{
			Username: myInfo.UserName, // anything except an empty string
			Password: myInfo.Password,
		},
	})
	CheckIfError(err)

}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
