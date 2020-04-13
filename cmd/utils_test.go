package cmd

import (
	"path/filepath"
	"testing"
)

func TestDestructurePath(t *testing.T){

	testPath :="src/http/wire.go"
	dir,file :=filepath.Split(testPath)
	t.Log(file)
	t.Log(dir)
}

func TestGetExecPath(t*testing.T){

	getExecPath()
}