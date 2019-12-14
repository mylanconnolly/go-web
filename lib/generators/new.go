package generators

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/markbates/pkger"
)

type newTmplCtx struct {
	Name    string
	Package string
}

// New is used to run the generator for a new project.
func New(pkg string) error {
	if pkg == "" {
		return fmt.Errorf("Must provide a package name")
	}
	dir := path.Base(pkg)
	_, err := os.Stat(dir)

	if !os.IsNotExist(err) {
		return fmt.Errorf("Directory %s already exists", dir)
	}
	if err = newDirectories(dir); err != nil {
		return err
	}
	file, err := pkger.Open("/templates/go.mod.tmpl")

	if err != nil {
		return err
	}
	defer file.Close()
	ctx := newTmplCtx{
		Name:    dir,
		Package: pkg,
	}
	return createFiles(dir, ctx)
}

func createFiles(dir string, ctx newTmplCtx) error {
	if err := createFile(dir, "go.mod.tmpl", ctx); err != nil {
		return err
	}
	if err := createFile(dir, "main.go.tmpl", ctx); err != nil {
		return err
	}
	return createFile(dir, "cmd/root.go.tmpl", ctx)
}

func newDirectories(dir string) error {
	if err := createDirectory(dir); err != nil {
		return err
	}
	if err := createDirectory(filepath.Join(dir, "cmd")); err != nil {
		return err
	}
	if err := createDirectory(filepath.Join(dir, "lib")); err != nil {
		return err
	}
	return nil
}

func createFile(dir, t string, ctx newTmplCtx) error {
	tmplPath := path.Join("/templates", t)
	newPath := filepath.Join(dir, strings.Replace(t, ".tmpl", "", 1))
	file, err := pkger.Open(tmplPath)

	if err != nil {
		return err
	}
	buf := bytes.Buffer{}

	if _, err := io.Copy(&buf, file); err != nil {
		return err
	}
	tmpl, err := template.New("").Parse(buf.String())

	if err != nil {
		return err
	}
	log.Printf("Creating file '%s'...", newPath)
	outFile, err := os.Create(newPath)

	if err != nil {
		return err
	}
	if err := tmpl.Execute(outFile, ctx); err != nil {
		return err
	}
	return outFile.Close()
}

func createDirectory(path string) error {
	log.Printf("Creating directory '%s'...", path)
	return os.Mkdir(path, 0755)
}
