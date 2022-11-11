package localize

import "errors"

func NewLocalizer(service, word string) Getter {
	wordMap, ok := WordSource[service][word]
	if !ok {
		return Getter{Error: errors.New("Error in service name or word. (Cannot find word)")}
	}
	return Getter{
		Word:  wordMap,
		Error: nil,
	}
}

type Getter struct {
	Word  map[string]string
	Error error
}

func (l *Getter) Uzbek() string {
	return l.Word["uz"]
}

func (l *Getter) Russian() string {
	return l.Word["ru"]
}

func (l *Getter) English() string {
	return l.Word["en"]
}

func (l *Getter) Other(language string) string {
	return l.Word[language]
}
