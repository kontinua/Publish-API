package main

import (
    "github.com/kataras/iris"
    "github.com/iris-contrib/plugin/typescript"
)

func main(){
    /* Options
    Bin ->            the typescript installation path/bin/tsc or tsc.cmd, if empty then it will search to the global npm modules
    Dir    ->          where to search for typescript files/project. Default "./"
    Ignore ->        comma separated ignore typescript files/project from these directories (/node_modules/ are always ignored). Default ""
  Tsconfig ->       &typescript.Tsconfig{}, here you can set all compilerOptions if no tsconfig.json exists inside the 'Dir'
  Editor ->         typescript.Editor(), if setted then alm-tools browser-based typescript IDE will be available. Default is nil.
    */

    config := typescript.Config {
        Dir: "./scripts/src",
        Tsconfig: &typescript.Tsconfig{Module: "commonjs", Target: "es5"}, // or typescript.DefaultTsconfig()
    }

    //if you want to change only certain option(s) but you want default to all others then you have to do this:
    config = typescript.DefaultConfig()
    //

    iris.Plugins.Add(typescript.New(config)) //or with the default options just: typescript.New()

    iris.Get("/", func (ctx *iris.Context){})

    iris.Listen(":8080")
}
