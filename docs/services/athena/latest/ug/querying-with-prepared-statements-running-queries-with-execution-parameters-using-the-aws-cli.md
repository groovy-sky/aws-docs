# Run queries with execution parameters using the AWS CLI

To use the AWS CLI to run queries with execution parameters, use the
`start-query-execution` command and provide a parameterized query in
the `query-string` argument. Then, in the
`execution-parameters` argument, provide the values for the execution
parameters. The following example illustrates this technique.

```nohighlight

aws athena start-query-execution
--query-string "SELECT * FROM table WHERE x = ? AND y = ?"
--query-execution-context "Database"="default"
--result-configuration "OutputLocation"="s3://amzn-s3-demo-bucket;/..."
--execution-parameters "1" "2"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use the Athena console

Use prepared
statements

All content copied from https://docs.aws.amazon.com/.
