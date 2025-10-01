package starter

import (
	"errors"
	"fmt"

	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/page"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/views"
	"github.com/Broderick-Westrope/tetrigo/internal/tui"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	cfg  *config.Config
	page page.Page
}

func NewInput(cfg *config.Config, page page.Page) *Input {
	return &Input{
		cfg:  cfg,
		page: page,
	}
}

type Model struct {
	child tea.Model

	cfg          *config.Config
	forceQuitKey key.Binding

	width  int
	height int

	ExitError error
}

func NewModel(in *Input) (*Model, error) {
	m := &Model{
		cfg: in.cfg,
		forceQuitKey: key.NewBinding(
			key.WithKeys([]string{"ctrl+c"}...), // TODO: make configurable
		),
	}

	err := m.setPage(in.page)
	if err != nil {
		return nil, fmt.Errorf("setting page: %w", err)
	}
	return m, nil
}

func (m *Model) Init() tea.Cmd {
	return m.initChild()
}

func (m *Model) View() string {
	return m.child.View()
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tui.FatalErrorMsg:
		m.ExitError = msg
		return m, tea.Quit

	case tea.KeyMsg:
		if key.Matches(msg, m.forceQuitKey) {
			return m, tea.Quit
		}

	case page.SwitchPageMsg:
		err := m.setPage(msg.Target)
		if err != nil {
			return m, tui.FatalErrorCmd(fmt.Errorf("setting page: %w", err))
		}
		cmd := m.initChild()
		return m, cmd

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	var cmd tea.Cmd
	m.child, cmd = m.child.Update(msg)
	return m, cmd
}

func (m *Model) setPage(target page.Page) error {
	var child tea.Model
	var err error

	switch target {
	case page.PageMain:
		child, err = views.NewMainModel(m.cfg)
	case page.PageLobby:
		child, err = views.NewLobbyModel(m.cfg)
	case page.PageGame:
		child, err = views.NewGameModel(m.cfg)
		m.child = child
	case page.PageSettings:
		child, err = views.NewSettingsModel(m.cfg)
		m.child = child
	default:
		return errors.New("invalid page")
	}

	if err != nil {
		return fmt.Errorf("creating child model for %s page: %w", target, err)
	}
	m.child = child
	return nil
}

func (m *Model) initChild() tea.Cmd {
	var cmds []tea.Cmd
	cmd := m.child.Init()
	cmds = append(cmds, cmd)
	m.child, cmd = m.child.Update(tea.WindowSizeMsg{Width: m.width, Height: m.height})
	cmds = append(cmds, cmd)
	return tea.Batch(cmds...)
}
