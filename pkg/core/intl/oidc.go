package intl

import (
	"fmt"
	"strings"
)

// LocalizeJSONObject returns the localized value of key in jsonObject according to preferredLanguageTags.
func LocalizeJSONObject(preferredLanguageTags []string, fallbackLanguageTag string, jsonObject map[string]interface{}, key string) string {
	prefix := fmt.Sprintf("%s#", key)
	m := map[string]string{}
	for k, v := range jsonObject {
		stringValue, ok := v.(string)
		if !ok {
			continue
		}
		if k == key {
			m[fallbackLanguageTag] = stringValue
		} else if strings.HasPrefix(k, prefix) {
			tag := strings.TrimPrefix(k, prefix)
			m[tag] = stringValue
		}
	}

	var supportedLanguageTags []string
	for tag := range m {
		supportedLanguageTags = append(supportedLanguageTags, tag)
	}
	supportedLanguageTags = SortSupported(supportedLanguageTags, fallbackLanguageTag)

	idx, _ := Match(preferredLanguageTags, supportedLanguageTags)
	tag := supportedLanguageTags[idx]
	value := m[tag]
	return value
}

// LocalizeStringMap returns the localized value of key in stringMap according to preferredLanguageTags.
func LocalizeStringMap(preferredLanguageTags []string, fallbackLanguageTag string, stringMap map[string]string, key string) string {
	prefix := fmt.Sprintf("%s#", key)
	m := map[string]string{}
	for k, stringValue := range stringMap {
		if k == key {
			m[fallbackLanguageTag] = stringValue
		} else if strings.HasPrefix(k, prefix) {
			tag := strings.TrimPrefix(k, prefix)
			m[tag] = stringValue
		}
	}

	var supportedLanguageTags []string
	for tag := range m {
		supportedLanguageTags = append(supportedLanguageTags, tag)
	}
	supportedLanguageTags = SortSupported(supportedLanguageTags, fallbackLanguageTag)

	idx, _ := Match(preferredLanguageTags, supportedLanguageTags)
	tag := supportedLanguageTags[idx]
	value := m[tag]
	return value
}
