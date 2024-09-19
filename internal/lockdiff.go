package internal

import (
	"errors"
	"io"
)

var (
	JSON_FILE_ERROR = errors.New("Error building from json file")
)

type Package struct {
	Version   string `json:"version"`
	Resolved  string `json:"resolved"`
	Integrity string `json:"integrity"`
}

type PackageLockFile struct {
	Name            string `json:"name"`
	LockfileVersion int    `json:"lockfileVersion"`
	Requires        bool   `json:"requires"`
	Packages        map[string]Package
}

type Diff struct {
	Name             string
	VersionLockFile1 string
	VersionLockFile2 string
}

type DependenciesAdapter interface {
	GetDiff(lockfile1, lockfile2 io.Reader) ([]Diff, error)
}

type ReporterAdapter interface {
	Print(diff []Diff) error
}

type ReporterService interface {
	GetReport(r1 io.Reader, r2 io.Reader) error
}
