package templater

import (
	"os"
	"text/template"
)

type TextEntry struct {
	Number int
	Square int
}

func TextTemplate(tFile string) {
	DATA := [][]int{{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16}}

	var Entries []TextEntry
	for _, i := range DATA {
		if len(i) == 2 {
			temp := TextEntry{i[0], i[1]}
			Entries = append(Entries, temp)
		}
	}

	t := template.Must(template.ParseGlob(tFile))
	t.Execute(os.Stdout, Entries)
}
