package git

import (
	"errors"
	"reflect"
	"testing"
)

func DeepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func TestGitParseCorrectPaths(t *testing.T) {
	var tests = []struct {
		input string
		want  interface{}
	}{
		{"git@example.com:group/repo.git", &Git{Address: "example.com", Group: "group", Subgroups: nil, Repository: "repo", Path: ""}},
		{"git@example.com:32222/group/repo.git", &Git{Address: "example.com", Port: 32222, Group: "group", Subgroups: nil, Repository: "repo", Path: ""}},
		{"git@example.com:group/subgroup/repo.git", &Git{Address: "example.com", Group: "group", Subgroups: nil, Repository: "repo", Path: ""}},
		{"", errors.New("empty url")},
		{"git@example.com/group/repo", errors.New("not a git repository url")},
		{"git@example.com:repo.git", errors.New("no userpath or group present")},
		{"example.com/group/repo.git", errors.New("no repository schema detected")},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var g Git
			err := g.Parse(tt.input)
			if err != tt.want {
				t.Errorf("Got\n %s, want\n %s\n", err, tt.want)
			}
			if DeepEqual(g, tt.want) {
				t.Errorf("Got %+v\n, want %+v\n", g, tt.want)
			}
		})
	}
}
