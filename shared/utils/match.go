package util

import (
	"regexp"
)

//Various match functions. Does what they say on the tin
func MatchAnyURL(s string) bool {
	//From https://gist.github.com/gruber/249502
	urlRXP := regexp.MustCompile(`(?i)\b((?:[a-z][\w-]+:(?:/{1,3}|[a-z0-9%])|www\d{0,3}[.]|[a-z0-9.\-]+[.][a-z]{2,4}/)(?:[^\s()<>]+|\(([^\s()<>]+|(\([^\s()<>]+\)))*\))+(?:\(([^\s()<>]+|(\([^\s()<>]+\)))*\)|[^\s!()\[\]{};:'\".,<>?«»“”‘’]))`)

	if res := urlRXP.FindString(s); res != "" {
		return true
	}
	return false
}

func MatchGitCommit(s string) bool {
	gitRXP := regexp.MustCompile(`\S*?/\S*?\s\S*\s.*?:`)

	if res := gitRXP.FindString(s); res != "" {
		return true
	}
	return false
}
