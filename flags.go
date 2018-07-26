package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strings"
)

// CassFetchFlags represents flags given to epoch program
type CassFetchFlags struct {
	Host     string
	Keyspace string
	Cf       string
	Columns  string
	Keyfield string
}

// GetCassFetchFlags gives an CassFetchFlags based on command line args
func GetCassFetchFlags() (CassFetchFlags, error) {
	flags := CassFetchFlags{}
	flag.StringVar(&flags.Host, "host", "", "Cassandra host")
	flag.StringVar(&flags.Keyspace, "keyspace", "", "keyspace")
	flag.StringVar(&flags.Cf, "cf", "", "Column Family")
	flag.StringVar(&flags.Columns, "columns", "", "Comma separated list of columns")
	flag.StringVar(&flags.Keyfield, "keyfield", "", "Key field (e.g. id or key)")
	flag.Parse()

	missing := missingParameters(flags)

	if len(missing) > 0 {
		message := fmt.Sprintf("Missing parameters: %s", strings.ToLower(strings.Join(missing, ", ")))
		usage(message)

		return flags, errors.New(message)
	}

	return flags, nil
}

func missingParameters(flags CassFetchFlags) (missing []string) {
	v := reflect.ValueOf(flags)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Interface()
		if field == "" {
			fieldName := v.Type().Field(i).Name
			missing = append(missing, fieldName)
		}
	}
	return
}

func usage(message string) {
	fmt.Fprintf(os.Stderr, "\n%s\n\n", message)
	flag.Usage()
}
