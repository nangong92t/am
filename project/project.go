package project

import (
	"os"
	"path/filepath"
)

type Work struct {
	work      []*Info
	gitDir    string
	directory string
}

type Info struct {
	name   string
	head   string
	branch []string
	commit []map[string]string
}

func NewWork(d string) *Work {
	d, err := filepath.Abs(d)
	_, err = os.Stat(d)

	if err != nil {
		panic("directory not fount")
	}

	wk := &Work{
		directory: d,
		gitDir:    ".git",
	}
	wk.setinfo()
	return wk
}

func (wk *Work) GetInfo(name string) *Info {
	for _, v := range wk.work {
		if v.name == name {
			return v
		}
	}
	return nil
}

//获取当前分支
func (wk *Work) GetHead(name string) string {
	info := wk.GetInfo(name)
	if info != nil {
		return info.head
	}

	return ""
}

//获取项目分支列表
func (wk *Work) GetBranchList(name string) []string {
	info := wk.GetInfo(name)
	if info != nil {
		return info.branch
	}

	return nil
}

//获取commit信息
func (wk *Work) GetCommit(name string, page, pagesize int) []map[string]string {
	info := wk.GetInfo(name)
	if info == nil {
		return nil
	}

	if page <= 0 {
		page = 1
	}
	if pagesize <= 0 {
		pagesize = 10
	}
	l := len(info.commit)
	maxpage := l/pagesize + 1
	if page > maxpage {
		page = maxpage
	}

	if page == maxpage {
		pagesize = l % pagesize
	}

	return info.commit[page-1 : pagesize]
}
