package domain

type TranslationService interface {
	Translate(content, sourceLanguage, targetLanguage string) (string, error)
}
