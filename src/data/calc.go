package data

import (
	"strings"

	"shylinux.com/x/ice"
	kit "shylinux.com/x/toolkits"
)

type calc struct {
	ice.Code
	ice.Hash
	short string `data:"input"`
	field string `data:"time,hash,input,output"`

	create string `name:"create input:textarea=1+1" help:"创建"`
	list   string `name:"list hash auto create" help:"计算器"`
}

func (s calc) Create(m *ice.Message, arg ...string) {
	op := map[string]int{"*": 3, "/": 3, "%": 3, "+": 2, "-": 2, "(": 1, ")": 1}
	stack := []string{}
	ops := func() bool {
		n := len(stack)
		if n < 3 {
			return false
		}
		res, a, op, b := 0, stack[n-3], stack[n-2], stack[n-1]
		switch op {
		case "*":
			res = kit.Int(a) * kit.Int(b)
		case "/":
			res = kit.Int(a) / kit.Int(b)
		case "%":
			res = kit.Int(a) % kit.Int(b)
		case "+":
			res = kit.Int(a) + kit.Int(b)
		case "-":
			res = kit.Int(a) - kit.Int(b)
		}
		stack = append(stack[:n-3], kit.Format(res))
		m.Debug("%v %s %s %s", stack, a, op, b)
		return true
	}
	for _, v := range kit.Split(m.Option("input"), "\t \n", strings.Join(kit.SortedKey(op), "")) {
		m.Debug("what %v %v", stack, v)
		if v == ")" {
			for ops() {
				if n := len(stack); stack[n-2] == "(" {
					stack = append(stack[:n-2], stack[n-1])
					break
				}
			}
			continue
		}
		if op[v] > 0 && len(stack) > 2 && op[stack[len(stack)-2]] >= op[v] {
			ops()
		}
		stack = append(stack, v)
	}
	for ops() {
	}
	s.Hash.Create(m, kit.Simple(m.OptionSimple("input"), "output", stack[0])...)
}
func (s calc) List(m *ice.Message, arg ...string) {
	s.Hash.List(m, arg...)
}

func init() { ice.CodeCtxCmd(calc{}) }
