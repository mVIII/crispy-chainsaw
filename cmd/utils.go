package cmd

import (
	"github.com/mVIII/crispy-chainsaw/src/utils"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const (
	file=`(\w+)\.(go)$`
	dir=`[a-zA-z\/]+\/`
)

type args struct{
	GitHoster string
	Owner string
	ProjectName string
}

func getExecPath()string{
	ex, err := os.Executable()
	utils.Check(err)

	exPath := filepath.Dir(ex)
	
	return exPath
}
//  Returns emptry string on path if none found
func destructurePath(p string)(path string, file string){
	return filepath.Split(p)
}

func makeDirs(p string,perm os.FileMode){
	err:=os.MkdirAll(p,perm)
	utils.Check(err)
}

func makeFile(p string)*os.File{

	f,err:=os.Create("./"+p)
	utils.Check(err)

	return f
}

func multipleContains(s string,strs ...string)bool{

	for _,str:=range strs{
		if strings.Contains(s,str) {
			return true
		}
	}
	return false
}

func fillTemplate(file io.Writer,fillable string ,data args){
	tmpl, err := template.New("bruh").Parse(fillable)
	if err!=nil{
		log.Fatal(err)
	}
	err = tmpl.Execute(file,data)
}