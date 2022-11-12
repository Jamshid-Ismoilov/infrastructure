package localize

import "errors"

// returns localizer for service according to name
func NewLocalizer(service string) (server ServiceGetter) {
	serviceMap, ok := WordSource[service]
	if !ok {
		return ServiceGetter{Error: errors.New("Error in service name. (Cannot find service)")}
	}
	return ServiceGetter{
		service: serviceMap,
		Error:   nil,
	}
}

// getter for service
type ServiceGetter struct {
	service map[string]map[string]string
	Error   error
}

// returns translation of word accourding to word name and language
func (s ServiceGetter) Localize(word string, lang string) (localizedWord string) {
	localizedWord, ok := s.service[word][lang]
	if !ok {
		return ""
	}
	return localizedWord
}

// returns word type, which includes localization settings
func (s *ServiceGetter) GetWord(word string) WordGetter {
	wordMap, ok := s.service[word]
	if !ok {
		return WordGetter{Error: errors.New("Error in word name. (Cannot find word)")}
	}
	return WordGetter{
		word:  wordMap,
		Error: nil,
	}
}

// getter for word
type WordGetter struct {
	word  map[string]string
	Error error
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
