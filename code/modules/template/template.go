package template

import (
	"bytes"
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

func Generate(path string, pool *pgxpool.Pool, parameters interface{}) ([]byte, error) {
	var templateBytes = bytes.NewBuffer([]byte{})

	if !strings.HasSuffix(path, ".tmpl") {
		path += ".tmpl"
	}

	fn := &Funcs{pool}

	tmpl := template.New(path)
	tmpl = tmpl.Funcs(template.FuncMap{
		"isRoleExist": fn.isRoleExist,
	})

	templateBody, err := assetsDir.ReadFile(filepath.Join("assets", path))
	if err != nil {
		return nil, fmt.Errorf("couldn't read template: %v", err)
	}

	tmpl, err = tmpl.Parse(string(templateBody))
	if err != nil {
		return nil, fmt.Errorf("coudn't parse template: %v", err)
	}

	if err := tmpl.ExecuteTemplate(templateBytes, path, parameters); err != nil {
		return nil, fmt.Errorf("couldn't execute template: %v", err)
	}

	return removeEmptyLines(templateBytes.Bytes()), nil
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
