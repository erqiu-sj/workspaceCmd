package core

import (
	"github.com/manifoldco/promptui"
	"strings"
	"workspaceCmd/utils"
)

// SelectWorkSpaceGroup returns FinalAnalysisConf
func SelectWorkSpaceGroup(list []FinalAnalysisConf) ([]FinalAnalysisConf, int) {
	createSelect := CreateSelect(list, "select work group")
	i, _, err := createSelect.Run()
	utils.ColdKiller(err)
	return list, i
}

func SelectWorkSpace(list []FinalAnalysisConf, groupName string) ([]WorkSpace, int) {
	var result []WorkSpace
	for _, item := range list {
		if item.GroupName != groupName {
			continue
		}
		result = item.SpaceList
		break
	}
	space := CreateSelectWithWorkSpace(result, "please select a workspace")
	i, _, err := space.Run()
	utils.ColdKiller(err)
	return result, i
}

func CreateSelect(list []FinalAnalysisConf, title string) promptui.Select {
	tem := promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🍅 {{ .GroupName | cyan }} ({{ .GroupRemarks | red }})\"",
		Inactive: "  {{ .GroupName | cyan }} ({{ .GroupRemarks | red }})",
		Selected: "🍅 {{ .GroupName | red | cyan }}",
		Details: `
--------- Profile ----------
{{ "Name:" | faint }}	{{ .GroupName }}
{{ "Remarks" | faint }}	{{ .GroupRemarks }}
{{ "Work group config file  path" | faint }}	{{ .WorkGroupConfigPath }}
{{ "Last open mode" | faint }}	{{ .LastOpenMethod }}
`,
	}
	return promptui.Select{
		Label:             title,
		Items:             list,
		HideHelp:          false,
		Templates:         &tem,
		StartInSearchMode: false,
		Searcher: func(input string, index int) bool {
			c := list[index]
			if strings.Contains(c.GroupName, input) {
				return true
			}
			return false
		},
	}
}

func CreateSelectWithWorkSpace(list []WorkSpace, title string) promptui.Select {
	tem := promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🍅 {{ .Name | cyan }} ({{ .Remarks | red }})\"",
		Inactive: "  {{ .Name | cyan }} ({{ .Remarks | red }})",
		Selected: "🍅 {{ .Name | red | cyan }}",
		Details: `
--------- Profile ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Remarks" | faint }}	{{ .Remarks }}
{{ "Path:" | faint }}	{{ .Path }}
{{ "Working group:" | faint }}	{{ .WithGroup }}
{{ "Last open mode" | faint }}	{{ .LastOpenMethod }}
`,
	}
	return promptui.Select{
		Label:             title,
		Items:             list,
		HideHelp:          false,
		Templates:         &tem,
		StartInSearchMode: false,
		Searcher: func(input string, index int) bool {
			c := list[index]
			if strings.Contains(c.Name, input) {
				return true
			}
			return false
		},
	}
}

func CreateOpenModeSelect() OpenMode {
	openMode := OpenMode{}
	list := openMode.GetOpenModeConfig()
	tem := promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🍅 {{ .Label | cyan }}",
		Inactive: "  {{ .Label | cyan }}",
		Selected: "🍅 {{ .Label | red | cyan }}",
		Details: `
--------- Profile ----------
{{ "Name:" | faint }}	{{ .Label }}
{{ "Remarks" | faint }}	{{ .CmdAlias }}
`,
	}
	modelSelect := promptui.Select{
		Label:             "please select the opening method",
		Items:             list,
		HideHelp:          false,
		Templates:         &tem,
		StartInSearchMode: false,
		Searcher: func(input string, index int) bool {
			c := list[index].Label
			if strings.Contains(c, input) {
				return true
			}
			return false
		},
	}
	i, _, err := modelSelect.Run()
	utils.ColdKiller(err)
	return list[i]
}
