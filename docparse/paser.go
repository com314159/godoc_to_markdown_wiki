/**
* author : gzc
* time   : 28/09/2016 10:51
*/
package docparse

import (
	"github.com/com314159/godoc_to_markdown_wiki/docparse/tomlparser"
	"os"
	"strings"
	"runtime"
	"path"
	"fmt"
	"go/token"
	"go/parser"
	"go/ast"
	"bytes"
)

func ParseToPackageName(pages []tomlparser.WikiPage ,packageDir string)  {
	
	gopath := os.Getenv("GOPATH")
	dirs := strings.Split(gopath,":")
	
	if runtime.GOOS == "windows" {
		dirs = strings.Split(gopath, ";")
	}
	
	
	for _,page := range pages {
		
		files := []string{}
		outFile := ""
		
		for _, d := range dirs {
			for _, f := range page.ApiFiles {
				goFile := path.Join(d,"src",f)
				
				if _,err := os.Stat(goFile); err == nil {
					if outFile == "" {
						outFile = path.Join(d,"src",packageDir,page.PageName+".markdown")
					}
					files = append(files,goFile)
				}
			}
			
		}
		
		parseFile(files,page.FuncNames,outFile)
		
	}
	
}



func ParseToOutDir(pages []tomlparser.WikiPage, outDir string)  {
	
	gopath := os.Getenv("GOPATH")
	dirs := strings.Split(gopath,":")
	
	if runtime.GOOS == "windows" {
		dirs = strings.Split(gopath, ";")
	}
	
	for _, d := range dirs {
		for _,page := range pages {
			files := []string{}
			outFile := path.Join(outDir,page.PageName+".markdown")
			for _, f :=range page.ApiFiles {
				goFile := path.Join(d,"src",f)
				
				if _,err := os.Stat(goFile); err == nil {
					files = append(files,goFile)
				}
			}
			
			parseFile(files,page.FuncNames,outFile)
		}
	}
	
}



func parseFile(goFiles []string, funcNames []string,outFile string)  {
	
	fmt.Println("start parse: ",goFiles)
	fmt.Println("start parse outfile: ",outFile)
	var bytesBuffer bytes.Buffer
	
	for _, goFile:= range goFiles {
			
		fileSet := token.NewFileSet()
		fileTree, err := parser.ParseFile(fileSet,goFile,nil,parser.ParseComments)
		
		if err != nil {
			fmt.Println("error",err)
			return
		}
		
		if fileTree.Decls == nil {
			fmt.Println("decl is nil")
			return
		}
		
		for _, dec := range fileTree.Decls{
			if fun,ok := dec.(*ast.FuncDecl);ok {
				
				for _, funcName := range funcNames {
					if funcName == fun.Name.Name {
						bytesBuffer.WriteString(fun.Doc.Text())
						bytesBuffer.WriteString("\n")
					}
				}
				
			}
		}
	}
	
	
	if bytesBuffer.Len() == 0 {
		fmt.Println("bytes is 0")
		return
	} else {
		fmt.Println("write :\n",bytesBuffer.String())
	}
	
	if _, err := os.Stat(outFile); err == nil {
		fmt.Println("file exist delete",outFile)
		os.Remove(outFile)
	}
	
	parentDir := GetParentDirectory(outFile)
	
	fmt.Println("parent dir ",parentDir)
	fmt.Println("file ",outFile)
	
	err := os.MkdirAll(parentDir,0777)
	if err != nil {
		fmt.Println("can't create dir: ",err)
		return
	}
		
	fd, err := os.Create(outFile)
	if err != nil {
		fmt.Println("Can not create document file: \n", err)
		return
	}
	defer fd.Close()
	fd.Write(bytesBuffer.Bytes())
	fmt.Println("write file:",outFile)
}


func GetParentDirectory(file string) string {
	return string(file[:strings.LastIndex(file,"/")])
}