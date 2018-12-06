package language

import (
	"github.com/stephengrider/gorunner/language/c"
	"github.com/stephengrider/gorunner/language/cpp"
	"github.com/stephengrider/gorunner/language/java"
	"github.com/stephengrider/gorunner/language/javascript"
	"github.com/stephengrider/gorunner/language/python"
)

type runFn func([]string, string) (string, string, error)

var languages = map[string]runFn{
	"c":            c.Run,
	"cpp":          cpp.Run,
	"java":         java.Run,
	"javascript":   javascript.Run,
	"python":       python.Run,
}

func IsSupported(lang string) bool {
	_, supported := languages[lang]
	return supported
}

func Run(lang string, files []string, stdin string) (string, string, error) {
	return languages[lang](files, stdin)
}
