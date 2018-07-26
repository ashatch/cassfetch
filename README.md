# cassfetch

Fetch rows from Cassandra using CQL. This tool assumes you are supplying primary keys as ids
over `stdin`.

## Huh?

Supply ids of rows in a columnfamily on `stdin` and `cassfetch` will output the rows
as json, e.g:

    cat my-list-of-ids.txt |
    cassfetch \
    --host example.com \
    --keyspace mykeyspace \
    --cf mycolumnfamily \
    --columns key,column1,column2 \
    --keyfield key


In the above example, if `my-list-of-ids.txt` contained

    key1
    key2
    
then the output might give this json:

    {"key": "key1", "column1":"value1", "column2":"value2"}
    {"key": "key2", "column1":"value11", "column2":"value22"}

Useful to pipe to `jq`.
