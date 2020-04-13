package git

import (
	"encoding/json"
	"github.com/mVIII/crispy-chainsaw/src/utils"
	"io/ioutil"
	"net/http"
	"sync"
)



type FileResponse struct {
	Path        string `json:"path"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	DownloadUrl string `json:"download_url"`
}

type GitHubClient struct {
	owner      string
	name       string
	httpClient http.Client
}

func NewGitHubClient(owner, name string) *GitHubClient {

	g := &GitHubClient{
		owner:      owner,
		name:       name,
		httpClient: http.Client{},
	}

	return g
}

func (g *GitHubClient) RepoStructure() []File {
	var paths []File
	ch := make(chan File)
	quit := make(chan bool)
	go g.getFilesWithPaths("", &paths, ch, quit, nil)

	for {
		select {
		case smt := <-ch:
			paths = append(paths, smt)
		case <-quit:
			return paths
		}

	}

}

func (g *GitHubClient) getFilesWithPaths(filep string, paths *[]File, ch chan<- File, quit chan bool, wgp *sync.WaitGroup) {
	current := g.owner +"/"+ g.name + "/contents" + filep
	var wg sync.WaitGroup

	objects, _ := g.githubFetch(current)

	for _, value := range objects {
		if value.Type != "dir" {
			res, _ := g.httpClient.Get(value.DownloadUrl+apiCredsGithub)
			body, _ := ioutil.ReadAll(res.Body)

			temp := File{Path: value.Path, Value: string(body)}
			ch <- temp
		} else {
			wg.Add(1)
			// need "/" because value path is like "src/internals/main.go"
			go g.getFilesWithPaths("/"+value.Path, paths, ch, nil, &wg)
		}
	}
	wg.Wait()
	if wgp != nil {
		wgp.Done()
	}
	if quit != nil {
		quit <- true
	}
}

//func filterDirs(r []response) []response {
//
//	var dirs []response
//
//	for _, n := range r {
//
//		if n.Type == "dir" {
//			dirs = append(dirs, n)
//		}
//	}
//	return dirs
//}

func (g *GitHubClient) githubFetch(p string) ([]FileResponse, error) {
	var respFiles []FileResponse
	r, err := g.httpClient.Get(apiUrlgithub + p+apiCredsGithub)
	utils.Check(err)
	body, err := ioutil.ReadAll(r.Body)
	utils.Check(err)

	err = json.Unmarshal(body, &respFiles)

	utils.Check(err)

	return respFiles, nil
}
