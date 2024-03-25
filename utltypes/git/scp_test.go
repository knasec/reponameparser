package git

import (
	"testing"
)

func TestScpParseCorrectPaths(t *testing.T) {
	var tests = []struct {
		input string
		want  interface{}
	}{
		{"git@example.com:group/repo.git", &Scp{User: "git", Host: "example.com", Port: "", Group: []string{"group"}, Repository: "repo"}},
		{"git@example.com:32222/group/repo.git", &Scp{User: "git", Host: "example.com", Port: "32222", Group: []string{"group"}, Repository: "repo"}},
		{"git@example.com:group/subgroup/repo.git", &Scp{User: "git", Host: "example.com", Port: "", Group: []string{"group"}, Repository: "repo"}},
		{"example.com:group/subgroup/repo.git", &Scp{User: "", Host: "example.com", Port: "", Group: []string{"group"}, Repository: "repo"}},
		{"git@example.com:32222/group/subgroup/repo.git", &Scp{User: "git", Host: "example.com", Port: "32222", Group: []string{"group"}, Repository: "repo"}},
		/*
			{"", errors.New("empty url")},
			{"git@example.com/group/repo", errors.New("not a git repository url")},
			{"git@example.com:repo.git", errors.New("no userpath or group present")},
			{"example.com/group/repo.git", errors.New("no repository schema detected")},
		*/
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var s Scp
			err := s.Parse(tt.input)
			if err != tt.want && err != nil {
				t.Errorf("\nGot\n %s, want\n %s\n", err, tt.want)
			}
			if DeepEqual(s, tt.want) {
				t.Errorf("\nGot %+v\n, want %+v\n", s, tt.want)
			}
		})
	}
}
