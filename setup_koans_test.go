package go_koans

import (
	"fmt"
	"testing"
)

func TestAboutArrays(t *testing.T) {
	aboutArrays()
}

func TestAboutBasics(t *testing.T) {
	aboutBasics()
}

func TestAboutStrings(t *testing.T) {
	AboutStrings()
}

func TestKoans(t *testing.T) {
	aboutSlices()
	aboutTypes()
	aboutControlFlow()
	aboutEnumeration()
	aboutAnonymousFunctions()
	aboutVariadicFunctions()
	aboutFiles()
	aboutInterfaces()
	aboutCommonInterfaces()
	aboutMaps()
	aboutPointers()
	aboutStructs()
	aboutAllocation()
	aboutChannels()
	aboutConcurrency()
	aboutPanics()

	fmt.Printf("\n%c[32;1mYou won life. Good job.%c[0m\n\n", 27, 27)
}
