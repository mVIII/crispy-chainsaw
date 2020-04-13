/*
Copyright © 2019 NAME HERE konkovac@me.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/mVIII/crispy-chainsaw/src/git"
	"github.com/mVIII/crispy-chainsaw/src/utils"
	"github.com/spf13/cobra"
	"os"
	"sync"

	//"sync"

	//"sync"

	//"sync"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates http api",
	Long:  `Creates http api`,
	Run: func(cmd *cobra.Command, args []string) {

		ProjectPath, _ := cmd.Flags().GetString("projectPath")

		if ProjectPath == "" {
			fmt.Println("Project path required!")
			os.Exit(1)
		}
		owner, _ := cmd.Flags().GetString("owner")

		if owner == "" {
			fmt.Println("Project path required!")
			os.Exit(1)
		}
		gitHoster, _ := cmd.Flags().GetString("gitHoster")

		if gitHoster == "" {
			fmt.Println("Git hoster service.")
			os.Exit(1)
		}
		create(ProjectPath,owner,gitHoster)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	flag := initCmd.Flags()
	flag.StringP("projectPath", "p", "", "Project path!")
	flag.StringP("owner", "o", "", "Owner of project!")
	flag.StringP("gitHoster", "k", "", "Hoster service of project!")

	err := cobra.MarkFlagRequired(flag, "projectPath")
	utils.Check(err)
	err = cobra.MarkFlagRequired(flag, "gitHoster")
	utils.Check(err)
	err = cobra.MarkFlagRequired(flag, "owner")
	utils.Check(err)

}

type out struct {
	path     string
	template string
}

func read (client git.GitClient) []git.File{
	return 	client.RepoStructure()
}

func write(files *[]git.File,data args) {
	var wg sync.WaitGroup

	for _, file := range *files {
		wg.Add(1)
		go func(wg *sync.WaitGroup,file git.File,data args) {
			fmt.Println(file.Path)
			p, _ :=destructurePath(file.Path)
			if p!="" {
				makeDirs(p,os.ModePerm)
			}
			wf:=makeFile(file.Path)
			fillTemplate(wf,file.Value,data)

			wf.Sync()
			wf.Close()
			wg.Done()
		}(&wg,file,data)
	}
	wg.Wait()
}

func create(projectPath string, owner string,gitHoster string) {

	var gitClient  git.GitClient
	var data args
	data.Owner= owner
	data.ProjectName = projectPath
	switch gitHoster {
	case "github.com":
		data.GitHoster = "github.com"
		gitClient = git.NewGitHubClient(owner,projectPath)
	}

	files := read(gitClient)
	write(&files,data)
}
