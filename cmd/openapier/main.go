package main

import (
	"bytes"
	"context"
	"encoding/json"
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
			&cli.StringFlag{
				Name:  "format",
				Usage: "Output format (yaml or json)",
				Value: "yaml",
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
			p := parser.NewParser(c.String("root"), c.String("main"))
			spec, err := p.Parse()
			if err != nil {
				return fmt.Errorf("failed to parse: %w", err)
			}

			format := c.String("format")
			var output []byte
			switch format {
			case "json":
				output, err = json.MarshalIndent(spec, "", "  ")
				if err != nil {
					return fmt.Errorf("failed to marshal spec to JSON: %w", err)
				}
			case "yaml":
				var buff bytes.Buffer
				encoder := yaml.NewEncoder(&buff)
				encoder.SetIndent(2)
				err = encoder.Encode(spec)
				if err != nil {
					return fmt.Errorf("failed to marshal spec to YAML: %w", err)
				}
				output = buff.Bytes()
			default:
				return fmt.Errorf("unsupported format: %s", format)
			}

			fmt.Println(string(output))
			return nil
		},
	}

	if err := cli.Run(context.Background(), os.Args); err != nil {
		slog.Error("fatal error running application", "err", err)
		os.Exit(1)
	}
}
