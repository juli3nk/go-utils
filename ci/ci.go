package ci

import (
	"fmt"
)

func ResolveVersion(version, tag, commit string, uncommitted bool) string {
	if version == "" {
		state := "clean"
		if uncommitted {
			state = "dirty"
		}

		switch {
		case state == "clean" && tag != "":
			version = tag
		default:
			version = fmt.Sprintf("%s-%s", commit, state)
		}
	}

	return version
}
