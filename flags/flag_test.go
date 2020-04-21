package flags_test

import (
	"flag"
	"fmt"
	"testing"

	"github.com/buguang01/util/flags"
)

func TestFlag(t *testing.T) {
	a := &FlagConfig{
		Port:    1000,
		Fdata:   1.1,
		Strname: "hello",
		Ifbool:  true,
		Arrstr:  []string{"a", "1"},
	}
	flags.SetFlagByStruct(a)
	flag.Parse()
	a.Port++
	fmt.Println(a)
	flags.LoadStruct(a)
	fmt.Println(a)
}

type FlagConfig struct {
	Port    int //说明
	Fdata   float64
	Strname string
	Ifbool  bool
	Arrstr  []string
}
