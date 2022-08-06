package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

var (
	inputDir    string
	pkg         string
	outFileName string
)

func parseFlags() {
	flag.StringVar(&inputDir, "in", "",
		"The directory where the .CS enum files are located")
	flag.StringVar(&pkg, "pkg", "",
		"The name of the output Go package")
	flag.StringVar(&outFileName, "out", "",
		"The name of the output .GO file")

	flag.Parse()
}

func main() {
	parseFlags()

	regexHeader, err := regexp.Compile(`^\s*\w*\s*enum (\w+)\s*\:*\s*(\w*)$`)
	handleError(err)
	regexMember, err := regexp.Compile(`^\s*(\w+)\s*=\s*(.+)[,;]+$`)
	handleError(err)

	files, err := ioutil.ReadDir(inputDir)
	handleError(err)

	var (
		enums       []Enum
		currentEnum Enum
	)

	for _, fileInfo := range files {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".cs") {
			fpath := path.Join(inputDir, fileInfo.Name())
			file, err := os.Open(fpath)
			handleError(err)
			defer file.Close()
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				line := scanner.Text()

				if strings.Contains(line, "}") && currentEnum.Name != "" {
					enums = append(enums, currentEnum)
					currentEnum = Enum{}
				}

				if parts := regexHeader.FindStringSubmatch(line); len(parts) > 0 {
					if len(parts) > 1 {
						currentEnum.Name = parts[1]
					}

					if len(parts) > 2 {
						currentEnum.CsType = CsEnumType(parts[2])
						currentEnum.GoType = EnumTypeCsToGo(currentEnum.CsType)
					}
				}

				if parts := regexMember.FindStringSubmatch(line); len(parts) > 0 {
					kvPair := KeyValuePair{
						Name:  parts[1],
						Value: parts[2],
					}
					kvPair.Value = strings.ReplaceAll(kvPair.Value, ".", "")

					currentEnum.KeyValuePairs = append(currentEnum.KeyValuePairs, kvPair)
				}
			}
		}
	}

	enumTemplate := template.New("enum.go.tmpl")
	_, err = enumTemplate.ParseFiles("enum.go.tmpl")
	handleError(err)

	outFile, err := os.Create(outFileName)
	handleError(err)
	defer outFile.Close()

	err = enumTemplate.Execute(outFile, map[string]interface{}{
		"PackageName": pkg,
		"enums":       enums,
	})
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
