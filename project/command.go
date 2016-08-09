package project

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

//"os/exec"
//"path/filepath"

func (wk *Work) setinfo() {
	/*cmd := &exec.Cmd{
		Path: "git",
		Dir:  wk.directory,
	}

	if filepath.Base(cmd.Path) == cmd.Path {
		if lp, err := exec.LookPath(cmd.Path); err != nil {
			panic("Did not find the GIT command")
		} else {
			cmd.Path = lp
		}
	}*/
	//dis := []string{"head", "branch", "remote", "remoteBranch"}
	//wk.
}

func (wk *Work) getHead() string {
	path := filepath.Join(wk.directory, wk.gitDir, "HEAD")
	buff, err := wk.readFile(path)
	if err != nil {
		return ""
	}

	reg := regexp.MustCompile(`[a-zA-Z0-9_\-]+$`)
	match := reg.FindAll(buff, -1)
	if len(match) == 0 {
		return ""
	}
	return string(match[0])
}

func (wk *Work) getBranch() []string {
	path := filepath.Join(wk.directory, wk.gitDir, "refs", "heads")
	var files []string
	err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		files = append(files, filepath.Base(path))
		return nil
	})
	if err != nil {
		return nil
	}

	return files
}

func (wk *Work) readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buff, err := ioutil.ReadAll(f)

	if err != nil {
		return nil, err
	}

	return buff, nil
}
