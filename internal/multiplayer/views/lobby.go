package views

import (
	"fmt"
	"time"

	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/colors"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/config"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/layout"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/page"
	"github.com/Broderick-Westrope/tetrigo/internal/multiplayer/validators"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type connectionStatus int

const (
	csNone connectionStatus = iota
	csConnecting
	csSuccess
	csFailure
)

var _ tea.Model = &LobbyModel{}
var _ page.Sizable = &LobbyModel{}

type LobbyModel struct {
	page.SizeableImpl

	form    *huh.Form
	spinner spinner.Model
	keys    *lobbyKeyMap

	formData      *lobbyFormData
	formCompleted bool

	connectionStatus
}

type lobbyFormData struct {
	username string
	isHost   bool
	hostAddr string
	port     string
	password string
}

type startConnectionMsg struct{}

func defaultLobbyFormData() *lobbyFormData {
	return &lobbyFormData{
		username: "tetrigo",
		isHost:   true,
		hostAddr: "127.0.0.1", // Default to localhost
		port:     "22",        // Default to SSH
	}
}

func defaultSpinner() spinner.Model {
	spn := spinner.New()
	spn.Spinner = spinner.Points
	spn.Style = lipgloss.NewStyle().Foreground(colors.Violet)
	return spn
}

func NewLobbyModel(_ *config.Config) (*LobbyModel, error) {
	formData := defaultLobbyFormData()
	keys := defaultLobbyKeyMap()
	spinner := defaultSpinner()

	return &LobbyModel{
		formData: formData,
		form: huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title("Host or Join?").
					Affirmative("Host").
					Negative("Join").
					Value(&formData.isHost),
				huh.NewInput().
					Value(&formData.username).
					Title("Username (Max 20):").
					CharLimit(20).
					Validate(validators.Username()),
				huh.NewInput().
					Value(&formData.hostAddr).
					Title("Host Address:").
					Description("IPv4 or IPv6 address").
					CharLimit(50).
					Validate(validators.HostAddr()),
				huh.NewInput().
					Value(&formData.port).
					Title("Port:").
					CharLimit(5).
					Validate(validators.Port()),
				huh.NewInput().
					Value(&formData.password).
					Title("Password (Max 20):").
					Description("Leave blank for no password.").
					CharLimit(20).
					EchoMode(huh.EchoModePassword),
			),
		).WithKeyMap(keys.formKeys),
		keys:    keys,
		spinner: spinner,
	}, nil
}

// Init implements tea.Model.
func (m *LobbyModel) Init() tea.Cmd {
	return m.form.Init()
}

func (m *LobbyModel) SetSize(width, height int) {
	m.SizeableImpl.SetSize(width, height)
	formWidth := width / 2
	formWidth = min(formWidth, lipgloss.Width(titleStr))
	m.form = m.form.WithWidth(formWidth)
}

// Update implements tea.Model.
func (m *LobbyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, m.keys.Exit) {
			return m, page.SwitchPageCmd(page.PageMain)
		}
	case startConnectionMsg:
		cmds := []tea.Cmd{m.spinner.Tick, m.startConnection()}
		return m, tea.Batch(cmds...)
	case spinner.TickMsg:
		spn, cmd := m.spinner.Update(msg)
		m.spinner = spn
		return m, cmd
	}

	var cmds []tea.Cmd
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted && !m.formCompleted {
		cmds = append(cmds, m.announceCompletion())
	}

	return m, tea.Batch(cmds...)
}

func (m *LobbyModel) announceCompletion() tea.Cmd {
	m.formCompleted = true
	return func() tea.Msg {
		return startConnectionMsg{}
	}
}

func (m *LobbyModel) startConnection() tea.Cmd {
	// TODO: Temporary, need to seperate connection logic from UI
	return func() tea.Msg {
		m.connectionStatus = csConnecting
		time.Sleep(time.Second * 3)
		m.connectionStatus = csSuccess
		time.Sleep(time.Second * 1)
		return page.SwitchPageMsg{Target: page.PageGame}
	}
}

// View implements tea.Model.
func (m *LobbyModel) View() string {
	output := lipgloss.JoinVertical(lipgloss.Center,
		titleStr,
		layout.GapV(1),
		lipgloss.NewStyle().
			Render("[Lobby]"),
		layout.GapV(1),
		m.form.View(),
		m.connectionView(),
	)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, output)
}

func (m *LobbyModel) connectionView() string {
	if !m.formCompleted {
		return ""
	}
	switch m.connectionStatus {
	case csNone:
		return ""
	case csConnecting:
		var desc string
		if m.formData.isHost {
			desc = "Hosting on"
		} else {
			desc = "Connecting to"
		}
		return lipgloss.JoinVertical(
			lipgloss.Center,
			fmt.Sprintf("%s [%s:%s]", desc, m.formData.hostAddr, m.formData.port),
			layout.GapV(1),
			m.spinner.View(),
		)
	case csSuccess:
		return "Connection Successful. Entering game..."
	case csFailure:
		return "Connection Failed"
	}
	return "Connection Status Unknown"
}
