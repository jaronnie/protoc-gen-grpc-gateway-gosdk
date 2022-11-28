package tpl

type GoModData struct {
	GoModule  string
	GoVersion string
}

var GoModTpl = `module {{.GoModule}}

go {{.GoVersion}}
`
