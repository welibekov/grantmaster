package template

import (
	"bytes"
	"context"
	"embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	//go:embed assets/*
	assetsDir embed.FS
)

type Template struct {
	path  string
	body  []byte
	pool  *pgxpool.Pool
	bytes *bytes.Buffer
}

func New(path string, pool *pgxpool.Pool) (*Template, error) {
	var templateBytes = bytes.NewBuffer([]byte{})

	if !strings.HasSuffix(path, ".tmpl") {
		path += ".tmpl"
	}

	templateBody, err := assetsDir.ReadFile(filepath.Join("assets", path))
	if err != nil {
		return nil, fmt.Errorf("couldn't read template: %v", err)
	}

	return &Template{
		path:  path,
		body:  templateBody,
		bytes: templateBytes,
		pool:  pool,
	}, nil
}

func (t *Template) Generate(ctx context.Context, parameters interface{}) ([]byte, error) {
	tmpl := template.New(t.path).Funcs(t.NewFuncs(ctx))

	tmpl, err := tmpl.Parse(string(t.body))
	if err != nil {
		return nil, fmt.Errorf("coudn't parse template: %v", err)
	}

	if err := tmpl.ExecuteTemplate(t.bytes, t.path, parameters); err != nil {
		return nil, fmt.Errorf("couldn't execute template: %v", err)
	}

	return removeEmptyLines(t.bytes.Bytes()), nil
}

// removeEmptyLines takes a []byte, splits it into lines,
// removes any that are empty (or only whitespace),
// and returns the result as a new []byte.
func removeEmptyLines(input []byte) []byte {
	// Split the input by newline.
	lines := bytes.Split(input, []byte("\n"))

	// Create a slice to hold non-empty lines.
	var nonEmptyLines [][]byte

	// Iterate over each line.
	for _, line := range lines {
		// Trim spaces to check if the line is empty.
		trimmed := bytes.TrimSpace(line)
		if len(trimmed) > 0 {
			nonEmptyLines = append(nonEmptyLines, line)
		}
	}

	// Join the non-empty lines using "\n" as the separator.
	return bytes.Join(nonEmptyLines, []byte("\n"))
}
