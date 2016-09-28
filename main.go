/**
* author : gzc
* time   : 28/09/2016 11:15
*/
package main

import (
	"flag"
	"fmt"
	"github.com/com314159/godoc_to_markdown_wiki/docparse/tomlparser"
	"github.com/com314159/godoc_to_markdown_wiki/docparse"
)

var tomlFile = flag.String("toml","","the toml file path")

func main()  {
	
	flag.Parse()
	
	fmt.Println("tomlFile: ",*tomlFile)
	
	if *tomlFile == "" {
		fmt.Println("need toml file path")
		return
	}
	
	err := tomlparser.ParseToml(*tomlFile)
	if err != nil {
		fmt.Println(" parse toml file error ")
	}
	
	if tomlparser.WikiOutDir != "" {
		docparse.ParseToOutDir(tomlparser.Pages,tomlparser.WikiOutDir)
	} else if tomlparser.WikiPackageName != "" {
		docparse.ParseToPackageName(tomlparser.Pages,tomlparser.WikiPackageName)
	} else {
		fmt.Println("must set out dir")
	}
		
	
}
