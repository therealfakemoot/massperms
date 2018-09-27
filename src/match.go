package perminator

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

// Match probably doesn't need to exist. All it does is call filepath.Match
func Match(pattern, name string) (bool, error) {
	return filepath.Match(pattern, name)
}

// Apply is a HOF that creates a filepath.WalkFunc that applies the provided ruleset to a target directory.
func Apply(rules RuleSet, targetDir string) filepath.WalkFunc {
	f := func(fname string, info os.FileInfo, err error) error {
		log.Printf("walking over %s", fname)
		for _, r := range rules {
			pattern := path.Join(targetDir, r.Pattern)
			log.Printf("matching against pattern: %s", pattern)
			m, err := Match(pattern, fname)
			if err != nil {
				return err
			}
			if m {
				log.Printf("updating permissions for %s: %s", fname, r.Mode)
				err := os.Chmod(fname, r.Mode)
				if err != nil {
					log.Printf("unable to modify %s: %s", fname, err)
					return err
				}
			}
		}
		return nil
	}

	return f
}
