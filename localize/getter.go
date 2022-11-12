package localize

import "errors"

func NewLocalizer(service string) (server ServiceGetter) {
	serviceMap, ok := WordSource[service]
	if !ok {
		return ServiceGetter{Error: errors.New("Error in service name. (Cannot find service)")}
	}
	return ServiceGetter{
		Service: serviceMap,
		Error:   nil,
	}
}

type WordGetter struct {
	word  map[string]string
	Error error
}
type ServiceGetter struct {
	Service map[string]map[string]string
	Error   error
}

func (s ServiceGetter) Localize(word string) WordGetter {
	wordMap, ok := s.Service[word]
	if !ok {
		return WordGetter{Error: errors.New("Error in word. (Cannot find word)")}
	}
	return WordGetter{
		word:  wordMap,
		Error: nil,
	}
}

func (l *WordGetter) Uzbek() string {
	return l.word["uz"]
}

func (l *WordGetter) Russian() string {
	return l.word["ru"]
}

func (l *WordGetter) English() string {
	return l.word["en"]
}

func (l *WordGetter) Other(language string) string {
	return l.word[language]
}
