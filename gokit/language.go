package gokit


import (
	"context"
)


func GetLanguageFromContext(ctx context.Context) string {
	lang := ctx.Value(LanguageKey)

	if lang != nil {
		return lang.(string)
	}

	return DefaultLanguage
}

