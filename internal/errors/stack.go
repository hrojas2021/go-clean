package errors

import (
	"bytes"
	"regexp"
	"runtime/debug"
)

var (
	ProjectPrefix     = "Boiler"
	ProjectPathPrefix = ProjectPrefix + "/"

	IgnoreCallerPrefixs = regexp.MustCompile(
		`(?i)(generated|/middleware\.go|_gen.go|pkg/errors/|pkg/io/database/database.go|` +
			`graphql/handle\.go|router\.go|rest/resp.go)`,
	)

	stackRegex = regexp.MustCompile(
		`(?i)((?:` + ProjectPrefix + `/|main)[^\n]+)\s+([^\n]+` + ProjectPathPrefix + `)([^\s]+)([^\n]*)`,
	)
)

func CallerByLevel(lvls ...int) string {
	var lvl int
	if len(lvls) > 0 {
		lvl = lvls[0]
	}

	lvl += 2

	_, stack := GetStack()
	if len(stack) > lvl {
		return stack[lvl]
	}

	return ""
}

func Caller() string {
	_, stack := GetStack()
	for _, s := range stack {
		return s
	}

	return ""
}

func GetStack() (string, []string) {
	tx := debug.Stack()
	matches := stackRegex.FindAllSubmatch(tx, -1)

	files := make([]string, 0, len(matches))
	for _, m := range matches {
		if !IgnoreCallerPrefixs.Match(m[3]) {
			files = append(files, string(m[3]))
		}
	}

	return string(tx[:bytes.IndexByte(tx, '\n')]), files
}
