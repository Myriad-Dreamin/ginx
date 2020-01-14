package artist

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"os"
)

type PublishedServices struct {
	svc         []ServiceDescription
	packageName string
}

func (c *PublishedServices) Publish() error {
	if err := c.writeToFiles(); err != nil {
		return err
	}
	return nil
}

func (c *PublishedServices) writeToFiles() (err error) {
	if err = c.writeSVCsAndDTOs(); err != nil {
		return
	}
	return
}

func (c *PublishedServices) writeSVCsAndDTOs() (err error) {
	for i := range c.svc {
		if err != nil {
			return
		}
		svc := c.svc[i]
		//fmt.Println(svc.filePath)
		var packages = make(map[string]bool)
		packages["github.com/Myriad-Dreamin/minimum-lib/controller"] = true
		for _, cate := range svc.GetCategories() {
			packages = inplaceMergePackage(packages, cate.GetPackages())
		}
		//fmt.Println(packages)
		sugar.WithWriteFile(func(f *os.File) {
			_, err = fmt.Fprintf(f, `
package %s

import (
%s
)

%s

%s`, c.packageName, depList(packages), svcIface(svc), svc.GenerateObjects())
		}, svc.GetFilePath())
	}
	return
}

func svcIface(svc ServiceDescription) string {
	return fmt.Sprintf(`
type %s interface {
%s
}`, svc.GetName(), svcMethods(svc))
}

func svcMethods(svc ServiceDescription) (res string) {
	res = fmt.Sprintf("    %sSignatureXXX() interface{}\n", svc.GetName())
	for _, cat := range svc.GetCategories() {
		for _, method := range cat.GetMethods() {
			res += "    " + method.GetName() + "(c controller.MContext)\n"
		}
	}
	return
}

func depList(pkgSet map[string]bool) (res string) {
	for k := range pkgSet {
		if len(k) > 0 {
			res += `    "` + k + `"
`
		}
	}
	return
}
