package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

//go:embed templates/*
var f embed.FS

type Copyright struct {
	Copyright string
}

// copied from:
// https://gist.github.com/kennwhite/306317d81ab4a885a965e25aa835b8ef
// TODO: test
func WordWrap(text string, lineWidth int) string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return text
	}
	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
			wrapped += "\n" + word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
	return wrapped
}

func ExplicitContributors(contributors []string) string {
	if len(contributors) == 0 {
		return ""
	}

	return strings.Join(contributors, ", ")
}

func ImplicitContributors(org, project string) string {
	s := "%s contributors (https://github.com/%s/%s/graphs/contributors)"
	return fmt.Sprintf(s, project, org, project)
}

func MakeContributorsList(contributors []string, org, project string) string {
	a := ExplicitContributors(contributors)
	b := ImplicitContributors(org, project)

	if a == "" {
		return b
	}

	return fmt.Sprintf("%s and %s", a, b)
}

func MakeCopyright(contributors []string, org, project, year string) string {
	cl := MakeContributorsList(contributors, org, project)
	together := fmt.Sprintf("Copyright (c) %s %s", year, cl)
	return WordWrap(together, 80)
}

func ParseProject(url string) (string, string, error) {
	re := regexp.MustCompile(`([^:]+)\.git`)
	matches := re.FindStringSubmatch(url)

	if len(matches) != 2 {
		return "", "", fmt.Errorf("Could not parse package name from remote: %s", url)
	}

	s := matches[1]

	ss := strings.Split(s, "/")
	if len(ss) != 2 {
		return "", "", fmt.Errorf("Must pass project in org/project format. Got: %s", s)
	}

	return ss[0], ss[1], nil
}

func GetProjectInfo(remoteName string) (string, string, error) {
	r, err := git.PlainOpen("")
	remote, err := r.Remote(remoteName)
	handle(err)

	url := remote.Config().URLs[0]
	return ParseProject(url)
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	currentYear := time.Now().Format("2006")

	y := flag.String("y", currentYear, "Year(s) that the license should cover")
	r := flag.String("r", "origin", "Organization/Project to cover")

	flag.Parse()

	org, proj, err := GetProjectInfo(*r)
	handle(err)

	contributors := flag.Args()

	c := Copyright{MakeCopyright(contributors, org, proj, *y)}

	// TODO: make map of licenses
	dat, err := f.ReadFile("templates/MIT.tpl")
	handle(err)

	tpl := string(dat)
	tmpl, err := template.New("test").Parse(tpl)
	if err != nil {
		panic(err)
	}

	tmpl.Execute(os.Stdout, c)
}
