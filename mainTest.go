/**
* author : gzc
* time   : 28/09/2016 08:36
*/
package main

import (
	"os"
	"strings"
	"path"
	"fmt"
	"go/token"
	goparser "go/parser"
	"go/ast"
)

var pkgNameFile string = "demos/gotest/api/api.go"

func main()  {
	
	gopath := os.Getenv("GOPATH")
	dirs := strings.Split(gopath,":")
	
	for _, d := range dirs {
		apifile := path.Join(d,"src",pkgNameFile)
		fmt.Println("apiFile",apifile)
		if _,err := os.Stat(apifile); err == nil {
			parseComment(apifile)
		}
	}
		
}

func parseComment(apiFile string)  {
	
	fmt.Println("start parse",apiFile)
	
	fileSet := token.NewFileSet()
	fileTree, err := goparser.ParseFile(fileSet,apiFile,nil,goparser.ParseComments)
	
	if err != nil {
		fmt.Println("error",err)
		return
	}
	
	if fileTree.Comments != nil {
		for _, comment := range fileTree.Comments {
			fmt.Println("comment: ",comment.Text())
		}
	} else {
		fmt.Println("comment nil")
	}
	
	if fileTree.Doc != nil {
		fmt.Println("doc: ",fileTree.Doc.Text())
	} else {
		fmt.Println("doc nil")
	}
	
	if fileTree.Decls != nil {
		for _, dec := range fileTree.Decls {
			
			if fun,ok := dec.(*ast.FuncDecl);ok {
				fmt.Println("funcDecl")
				fmt.Println("funcName: ",fun.Name)
				fmt.Println("commentText: ",fun.Doc.Text())
			}
			
		}
	} else {
		fmt.Println("decl nil")
	}
}