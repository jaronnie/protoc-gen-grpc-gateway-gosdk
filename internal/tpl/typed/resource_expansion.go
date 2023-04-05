package typed

type ResourceExpansionData struct {
	ScopeVersion string // corev1
	UpResource   string // Credential
}

var ResourceExpansionTpl = `
package {{.ScopeVersion}}

type {{.UpResource}}Expansion interface {}
`
