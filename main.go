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
	"strings"
	"os"
	"runtime"
	"path"
)

var tomlFile = flag.String("toml","","the toml file path")

func main()  {
	
	flag.Parse()
	
	if *tomlFile == "" {
		fmt.Println("need toml file path")
		return
	}
	
	if strings.HasPrefix(*tomlFile,"$GOPATH") {
		
		gopath := os.Getenv("GOPATH")
		dirs := strings.Split(gopath,":")
		
		if runtime.GOOS == "windows" {
			dirs = strings.Split(gopath, ";")
		}
		
		for _,d := range dirs {
			filePath := strings.Replace(*tomlFile,"$GOPATH",path.Join(d,"src"),1)
			if _,err := os.Stat(filePath); err == nil {
				*tomlFile = filePath
				break
			}
		}
	}
	
	fmt.Println("tomlFile: ",*tomlFile)
	
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
