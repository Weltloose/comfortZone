package model

import (
	"io/ioutil"
)

func AddSource() int {
	tmp := GetCurSources()
	addCurSources()
	return tmp
}

func TransferSource(sourceFile, destFile, realFile string) {
	bytes, _ := ioutil.ReadFile(sourceFile)

	ioutil.WriteFile(destFile, bytes, 0666)
	newSource(destFile, realFile)
}
