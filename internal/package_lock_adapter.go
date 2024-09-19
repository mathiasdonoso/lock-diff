package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func buildPackageLock(file []byte) (PackageLockFile, error) {
	var lockFile PackageLockFile
	if err := json.Unmarshal(file, &lockFile); err != nil {
		fmt.Printf("Error unmarshalling the json: %v\n", err)
		return PackageLockFile{}, errors.Join(JSON_FILE_ERROR, err)
	}

	type Alias struct {
		Packages     map[string]Package `json:"packages"`
		Dependencies map[string]Package `json:"dependencies"`
	}

	var aux Alias

	if err := json.Unmarshal(file, &aux); err != nil {
		fmt.Printf("Error unmarshalling the aux json: %v\n", err)
		return PackageLockFile{}, errors.Join(JSON_FILE_ERROR, err)
	}

	if aux.Packages != nil {
		lockFile.Packages = aux.Packages
	} else if aux.Dependencies != nil {
		lockFile.Packages = aux.Dependencies
	}

	return lockFile, nil
}

type PackageLockAdapter struct{}

func NewPackageLockAdapter() *PackageLockAdapter {
	return &PackageLockAdapter{}
}

func (p *PackageLockAdapter) GetDiff(lockfile1, lockfile2 io.Reader) ([]Diff, error) {
	file1, err := io.ReadAll(lockfile1)
	if err != nil {
		return []Diff{}, err
	}

	lockFile1, err := buildPackageLock(file1)
	if err != nil {
		return []Diff{}, err
	}

	file2, err := io.ReadAll(lockfile2)
	if err != nil {
		return []Diff{}, err
	}

	lockFile2, err := buildPackageLock(file2)
	if err != nil {
		return []Diff{}, err
	}

	var diff []Diff
	for p1, d1 := range lockFile1.Packages {
		for p2, d2 := range lockFile2.Packages {
			if p1 == "" || p2 == "" {
				continue
			}

			p1Formatted := strings.Replace(p1, "node_modules/", "", 1)
			p2Formatted := strings.Replace(p2, "node_modules/", "", 1)

			if strings.Compare(p1Formatted, p2Formatted) == 0 {
				if strings.Compare(d1.Version, d2.Version) != 0 {
					diff = append(diff, Diff{
						Name:             p1Formatted,
						VersionLockFile1: d1.Version,
						VersionLockFile2: d2.Version,
					})
				}
			}
		}
	}

	return diff, nil
}
