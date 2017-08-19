package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"strings"
	"testing"
)

func TestTransposeCSV(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  string
	}{
		{
			"",
			`header1,header2,header3
1,"asdf,a",3
2,asd,4
3,as,5`,
			`header1,1,2,3,
header2,"asdf,a",asd,as,
header3,3,4,5,
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var tempFiles []string
			defer func() {
				for _, name := range tempFiles {
					if fileExists(name) {
						t.Errorf("file %s shouldn't exists anymore", name)
					}
				}
			}()

			csvFile := strings.NewReader(tc.input)
			var b bytes.Buffer

			r := csv.NewReader(csvFile)
			buf := &FileBuffer{
				size: 32,
				sep:  byte(','),
			}
			defer buf.Remove()

			err := buf.Store(r)
			if err != nil {
				t.Error("buf.WriteTo returned unexpected error", err)
			}

			fmt.Println(buf.names)
			tempFiles = append(tempFiles, buf.names...)

			_, err = buf.WriteTo(&b)
			if err != nil {
				t.Error("buf.WriteTo returned unexpected error", err)
			}

			if got := b.String(); got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
				t.Errorf("got %d; want %d", len(got), len(tc.want))
			}
		})
	}
}
