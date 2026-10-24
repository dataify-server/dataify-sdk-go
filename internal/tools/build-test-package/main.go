// Command build-test-package builds distributable artifacts for this Go module.
//
// It writes:
//
//	artifacts/dataify-sdk-go-<version>.zip              plain source archive
//	artifacts/goproxy/<module>/@v/<version>.zip         Go module proxy zip
//	artifacts/goproxy/<module>/@v/<version>.mod         Go module proxy go.mod
//	artifacts/goproxy/<module>/@v/<version>.info        Go module proxy info
//	artifacts/SHA256SUMS.txt
//	artifacts/README-test.md
//
// Run from the repository root:
//
//	go run ./internal/tools/build-test-package
package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"
)

const defaultVersion = "v0.2.0-test.4"

var versionPattern = regexp.MustCompile(`^v\d+\.\d+\.\d+(?:-[0-9A-Za-z.-]+)?(?:\+[0-9A-Za-z.-]+)?$`)

type fileEntry struct {
	rel string
	abs string
}

type artifact struct {
	path string
	name string
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "build-test-package: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	version := flag.String("version", defaultVersion, "module version, for example v0.2.0-test.1")
	moduleFlag := flag.String("module", "", "module path; defaults to the module directive in go.mod")
	outFlag := flag.String("out", "artifacts", "output directory")
	rootFlag := flag.String("root", ".", "module root directory")
	flag.Parse()

	if !versionPattern.MatchString(*version) {
		return fmt.Errorf("invalid module version %q: expected a semantic version such as v0.2.0-test.1", *version)
	}

	root, err := filepath.Abs(*rootFlag)
	if err != nil {
		return err
	}
	modulePath := strings.TrimSpace(*moduleFlag)
	goModPath := filepath.Join(root, "go.mod")
	goModBytes, err := os.ReadFile(goModPath)
	if err != nil {
		return fmt.Errorf("read go.mod: %w", err)
	}
	if modulePath == "" {
		modulePath = moduleDirective(goModBytes)
	}
	if modulePath == "" {
		return errors.New("module path is empty; pass -module or add a module directive to go.mod")
	}

	outDir, err := filepath.Abs(filepath.Join(root, *outFlag))
	if err != nil {
		return err
	}
	if !withinDir(root, outDir) {
		return fmt.Errorf("output directory %s must be inside the module root %s", outDir, root)
	}
	if err := os.RemoveAll(outDir); err != nil {
		return fmt.Errorf("clean output directory: %w", err)
	}
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	allFiles, moduleFiles, err := collectFiles(root, outDir)
	if err != nil {
		return err
	}
	if len(allFiles) == 0 {
		return errors.New("no files to package")
	}

	baseName := filepath.Base(root)
	versionName := strings.TrimPrefix(*version, "v")
	modTime := latestFileTime(allFiles)

	plainZip := filepath.Join(outDir, fmt.Sprintf("%s-%s.zip", baseName, versionName))
	if err := writeZip(plainZip, baseName+"-"+versionName+"/", allFiles, modTime); err != nil {
		return err
	}

	proxyDir := filepath.Join(outDir, "goproxy", filepath.FromSlash(modulePath), "@v")
	if err := os.MkdirAll(proxyDir, 0o755); err != nil {
		return err
	}
	moduleZip := filepath.Join(proxyDir, *version+".zip")
	if err := writeZip(moduleZip, modulePath+"@"+*version+"/", moduleFiles, modTime); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(proxyDir, *version+".mod"), goModBytes, 0o644); err != nil {
		return err
	}
	infoBytes, err := json.Marshal(struct {
		Version string    `json:"Version"`
		Time    time.Time `json:"Time"`
	}{Version: *version, Time: modTime})
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(proxyDir, *version+".info"), append(infoBytes, '\n'), 0o644); err != nil {
		return err
	}

	artifacts := []artifact{
		{path: plainZip},
		{path: moduleZip},
		{path: filepath.Join(proxyDir, *version+".mod")},
		{path: filepath.Join(proxyDir, *version+".info")},
	}
	if err := writeChecksums(filepath.Join(outDir, "SHA256SUMS.txt"), outDir, artifacts); err != nil {
		return err
	}

	readmePath := filepath.Join(outDir, "README-test.md")
	if err := os.WriteFile(readmePath, []byte(testReadme(modulePath, *version, baseName, versionName)), 0o644); err != nil {
		return err
	}

	fmt.Printf("built %s %s\n", modulePath, *version)
	fmt.Printf("  plain archive: %s\n", plainZip)
	fmt.Printf("  file proxy:    %s\n", filepath.Join(outDir, "goproxy"))
	return nil
}

func moduleDirective(goMod []byte) string {
	for _, line := range strings.Split(string(goMod), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(strings.TrimPrefix(line, "module "), "\t"))
		}
	}
	return ""
}

