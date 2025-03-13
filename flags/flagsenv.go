package flags

import (
	"flag"
	"os"
	"strings"

	"github.com/buguang01/util"
)

var (
	flagmap map[string]*FlagModel
	envmap  map[string]*FlagModel
)

func init() {
	flagmap = make(map[string]*FlagModel)
	envmap = make(map[string]*FlagModel)
}

type FlagModel struct {
	FlagName string
	EnvName  string
	Usage    string
	Val      string
	Default  string
}

// 设置要读的参数
// flagname 命令行参数
// envname 环境参数
// note 命令行描述信息
// def 默认值
func SetFlag(flagname, envname, note, def string) {
	m := &FlagModel{
		FlagName: flagname,
		EnvName:  envname,
		Usage:    note,
		Default:  def,
		Val:      def,
	}
	if m.FlagName != "" {
		flagmap[m.FlagName] = m
		flag.StringVar(&m.Val, m.FlagName, m.Default, m.Usage)
	}
	if m.EnvName != "" {
		envmap[m.EnvName] = m
	}
}

func Parse() {
	for _, m := range envmap {
		if v, ok := os.LookupEnv(m.EnvName); ok {
			m.Val = v
		}
	}
	flag.Parse()
}

func GetFlagByInt(name string) int {
	if m, ok := flagmap[name]; ok {
		return util.NewString(m.Val).ToIntV()
	}
	return 0
}

func GetFlagByUint(name string) uint {
	if m, ok := flagmap[name]; ok {
		return uint(util.NewString(m.Val).ToUint64V())
	}
	return 0
}

func GetFlagByString(name string) string {
	if m, ok := flagmap[name]; ok {
		return m.Val
	}
	return ""
}

func GetFlagByFloat64(name string) float64 {
	if m, ok := flagmap[name]; ok {
		return util.NewString(m.Val).ToFloatV()
	}
	return 0
}

func GetFlagBySlice(name string) []string {
	if m, ok := flagmap[name]; ok {
		m.Val = strings.ReplaceAll(m.Val, " ", ",")
		return strings.Split(m.Val, ",")
	}
	return []string{}
}

func GetFlagByBool(name string) bool {
	if m, ok := flagmap[name]; ok {
		return util.NewString(m.Val).ToBoolV()
	}
	return false
}

func GetEnvByInt(name string) int {
	if m, ok := envmap[name]; ok {
		return util.NewString(m.Val).ToIntV()
	}
	return 0
}

func GetEnvByString(name string) string {
	if m, ok := envmap[name]; ok {
		return m.Val
	}
	return ""
}

func GetEnvByFloat64(name string) float64 {
	if m, ok := envmap[name]; ok {
		return util.NewString(m.Val).ToFloatV()
	}
	return 0
}

func GetEnvBySlice(name string) []string {
	if m, ok := envmap[name]; ok {
		m.Val = strings.ReplaceAll(m.Val, " ", ",")
		return strings.Split(m.Val, ",")
	}
	return []string{}
}

func GetEnvByBool(name string) bool {
	if m, ok := envmap[name]; ok {
		return util.NewString(m.Val).ToBoolV()
	}
	return false
}
