package git

import (
	"testing"
)

func TestGetRepo(t *testing.T) {
	gc :=NewGitHubClient("mVIII","http-api-template")
	p ,err:=gc.githubFetch(gc.owner+"/"+gc.name)
	if err!=nil {
		t.Log(err)
		t.Fail()
	}
	//t.Log(p.ContentsUrl)
	for _,r := range p  {
		//fmt.Println(r.Path)
		t.Log(r.Url)
	}
}
func TestGetFilePaths(t *testing.T) {
	gc :=NewGitHubClient("mVIII","http-api-template")
	p :=gc.RepoStructure()

	//t.Log(p.ContentsUrl)
	for _,r := range p  {
		//fmt.Println(r.Path)
		t.Log(r.Value)
	}
}

