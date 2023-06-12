package cmds

import (
	"context"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/cmds"
	"shylinux.com/x/golang-story/src/project/server/infrastructure/utils/deploy"
)

const TUI = "tui"

type TuiCmds struct {
	deploy *deploy.Deploy
}

func (s TuiCmds) Init() tea.Cmd {
	return nil
}
func (s TuiCmds) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return s, nil
}
func (s TuiCmds) View() string {
	return "hello world"
}
func (s *TuiCmds) Open(ctx context.Context, arg ...string) {
	p := tea.NewProgram(s)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
func NewTuiCmds(cmds *cmds.Cmds, deploy *deploy.Deploy) *TuiCmds {
	tui := &TuiCmds{deploy: deploy}
	cmds.Register(TUI, `tui command
  open
`, tui)
	return tui
}
