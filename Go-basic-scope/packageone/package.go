package packageone

import "fmt"

var PackageVar string = "PackageVar-variable is a package level from packageone"

//note if function is a lowwer cap, it only available in a package where're declare
func Printme(x1, x2 string) {
	fmt.Println(x1, x2, PackageVar)

}