func collectFiles(root string, outDir string) ([]fileEntry, []fileEntry, error) {
	var allFiles []fileEntry
	var moduleFiles []fileEntry
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == outDir {
			return fs.SkipDir
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		rel = filepath.ToSlash(rel)
		if rel == "." {
			return nil
		}
		if d.IsDir() {
			if skipDir(d.Name()) {
				return fs.SkipDir
			}
			return nil
		}
		if !d.Type().IsRegular() {
			return nil
		}
		if skipFile(rel) {
			return nil
		}
		entry := fileEntry{rel: rel, abs: path}
		allFiles = append(allFiles, entry)
		if !dotPrefixed(rel) {
			moduleFiles = append(moduleFiles, entry)
		}
		return nil
	})
	sort.Slice(allFiles, func(i, j int) bool { return allFiles[i].rel < allFiles[j].rel })
	sort.Slice(moduleFiles, func(i, j int) bool { return moduleFiles[i].rel < moduleFiles[j].rel })
	return allFiles, moduleFiles, err
}

func withinDir(root string, target string) bool {
	rel, err := filepath.Rel(filepath.Clean(root), filepath.Clean(target))
	if err != nil {
		return false
	}
	if rel == "." || rel == ".." {
		return false
	}
	return !strings.HasPrefix(rel, ".."+string(os.PathSeparator))
}

func latestFileTime(files []fileEntry) time.Time {
	latest := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	for _, entry := range files {
		info, err := os.Stat(entry.abs)
		if err == nil && info.ModTime().After(latest) {
			latest = info.ModTime().UTC()
		}
	}
	return latest.Truncate(time.Second)
}

func skipDir(name string) bool {
	switch name {
	case ".git", ".idea", ".gocache", ".gotmp", "artifacts", "dist", "node_modules":
		return true
	}
	return false
}

func skipFile(rel string) bool {
	base := filepath.Base(rel)
	switch {
	case base == ".DS_Store", base == "Thumbs.db":
		return true
	case strings.HasPrefix(base, ".env"):
		return true
	case strings.HasPrefix(rel, "docs/ts-go-diff."):
		return true
	}
	switch strings.ToLower(filepath.Ext(base)) {
	case ".exe", ".test", ".out", ".dll", ".so", ".dylib", ".log", ".tmp", ".pem", ".key",
		".zip", ".tgz", ".gz", ".tar", ".whl", ".jar":
		return true
	}
	return false
}

func dotPrefixed(rel string) bool {
	for _, part := range strings.Split(rel, "/") {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

func writeZip(path string, prefix string, files []fileEntry, modTime time.Time) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	writer := zip.NewWriter(file)
	for _, entry := range files {
		header := &zip.FileHeader{
			Name:   prefix + entry.rel,
			Method: zip.Deflate,
		}
		header.SetMode(0o644)
		header.SetModTime(modTime)
		part, err := writer.CreateHeader(header)
		if err != nil {
			writer.Close()
			file.Close()
			return err
		}
		source, err := os.Open(entry.abs)
		if err != nil {
			writer.Close()
			file.Close()
			return err
		}
		if _, err := io.Copy(part, source); err != nil {
			source.Close()
			writer.Close()
			file.Close()
			return err
		}
		if err := source.Close(); err != nil {
			writer.Close()
			file.Close()
			return err
		}
	}
	if err := writer.Close(); err != nil {
		file.Close()
		return err
	}
	return file.Close()
}

func writeChecksums(path string, baseDir string, artifacts []artifact) error {
	lines := make([]string, 0, len(artifacts))
	for _, item := range artifacts {
		sum, err := fileSHA256(item.path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(baseDir, item.path)
		if err != nil {
			return err
		}
		lines = append(lines, fmt.Sprintf("%s  %s", sum, filepath.ToSlash(rel)))
	}
	sort.Strings(lines)
	return os.WriteFile(path, []byte(strings.Join(lines, "\n")+"\n"), 0o644)
}

func fileSHA256(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func testReadme(modulePath string, version string, baseName string, versionName string) string {
	return fmt.Sprintf(`# Test package: %s %s

This directory contains a local, non-published build of the Dataify Go SDK.
It is meant for test environments only.

## Contents

- %s-%s.zip: plain source archive. Unzip it and add a local replace directive.
- goproxy/: standard Go module proxy layout for %s@%s.
- SHA256SUMS.txt: checksums for this build.

## Option 1: local replace

    Expand-Archive -LiteralPath .\%s-%s.zip -DestinationPath .

Then add to the consumer project:

    replace %s => ../%s-%s

## Option 2: local file proxy

    $env:GOPROXY = "file:///<abs-path-to-artifacts>/goproxy"
    $env:GOSUMDB = "off"
    go get %s@%s

Replace the file URL with the absolute path to this artifacts directory.

## AI Web configuration

AI Web defaults to the production gateway https://aiwebapi.dataify.com.
To use a test gateway, set AI_WEB_API_BASE_URL before creating the client:

    $env:AI_WEB_API_BASE_URL = "https://<test-ai-web-gateway>"
    $env:DATAIFY_API_TOKEN = "<test-token>"

or pass it explicitly:

    client := dataify.NewClient(
        dataify.TokenFromEnv(),
        dataify.WithAIWebBaseURL("https://<test-ai-web-gateway>"),
    )

Do not commit real test tokens.

## Rebuild

From the repository root:

    go run ./internal/tools/build-test-package -version %s
`, modulePath, version, baseName, versionName, modulePath, version, baseName, versionName, modulePath, baseName, versionName, modulePath, version, version)
}
