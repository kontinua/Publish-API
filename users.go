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


func PrintHugeParams(fset *token.FileSet, info *types.Info, files []*ast.File) {
	checkTuple := func(descr string, tuple *types.Tuple) {
		for i := 0; i < tuple.Len(); i++ {
			v := tuple.At(i)
			if sz := sizeof(v.Type()); sz > int64(*bytesFlag) {
				fmt.Printf("%s: %q %s: %s = %d bytes\n",
					fset.Position(v.Pos()),
					v.Name(), descr, v.Type(), sz)
			}
		}
	}
	checkSig := func(sig *types.Signature) {
		checkTuple("parameter", sig.Params())
		checkTuple("result", sig.Results())
	}
	for _, file := range files {
		ast.Inspect(file, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncDecl:
				checkSig(info.Defs[n.Name].Type().(*types.Signature))
			case *ast.FuncLit:
				checkSig(info.Types[n.Type].Type.(*types.Signature))
			}
			return true
		})
	}
}
