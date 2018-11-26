package parser

import (
	"regexp"
	"strconv"

	"github.com/learngo/crawler/model"

	"github.com/learngo/crawler/engine"
)

var ageRegexp = regexp.MustCompile(`<td><span class="label">年龄：</span>(.+)</td>`)

var marriageRegexp = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	profile := model.Profile{}
	profile.Name = name

	age, err := strconv.Atoi(extractString(contents, ageRegexp))
	if err == nil {
		profile.Age = age
	} else {
		profile.Age = 0
	}

	profile.Marriage = extractString(contents, marriageRegexp)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)

	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
