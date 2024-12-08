package i18n

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocal(t *testing.T) {
	s, err := Locale.Localize(&i18n.LocalizeConfig{
		MessageID: "PersonCats",
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 2,
		},
	})
	assert.NoError(t, err)
	assert.Equal(t, "Nick 有 2 猫.", s)
}
