# List prepared statements using the AWS CLI

To list the prepared statements for a specific workgroup, you can use the
Athena [list-prepared-statements](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/athena/list-prepared-statements.html) AWS CLI command or the [ListPreparedStatements](../../../../reference/athena/latest/apireference/api-listpreparedstatements.md) Athena API action. The
`--work-group` parameter is required.

```nohighlight

aws athena list-prepared-statements --work-group primary
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Execute

Additional resources

All content copied from https://docs.aws.amazon.com/.
