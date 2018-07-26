package main

import (
	"errors"
	"flag"
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

	v := reflect.ValueOf(flags)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i).Interface()
		if field == "" {
			return flags, errors.New("missing argument: " + strings.ToLower(v.Type().Field(i).Name))
		}
	}

	return flags, nil
}
