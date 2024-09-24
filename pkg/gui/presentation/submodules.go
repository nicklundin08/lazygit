package presentation

import (
	"fmt"

	"github.com/jesseduffield/lazygit/pkg/commands/models"
	"github.com/jesseduffield/lazygit/pkg/theme"
	"github.com/samber/lo"
)

func GetSubmoduleListDisplayStrings(submodules []*models.SubmoduleConfig) [][]string {
	return lo.Map(submodules, func(submodule *models.SubmoduleConfig, _ int) []string {
		return getSubmoduleDisplayStrings(submodule)
	})
}

func getSubmoduleDisplayStrings(s *models.SubmoduleConfig) []string {
	// Arbitrarily add 30 spaces to the format. Not married to this but I just wanted it to look aligned
	// Put the HEAD first because those are more likely to be similar lengths?
	name := fmt.Sprintf("%-30s\t%s", s.Head,
		s.Name,
	)
	if s.ParentModule != nil {
		indentation := ""
		for p := s.ParentModule; p != nil; p = p.ParentModule {
			indentation += "  "
		}

		name = indentation + "- " + s.Name
	}

	if s.NumStagedChanges != 0 {
		name = fmt.Sprintf(
			"%s +%d",
			name,
			s.NumStagedChanges,
		)
	}

	if s.NumUnstagedChanges != 0 {
		name = fmt.Sprintf(
			"%s !%d",
			name,
			s.NumUnstagedChanges,
		)
	}

	if s.NumUntrackedChanges != 0 {
		name = fmt.Sprintf(
			"%s ?%d ",
			name,
			s.NumUntrackedChanges,
		)
	}

	return []string{theme.DefaultTextColor.Sprint(name)}
}
