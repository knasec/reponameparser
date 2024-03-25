package git

import (
	"errors"
	"strconv"
	"strings"
)

type Git struct {
	User       string   `json:"user"`
	Address    string   `json:"address"`
	Port       int      `json:"port"`
	Group      string   `json:"group"`
	Subgroups  []string `json:"subgroups"`
	Repository string   `json:"repository"`
	Path       string   `json:"path"`
}

func (g Git) Check(URL string) bool {

	// Patterns
	// @.://.git
	// @.:[1-9]/.
	// Pattern breakdown
	// @ - 1
	// . - many
	// : - 1
	// / - many
	// .git - 1 at the end

	for i := 0; i < len(URL); i++ {
		c := URL[i]

		switch {
		case c == '@':
			if URL[4:] == ".git" {
				return true
			}
		case c == '.':
			//
		case c == ':':
			// Fire 1
		case '0' <= c && c <= '9':
			//
		}
	}
	return false
}

// git@example.com:32222/group/subgroup/reponame.git
// git@example.com:group/subgroup/reponame.git
func (g Git) Parse(URL string) error {

	// Pattern: @.:[0-9]/.

	// Check if its real git repo sintax
	gitprefix := URL[:4]
	if gitprefix != "git@" {
		return errors.New("not a git repository url")
	}

	// Lets parse URL
	tempURL := URL[4:]
	splittedURL := strings.Split(tempURL, ":")
	// Now lets check that we had that ":" in our URL
	if len(splittedURL) < 2 {
		return errors.New("not a git repository url")
	}
	// Fill address
	g.Address = splittedURL[0]

	tempURLPath := strings.Split(splittedURL[1], "/")
	// Check that we have at least root group or username
	if len(tempURLPath) < 1 {
		return errors.New("no userpath or group present")
	}
	// Now lets check what we have port number or root group
	port, err := strconv.Atoi(tempURLPath[0])
	if err != nil {
		// There is no port, its root group
		g.Group = tempURLPath[0]
	} else {
		g.Port = port
	}

	return nil
}

// GL65 9ESK
