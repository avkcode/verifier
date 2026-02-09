package helpui

import "strings"

var commandOwners = map[string][]string{
	"verifier apply":      {"internal/deploy", "internal/ui"},
	"verifier apply plan": {"internal/deploy", "internal/ui"},
	"verifier build":      {"internal/workflows/buildsvc"},
	"verifier delete":     {"internal/deploy", "internal/ui"},
	"verifier init":       {"internal/appconfig"},
	"verifier help":       {"internal/helpui"},
	"verifier logs":       {"internal/tailer"},
	"verifier secrets":    {"internal/secretstore"},
	"verifier stack":      {"internal/stack"},
}

func ownersForCommand(path string) []string {
	path = strings.TrimSpace(path)
	if path == "" {
		return nil
	}
	// Allow subcommands to inherit ownership from parent commands.
	for candidate := path; candidate != ""; {
		if owners, ok := commandOwners[candidate]; ok && len(owners) > 0 {
			out := make([]string, 0, len(owners))
			out = append(out, owners...)
			return out
		}
		idx := strings.LastIndex(candidate, " ")
		if idx < 0 {
			break
		}
		candidate = strings.TrimSpace(candidate[:idx])
	}
	return nil
}
