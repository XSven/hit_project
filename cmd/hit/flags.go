package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

type flags struct {
	url  string
	n, c int
}

func (f *flags) parse() error {
	flag.StringVar(&f.url, "url", "", "HTTP server `URL` to make requests (required)")
	flag.IntVar(&f.n, "n", f.n, "Number of requests to make")
	flag.IntVar(&f.c, "c", f.c, "Concurrency level")
	flag.Parse()
	if err := f.validate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		return err
	}
	return nil
}

// validate post-conditions after parsing the flags.
func (f *flags) validate() error {
	if f.c > f.n {
		return fmt.Errorf("-c=%d: should be less than or equal to -n=%d", f.c, f.n)
	}
	if err := validateURL(f.url); err != nil {
		return fmt.Errorf("invalid value %q for flag -url: %w", f.url, err)
	}
	return nil
}

func validateURL(s string) error {
	if strings.TrimSpace(s) == "" {
		return errors.New("required")
	}
	_, err := url.Parse(s)
	return err
}
