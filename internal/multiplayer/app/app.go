package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
	"github.com/muesli/termenv"
)

const (
	host = "localhost"
	port = "23234"
)

type progID int

// App contains a wish server and the list of running programs.
type App struct {
	server *ssh.Server
	progs  map[progID]*tea.Program
}

func New() *App {
	a := &App{
		progs: make(map[progID]*tea.Program),
	}
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPath(".ssh/id_ed25519"),
		wish.WithMiddleware(
			bubbletea.MiddlewareWithProgramHandler(a.ProgramHandler, termenv.ANSI256),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Error("Could not start server", "error", err)
	}

	a.server = s
	return a
}

func (a *App) Start() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	log.Info("Starting SSH server", "host", host, "port", port)
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			log.Error("Could not start server", "error", err)
			done <- nil
		}
	}()

	<-done
	log.Info("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := a.server.Shutdown(ctx); err != nil {
		log.Error("Could not stop server", "error", err)
	}
}

func (a *App) ProgramHandler(s ssh.Session) *tea.Program {
	model := initialModel()
	model.app = a
	model.id = s.User()

	p := tea.NewProgram(model, bubbletea.MakeOptions(s)...)
	a.progs[progID(len(a.progs))] = p

	return p
}

func (a *App) send(pid progID, msg tea.Msg) {
	if p, ok := a.progs[pid]; ok {
		p.Send(msg)
	}
}
