# cassfetch

Supply ids of rows in a colunfamily and `cassfetch` will output the rows as json.

e.g:

    cassfetch \
    --host example.com \
    --keyspace mykeyspace \
    --cf mycolumnfamily \
    --columns key,column1,column2 \
    --keyfield key

Might give json like this:

    {"key": "key1", "column1":"value1", "column2":"value2"}
    {"key": "key2", "column1":"value11", "column2":"value22"}

Useful to pipe to `jq`.
