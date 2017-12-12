package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type Environ struct {
	dir    string
	env    []string
	stdout io.Writer
	stderr io.Writer
}

func NewEnviron(dir string, env []string, stdout, stderr io.Writer) *Environ {
	return &Environ{
		dir:    dir,
		env:    env,
		stdout: stdout,
		stderr: stderr,
	}
}

// Run executes the given program.
func (e *Environ) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	cmd.Dir = e.dir
	cmd.Env = e.env
	cmd.Stdout = e.stdout
	cmd.Stderr = e.stderr

	// TODO: Extract this
	fmt.Println()
	fmt.Println("$", strings.Join(cmd.Args, " "))
	//--

	return cmd.Run()
}

const keyPath = "/tmp/gcloud.json"

func main() {
	err := wrapMain()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func wrapMain() error {
	fmt.Println("Drone Google Source Repositories Plugin start.")

	var (
		project     = getenv("PLUGIN_PROJECT")
		credentials = getenv("PLUGIN_JSON_KEY", "GOOGLE_CREDENTIALS")
	)

	if project == "" {
		return fmt.Errorf("Google Sourece Repositories Not Specified.")
	}

	credentials = strings.TrimSpace(credentials)
	if credentials == "" {
		return fmt.Errorf("Missing required parameter: google_credentials")
	}

	// Write ephemeral service account json file in plugin container
	err := ioutil.WriteFile(keyPath, []byte(credentials), 0600)
	if err != nil {
		return fmt.Errorf("error writing token file: %s", err)
	}

	defer func() {
		err := os.Remove(keyPath)
		if err != nil {
			fmt.Printf("Warning: error removing token file: %s\n", err)
		}
	}()

	e := os.Environ()
	e = append(e, fmt.Sprintf("GOOGLE_APPLICATION_CREDENTIALS=%s", keyPath))

	runner := NewEnviron(keyPath, e, os.Stdout, os.Stderr)

	err = runner.Run("gcloud", "auth", "activate-service-account", "--key-file", keyPath)
	if err != nil {
		return fmt.Errorf("Error: %s", err)
	}

	// ----> TODO: Solve error "Google Sourece Repositories Not Specified." using command below
	// err = runner.Run("gcloud", "source", "repos", "clone", "default", "default")
	// if err != nil {
	// 	return fmt.Errorf("Error: %s", err)
	// }
	// <----
	return nil
}

func getenv(key ...string) (s string) {
	for _, k := range key {
		s = os.Getenv(k)
		if s != "" {
			return
		}
	}
	return
}
