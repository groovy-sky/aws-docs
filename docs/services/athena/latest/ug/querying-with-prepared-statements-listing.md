# List prepared statements using the AWS CLI

To list the prepared statements for a specific workgroup, you can use the
Athena [list-prepared-statements](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/athena/list-prepared-statements.html) AWS CLI command or the [ListPreparedStatements](../../../../reference/athena/latest/apireference/api-listpreparedstatements.md) Athena API action. The
`--work-group` parameter is required.

```nohighlight

aws athena list-prepared-statements --work-group primary
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Execute

Additional resources
