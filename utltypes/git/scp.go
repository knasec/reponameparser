package git

import (
	"errors"
	"fmt"
)

type Scp struct {
	User       string   `json:"user"`
	Host       string   `json:"host"`
	Port       string   `json:"port"`
	Group      []string `json:"Group"`
	Repository string   `json:"repository"`
}

func (s *Scp) Check(URL string) bool {
	return true
}

func (s *Scp) Parse(URL string) error {

	// Lets verify that URL is realy in SCP format
	if !s.Check(URL) {
		error_text := fmt.Sprintf("not a scp name url. expected format: [user@]example.com/group/subgroup/repository.git, got: %s", URL)
		return errors.New(error_text)
	}

	// We dont modify our URL so lets calculate its lenght one time
	len := len(URL)
	lastTrigger := 0

	for i := 0; i < len; i++ {
		c := URL[i]
		switch {

		case c == '@':
			// If we have a User lets parse him
			s.User = URL[:i]
			lastTrigger = i
		case c == ':':
			// We have a port or group
			if i+1 < len {
				// Check if there is user in URL
				if lastTrigger != 0 {
					s.Host = URL[lastTrigger+1 : i]

				} else {
					// There are not user in URL
					s.Host = URL[lastTrigger:i]
				}
				lastTrigger = i
			}
		case c == '/':
			// We have a port
			if '0' <= URL[lastTrigger+1] && URL[lastTrigger+1] <= '9' {
				if i+1 < len {
					s.Port = URL[lastTrigger+1 : i]
				}
			} else {
				// Its a group
				if i+1 < len {
					s.Group = append(s.Group, URL[lastTrigger+1:i])
				}
			}
			lastTrigger = i
		}
	}
	// Lets check the repository name
	if lastTrigger < len {
		s.Repository = URL[lastTrigger+1 : len-4]
	}
	return nil
}
