package main

import (
	"github.com/google/go-github/github"
)

client := github.NewClient(nil)
orgs, _, err := client.Organizations.List("kontinua", nil)

func main() {
  ts := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: "... your access token ..."},
  )
  tc := oauth2.NewClient(oauth2.NoContext, ts)

  client := github.NewClient(tc)

  // list all repositories for the authenticated user
  repos, _, err := client.Repositories.List("", nil)
}
