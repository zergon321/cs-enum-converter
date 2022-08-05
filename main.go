package main

import (
	"fmt"
	"regexp"
)

func main() {
	regexHeader, err := regexp.Compile(`^\s*\w*\s*enum (\w+)\s*\:*\s*(\w*)$`)
	handleError(err)

	fmt.Println(regexHeader.Match([]byte("public enum CoverageMask : uint")))
	fmt.Println(regexHeader.FindStringSubmatch("public enum CoverageMask : uint"))
	fmt.Println(regexHeader.Match([]byte("public enum CoverageMask")))
	fmt.Println(regexHeader.FindStringSubmatch("public enum CoverageMask"))

	regexMember, err := regexp.Compile(`^\s*(\w+)\s*=\s*(.+)[,;]+$`)
	handleError(err)

	fmt.Println(regexMember.Match([]byte("Unknown                 = 0x00000001,")))
	fmt.Println(regexMember.FindStringSubmatch("Unknown                 = 0x00000001,"))
	fmt.Println(regexMember.Match([]byte("Unknown                 = 0x00000001;")))
	fmt.Println(regexMember.FindStringSubmatch("Unknown                 = 0x00000001;"))
	fmt.Println(regexMember.Match([]byte("UnderwearLegs = CoverageMask.UnderwearUpperLegs | CoverageMask.UnderwearLowerLegs,")))
	fmt.Println(regexMember.FindStringSubmatch("UnderwearLegs = CoverageMask.UnderwearUpperLegs | CoverageMask.UnderwearLowerLegs,"))
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
