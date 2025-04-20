package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/daipham3213/selfctl/command"
	"github.com/urfave/cli/v3"
)

type SelfCtl struct {
	ctx      context.Context
	logger   *slog.Logger
	commands []command.Command
	app      *cli.Command
}

func NewSelfCtl(ctx context.Context, logger slog.Logger) *SelfCtl {
	return &SelfCtl{
		ctx:      ctx,
		logger:   &logger,
		commands: command.GetCommands(),
	}
}

func (s *SelfCtl) Run(args []string) error {
	var commands = make([]*cli.Command, len(s.commands))
	for i, cmd := range s.commands {
		commands[i] = &cli.Command{
			Name:    cmd.Name,
			Usage:   cmd.Help,
			Aliases: []string{cmd.Alias},
			Action: func(ctx context.Context, c *cli.Command) error {
				fmt.Println(cmd.Template)
				return nil
			},
		}
	}
	app := &cli.Command{
		Name:     "selfctl",
		Usage:    "a CLI tool to tell you about me",
		Commands: commands,
	}
	if err := app.Run(s.ctx, args); err != nil {
		s.logger.Error("Error running command", "error", err)
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	selfCtl := NewSelfCtl(ctx, *logger)
	if err := selfCtl.Run(os.Args); err != nil {
		logger.Error("Error running selfctl", "error", err)
		os.Exit(1)
	}
}
