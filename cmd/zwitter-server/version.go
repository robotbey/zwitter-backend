package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Version struct {
	Version    string `json:"version"`
	BuildNo    string `json:"buildNo"`
	BuildDate  string `json:"buildDate"`
	CommitHash string `json:"commitHash"`
}

func (v Version) ToShortString() string {
	return v.Version
}

func (v Version) ToLongString() string {
	return fmt.Sprintf("%s (build %s, commit %s, built at %s)", v.Version, v.BuildNo, v.CommitHash, v.BuildDate)
}

//go:embed version.txt
var versionLines string

var VERSION Version

func init() {
	parts := strings.Split(versionLines, "\n")
	VERSION = Version{
		Version:    strings.TrimSpace(parts[0]),
		BuildNo:    strings.TrimSpace(parts[1]),
		BuildDate:  strings.TrimSpace(parts[2]),
		CommitHash: strings.TrimSpace(parts[3]),
	}
}
