/**
* author : gzc
* time   : 28/09/2016 10:52
*/
package tomlparser

import (
	"fmt"
	"github.com/spf13/viper"
)

type WikiPage struct {
	ApiFile string
	FuncNames []string
	PageName string
}

var Pages []WikiPage = []WikiPage{}
var WikiPackageName string
var WikiOutDir string

func ParseToml(tomlfile string) error {
	
	fmt.Println("start parse toml file ",tomlfile)
	viper.SetConfigFile(tomlfile)
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read toml file err")
		fmt.Println(err)
		return err
	}
	
	WikiPackageName = viper.GetString("wikiPackageName")
	WikiOutDir = viper.GetString("wikiOutDir")
	
	pMap := viper.GetStringMap("pages")
	
	for k,_ :=range pMap {
		
		page := WikiPage{}
		
		key := "pages." + k + "."

		page.PageName = k
		page.ApiFile = viper.GetString(key + "apifile")
		page.FuncNames = viper.GetStringSlice(key + "funcnames")

		Pages = append(Pages,page)
	}
	
	return nil
}