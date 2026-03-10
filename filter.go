package prompt

import (
	"strings"
	"unicode/utf8"
)

// Filter is the type to filter the prompt.Suggestion array.
type Filter func([]Suggest, string, bool) []Suggest

// FilterHasPrefix checks whether the string completions.Text begins with sub.
func FilterHasPrefix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.HasPrefix)
}

// FilterHasSuffix checks whether the completion.Text ends with sub.
func FilterHasSuffix(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.HasSuffix)
}

// FilterContains checks whether the completion.Text contains sub.
func FilterContains(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, strings.Contains)
}

// FilterFuzzy checks whether the completion.Text fuzzy matches sub.
// Fuzzy searching for "dog" is equivalent to "*d*o*g*". This search term
// would match, for example, "Good food is gone"
//                               ^  ^      ^
func FilterFuzzy(completions []Suggest, sub string, ignoreCase bool) []Suggest {
	return filterSuggestions(completions, sub, ignoreCase, fuzzyMatch)
}

func fuzzyMatch(s, sub string) bool {
	sIdx := 0

	for _, c := range sub {
		found := false
		for sIdx < len(s) {
			r, size := utf8.DecodeRuneInString(s[sIdx:])
			sIdx += size
			if r == c {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func filterSuggestions(suggestions []Suggest, sub string, ignoreCase bool, function func(string, string) bool) []Suggest {
	if sub == "" {
		return suggestions
	}
	if ignoreCase {
		sub = strings.ToUpper(sub)
	}

	ret := make([]Suggest, 0, len(suggestions))
	for i := range suggestions {
		c := suggestions[i].Text
		if ignoreCase {
			c = strings.ToUpper(c)
		}
		if function(c, sub) {
			ret = append(ret, suggestions[i])
		}
	}
	return ret
}
