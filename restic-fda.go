// https://github.com/restic/restic/issues/2051#issuecomment-440477438
// https://n8henrie.com/2018/11/how-to-give-full-disk-access-to-a-binary-in-macos-mojave/

// Runrestic provides a binary to run my restic backup script in MacOS Mojave with Full Disk Access
package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(ex)
	script := filepath.Join(dir, "restic-backup.sh")

	cmd := exec.Command("/bin/bash", script)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
