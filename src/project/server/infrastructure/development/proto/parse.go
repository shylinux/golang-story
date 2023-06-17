package proto

import (
	"os"
	"strings"

	"shylinux.com/x/golang-story/src/project/server/infrastructure/logs"
)

const (
	SYNTAX  = "syntax"
	OPTION  = "option"
	PACKAGE = "package"
	SERVICE = "service"
	MESSAGE = "message"
	METHOD  = "method"
	PARAMS  = "params"
)

type Item struct {
	Comment string `json:"omitempty"`
	Name    string `json:"omitempty"`
	List    []string
}

func (s *GenerateCmds) ParseProto(file *os.File) map[string]*Item {
	items := map[string]*Item{}
	comment, block, name := "", "", ""
	s.ScanProto(file, func(ls []string, text string) {
		if strings.HasPrefix(text, "//") {
			comment = strings.TrimSpace(strings.TrimPrefix(text, "//"))
			return
		}
		switch ls[0] {
		case "}":
			block, name = "", ""
		case SYNTAX:
		case OPTION:
		case PACKAGE:
			items[ls[0]] = &Item{Comment: comment, Name: ls[1]}
		case SERVICE:
			block, name = ls[0], ls[1]
			items[name] = &Item{Comment: comment}
			items[PACKAGE].List = append(items[PACKAGE].List, ls[1])
		case MESSAGE:
			block, name = ls[0], ls[1]
			items[name] = &Item{Comment: comment}
		default:
			switch block {
			case SERVICE:
				items[name].List = append(items[name].List, ls[1])
				items[ls[1]] = &Item{Comment: comment, List: []string{
					strings.TrimSuffix(strings.TrimPrefix(ls[2], "("), ")"),
					strings.TrimSuffix(strings.TrimPrefix(ls[4], "("), ")"),
				}}
			case MESSAGE:
				items[name].List = append(items[name].List, comment, ls[0], ls[1])
			default:
				logs.Warnf("not parse %s", text)
			}
		}
		comment = ""
	})
	return items
}
