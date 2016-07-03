package main

import "github.com/google/go-github/github"

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

func reverse() {
	fmt.Println(stringutil.Reverse("!selpmaxe oG ,olleH"))
}

// init is run before the application starts serving.
func init() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/hello", helloHandler)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from the Go app")
}
