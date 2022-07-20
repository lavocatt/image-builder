package common

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type AllowList map[string][]string

// Loads the allow list from an allowFile and returns an AllowList. If an empty
// string is given as the argument, returns an empty AllowList.
//
// The allow file must conform to the following json schema:
// {
// 	"000000": ["fedora-*"]
//  "000001": ["fedora-34", "fedora-35", "fedora-36"]
//  "000002": []
// }
func LoadAllowList(allowFile string) (AllowList, error) {
	if allowFile == "" {
		return AllowList{}, nil
	}

	jsonFile, err := os.Open(filepath.Clean(allowFile))
	if err != nil {
		return nil, fmt.Errorf("No allow file found at %s: %v", allowFile, err)
	}
	defer func() {
		if err := jsonFile.Close(); err != nil {
			logrus.Errorln(fmt.Sprintf("Error closing file: %s", err))
		}
	}()

	rawJsonFile, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("Failed to read allow file %q: %s", allowFile, err.Error())
	}

	var allowList AllowList
	err = json.Unmarshal(rawJsonFile, &allowList)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal allow file %q: %s", allowFile, err.Error())
	}

	return allowList, nil
}

func (a AllowList) IsAllowed(orgId, distro string) (bool, error) {
	for _, allowedDistro := range a[orgId] {
		// path.Match() supports matching glob patterns for distros, e.g. fedora-*
		match, err := path.Match(allowedDistro, distro)
		if err != nil {
			return false, err
		}
		if match {
			return true, nil
		}
	}
	return false, nil
}
