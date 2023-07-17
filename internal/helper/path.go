package helper

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	rootpath   = filepath.Join(basepath, "..", "..")
)

func GetRootPath() string {
	return rootpath
}

func GetInternalPath() string {
	return filepath.Join(GetRootPath(), "internal")
}
