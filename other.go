package rules_go_pkgsite

import "fmt"

func (m *Foo) SomeOtherMethod() string {
	if m.ADouble != nil {
		return fmt.Sprintf("%s-%f", m.AString, *m.ADouble)
	}
	return m.AString
}
