package serial

import (
	"fmt"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"os"
)

type PublishedServices struct {
	svc         []*serviceDescription
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
		fmt.Println(svc.filePath)
		svc.packages["github.com/Myriad-Dreamin/minimum-lib/controller"] = 1
		sugar.WithWriteFile(func(f *os.File) {
			_, err = fmt.Fprintf(f, `
package %s

import (
%s
)

%s

%s`, c.packageName, depList(svc.packages), svcIface(svc), svc.generateObjects())
		}, svc.filePath)
	}
	return
}

func svcIface(svc *serviceDescription) string {
	return fmt.Sprintf(`
type %s interface {
%s
}`, svc.name, svcMethods(svc))
}

func svcMethods(svc *serviceDescription) (res string) {
	res = fmt.Sprintf("    %sSignatureXXX() interface{}\n", svc.name)
	for _, cat := range svc.categories {
		for _, method := range cat.methods {
			res += "    " + method.name + "(c controller.MContext)\n"
		}
	}
	return
}

func depList(pkgSet map[string]int) (res string) {
	for k := range pkgSet {
		if len(k) > 0 {
			res += `    "` + k + `"
`
		}
	}
	return
}
