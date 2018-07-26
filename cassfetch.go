package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gocql/gocql"
)

func main() {
	flags, err := GetCassFetchFlags()

	if err != nil {
		log.Fatal(err)
	}

	columns := strings.Split(flags.Columns, ",")
	cluster := gocql.NewCluster(flags.Host)
	cluster.Consistency = gocql.LocalQuorum
	cluster.Keyspace = flags.Keyspace
	columnFamily := flags.Cf
	keyField := flags.Keyfield

	session, sessionErr := cluster.CreateSession()

	if sessionErr != nil {
		log.Fatal("Could not connect to " + flags.Host + " with keyspace " + flags.Keyspace)
	}

	defer session.Close()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		id := scanner.Text()

		columnsString := strings.Join(columns, "\",\"")

		query := fmt.Sprintf("SELECT \"%v\" FROM %v WHERE %v=?", columnsString, columnFamily, keyField)
		result := make(map[string]interface{})

		if err := session.Query(query, id).Consistency(gocql.One).MapScan(result); err != nil {
			log.Fatal(err)
		}

		resultStringMap := make(map[string]string)

		for k, v := range result {
			resultStringMap[k] = fmt.Sprintf("%v", v)
		}

		jsonString, err := json.Marshal(resultStringMap)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonString))
	}
}
