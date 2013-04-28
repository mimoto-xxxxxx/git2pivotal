package pivotal

import (
	"regexp"
	"strings"
)

type Story struct {
	Id string
}

var finder = regexp.MustCompile(`https://www\.pivotaltracker\.com/story/show/(\d+)|pivotal:(\d+)`)

func FindStoryFromString(s string) []*Story {
	matches := finder.FindAllStringSubmatch(s, -1)
	ret := make([]*Story, len(matches))
	for i, v := range matches {
		ret[i] = &Story{
			Id: v[1] + v[2],
		}
	}
	return ret
}

func RemoveStoryFromString(s string) string {
	return strings.TrimSpace(finder.ReplaceAllString(s, ""))
}
