package i18n

import (
	"fmt"
	"github.com/BurntSushi/toml"
	stdi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc/metadata"
	"context"
	"os"
)
var (
	defaultLanguage = "en"
	defaultAcceptLanguage = "en, en-US, vi"
)

type Locale interface {
	T (ID string, message *Message, params map[string]string) (string)
	SetLanguage(language string) (error)
}

type locale struct {
	Language string
	AcceptLanguage string
	Path string
	Bundle *stdi18n.Bundle
	Localizer *stdi18n.Localizer
}

type Message struct {
	Description string
	One string
	Other string
}

func New(languagePath string, acceptLanguage string) Locale {
	bundle := stdi18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	localizer := stdi18n.NewLocalizer(bundle, defaultLanguage)
	if acceptLanguage == "" {
		acceptLanguage = defaultAcceptLanguage
	}
	return &locale{
		Path: languagePath,
		Bundle: bundle,
		AcceptLanguage: acceptLanguage,
		Language: defaultLanguage,
		Localizer: localizer,
	}
}

func (i *locale) SetLanguage(language string) (error){

	languageFile := fmt.Sprintf("%s/active.%s.toml", i.Path, language)
	//fmt.Println(languageFile)
	if FileExists(languageFile) {
		//fmt.Println("Support Language")
		i.Language = language
		i.Bundle.MustLoadMessageFile(languageFile)
		i.Localizer = stdi18n.NewLocalizer(i.Bundle, i.Language, i.AcceptLanguage)
		return nil
	}else{
		//fmt.Println("Not Support Language")
		return fmt.Errorf("Language %s does not exist",language)
	}
}



func (i *locale) T(ID string, message *Message, params map[string]string) (string){
	//fmt.Println("current language:" + i.Language)
	apiMessage := &stdi18n.Message{
		ID:    ID,
		Description: message.Description,
		One: message.One,
		Other: message.Other,
	}

	msg := i.Localizer.MustLocalize(&stdi18n.LocalizeConfig{
		DefaultMessage: apiMessage,
		TemplateData:   params,
	})

	return msg
}


func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func GRPCClientlanguage(lang string) grpctransport.ClientRequestFunc {
	return func(ctx context.Context, md *metadata.MD) context.Context {
		(*md)["content-language"] = []string{lang}
		ctx = context.WithValue(ctx, "content-language", lang)
		return ctx
	}
}


// GRPCToContext moves a JWT from grpc metadata to context. Particularly
// userful for servers.
func LanguageToContext(i18N locale) grpctransport.ServerRequestFunc {
	return func(ctx context.Context, md metadata.MD) context.Context {
		// capital "Key" is illegal in HTTP/2.
		languageHeader, ok := md["content-language"]
		if !ok {
			return ctx
		}

		lang := languageHeader[0]

		i18N.SetLanguage(lang)

		ctx = context.WithValue(ctx, "content-language", lang)
		return ctx
	}
}