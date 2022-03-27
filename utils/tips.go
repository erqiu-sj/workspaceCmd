package utils

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"github.com/manifoldco/promptui"
	"strings"
)

func GreenTips(text string) {
	fmt.Println(
		aurora.Bold(aurora.Green(text)),
	)
	//fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+"  \x1b[0m\n", 32)
}

func FillGrayTips(text interface{}) {
	fmt.Println(
		aurora.Bold(aurora.Gray(100, text)),
	)
}

func RedTips(text interface{}) {
	fmt.Println(
		aurora.Bold(aurora.Red(text)),
	)
	//fmt.Printf("\x1b[%dm"+fmt.Sprint(text)+" \x1b[0m\n", 31)
}

type ConciseSelector struct {
	Label string
	Desc  string
}

func WarningModelTips(label string, list []ConciseSelector) promptui.Select {
	tem := promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🍅 {{ .Label | cyan }}",
		Inactive: "  {{ .Label | cyan }}",
		Selected: "🍅 {{ .Label | red | cyan }}",
		Details: `
--------- warning ----------
{{ "label:" | faint }}	{{ .Label }}
{{ "desc" | faint }}	{{ .Desc }}`,
	}
	return promptui.Select{
		Label:             label,
		Items:             list,
		HideHelp:          false,
		Templates:         &tem,
		StartInSearchMode: false,
		Searcher: func(input string, index int) bool {
			c := list[index]
			if strings.Contains(c.Label, input) {
				return true
			}
			return false
		},
	}
}

// InquiryTips 询问
func InquiryTips(label string) promptui.Prompt {
	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	return promptui.Prompt{
		Label:     label,
		Templates: templates,
	}
}
