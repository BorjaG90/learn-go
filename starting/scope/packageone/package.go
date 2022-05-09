package packageone

import "fmt"

var privateVar = "I am private"
var PublicVar = "I am public (or exported)"
var PackageVar = "This is a Package level Variable in a package, "

func notExported() {

}

func Exported() {

}

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
