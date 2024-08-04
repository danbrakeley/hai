package main

import (
	"bytes"
	"strings"
	"time"

	"github.com/danbrakeley/bsh"
	"github.com/magefile/mage/mg"
)

var sh = &bsh.Bsh{}
var hai = "hai"

// Test runs tests for all packages
func Test() {
	sh.Echo("Running unit tests...")
	sh.Cmd("go test ./...").Run()
}

// Gen runs go generate for all packages
func Gen() {
	sh.Echo("Running go generate...")
	sh.Cmd("go generate ./...").Run()
}

// BuildHai builds cmd/hai (output goes to "local" folder)
func BuildHai() {
	target := sh.ExeName(hai)

	sh.Echof("Building %s...", target)
	sh.MkdirAll("local/")

	// grab git commit hash to use as version for local builds
	commit := "(dev)"
	var b bytes.Buffer
	n := sh.Cmd(`git log --pretty=format:'%h' -n 1`).Out(&b).RunExitStatus()
	if n == 0 {
		commit = strings.TrimSpace(b.String())
	}

	sh.Cmdf(
		`go build -ldflags '`+
			`-X "github.com/danbrakeley/hai/internal/buildvar.Version=%s" `+
			`-X "github.com/danbrakeley/hai/internal/buildvar.BuildTime=%s" `+
			`-X "github.com/danbrakeley/hai/internal/buildvar.ReleaseURL=https://github.com/danbrakeley/hai"`+
			`' -o local/%s ./cmd/%s`, commit, time.Now().Format(time.RFC3339), target, hai,
	).Run()
}

// Build tests and builds all apps
func Build() {
	mg.SerialDeps(Test, BuildHai)
}

// RunWaiig runs unit tests, builds hai until /local, then runs it
func RunWaiig() {
	mg.SerialDeps(Test, BuildHai)

	target := sh.ExeName(hai)

	sh.Echo("Running...")
	sh.Cmdf("./%s", target).Dir("local").Run()
}

// Setup installs cli apps needed for development (not including 'go' or 'mage')
func Setup() {
	sh.Echo("Installing enumer...")
	sh.Cmd("go install github.com/dmarkham/enumer@latest").Run()
}

// CI runs all CI tasks
func CI() {
	mg.SerialDeps(Setup, Build)
}
