package reponameparser

// Detect repo name and components from URL
// Supported types of names
// user@example.com:group/reponame.git
// user@example.com:32222/group/reponame.git
// https://example.com/group/reponame.git
// hub.example.com/team/imagename:version
// registry.example.com/team/imagename:version
// registry.example.com:32222/team/imagename:version

// Patterns
// ssh: @.:/.
// ssh: @.:[1-9]/.
// https: ://.//.
// https: ://.:[1-9]/.
// image: ..//:
// image: ..:[1-9]//:

// git@gitlab.ssdlc.sec.nctmsk.ru:ssdlc/devsecops/project/opensource/3pl-analyze/3pl-scan.git
// @....://///.
// shema://user@host:/

func Parse(URL string) {

}
