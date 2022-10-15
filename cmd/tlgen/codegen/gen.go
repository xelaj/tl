package codegen

import (
	"fmt"

	"github.com/dave/jennifer/jen"

	"github.com/xelaj/tl/schema"
)

// Some terms (id, api, url etc.) it's important to write in uppercase letters
// or lowercase only (e.g, you can't write Id, or Api, only ID, API, URL, etc.)
func capitalizePatterns() []string {
	return []string{
		"id",
		"api",
		"url",
		"p2p",
		"sha",
		"srp",
	}
}

// TODO: test it
func createCrcFunc(typ string, crc uint32) *jen.Statement {
	hex := fmt.Sprintf("0x%x", crc)
	return jen.Func().Params(jen.Id(typ)).Id("CRC").Params().Uint32().
		Id("{" + jen.Return(jen.Id(hex)).GoString() + "}")
}

func generateMethod(m *schema.ObjectMethod) *jen.Statement {
	return nil
}
