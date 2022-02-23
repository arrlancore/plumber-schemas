// This script will generate a helper func for generating connection options.
// This func is used by the CLI to avoid having to write out a switch manually.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

const (
	tunnelGrepFile   = "build/go/protos/opts/ps_opts_tunnel.pb.go"
	tunnelOutputFile = "./build/go/protos/opts/batch_merge_relay_opts.pb.go"
)

type TunnelMapping struct {
	Name string
	Capitalized string
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("unable to get working dir: %s", err)
	}

	grepArgs := [][]string {
		{`\*TunnelGroup.\+protobuf:`, path + "/"  + tunnelGrepFile},
	}

	// key = backend name
	mappings := make(map[string]TunnelMapping, 0)

	for _, cmd := range grepArgs {
		out, err := exec.Command("grep", cmd ... ).CombinedOutput()
		if err != nil {
			log.Fatalf("unable to exec grep\nOutput: %s\nError: %s", string(out), err)
		}

		lines := strings.Split(string(out), "\n")

		for _, line := range lines {
			if len(line) == 0 {
				continue
			}

			r, err := regexp.Compile(`\t?([a-zA-Z0-9]+)\s+\*TunnelGroup(.+)Options\s+.+`)
			if err != nil {
				log.Fatalf("unable to compile regex: %s", err)
			}

			found := r.FindStringSubmatch(line)

			if len(found) != 3 {
				log.Fatalf("unexpected number of regex matches '%d'", len(found))
			}

			normalizedBackendName := strings.ToLower(found[1])
			normalizedBackendName = strings.Replace(normalizedBackendName, " ", "", -1)

			bm := TunnelMapping{
				Name:     found[1],
				Capitalized: found[2],
			}

			mappings[normalizedBackendName] = bm
		}
	}

	if len(mappings) == 0 {
		log.Fatalf("something went wrong - 0 mappings found for relays")
	}

	fileContents, err := createTunnelCode(mappings)
	if err != nil {
		log.Fatalf("unable to create file contents: %s", err)
	}

	fmt.Println(string(fileContents))

	//if err := writeTunnelCode(tunnelOutputFile, fileContents); err != nil {
	//	log.Fatalf("unable to write file '%s': %s", tunnelOutputFile, err)
	//}

	fmt.Printf("Successfully generated file '%s' with '%d' relays\n", tunnelOutputFile, len(mappings))
}

func createTunnelCode(mappings map[string]TunnelMapping) ([]byte, error) {
	contents := `// Code generated by generate-merge-relay-opts.go. DO NOT EDIT.

package opts

import (
	"errors"
)

func MergeTunnelOptions(backend string, tunnelOpts *TunnelOptions, createTunnelOpts *CreateTunnelOptions) error {
	switch backend {
`
	// We want to have the same generation order (to not trigger repo updates
	keys := make([]string, 0)

	for backendName, _ := range mappings {
		keys = append(keys, backendName)
	}

	sort.Strings(keys)

	for _, key := range keys {
		mapping, ok := mappings[key]
		if !ok {
			log.Fatalf("unable to find mapping for backend '%s'", key)
		}

		caseEntry := fmt.Sprintf("\tcase \"%s\":\n", key)
		caseEntry += fmt.Sprintf("\t\ttunnelOpts.%s = &TunnelGroup%sOptions{Args: createTunnelOpts.%s}\n", mapping.Name, mapping.Capitalized, mapping.Name)

		contents += caseEntry
	}

	// Close the switch and func
	contents += fmt.Sprint("\tdefault:\n\t\treturn errors.New(\"unknown backend\")\n\t}\n\n")
	contents += fmt.Sprint("\treturn nil\n}\n")

	return []byte(contents), nil
}

func writeTunnelCode(filename string, contents []byte) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("unable to create file '%s': %s", filename, err)
	}

	if _, err := f.Write(contents); err != nil {
		return fmt.Errorf("unable to write to file '%s': %s", filename, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("unable to close file '%s': %s", filename, err)
	}

	return nil
}