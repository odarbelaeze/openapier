package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/odarbelaeze/openapier/pkg/parser"
	"github.com/urfave/cli/v3"
	"go.yaml.in/yaml/v4"
)

func main() {
	cli := cli.Command{
		Name:  "openapier",
		Usage: "A tool to generate OpenAPI specifications from Go code",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:      "main",
				Usage:     "Path to the file containing the root spec definition.",
				Value:     "main.go",
				TakesFile: true,
			},
			&cli.StringFlag{
				Name:        "root",
				Usage:       "Path to the root directory of the Go code to parse. ",
				Value:       "./",
				DefaultText: "current working directory",
				TakesFile:   false,
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
		},
		Action: func(ctx context.Context, c *cli.Command) error {
			if c.Bool("debug") {
				slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug})))
			}
			p := parser.NewParser()
			spec, err := p.Parse(c.String("root"), c.String("main"))
			if err != nil {
				return fmt.Errorf("failed to parse: %w", err)
			}
			bytes, err := yaml.Marshal(spec)
			fmt.Println(string(bytes))
			return err
		},
	}

	if err := cli.Run(context.Background(), os.Args); err != nil {
		slog.Error("fatal error running application", "err", err)
		os.Exit(1)
	}
}
