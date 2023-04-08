package fake

type FakeResourceExpansionData struct {
	UpResource string // Credential
}

var FakeResourceExpansionTpl = `
package fake

type Fake{{.UpResource}}Expansion interface {}
`
