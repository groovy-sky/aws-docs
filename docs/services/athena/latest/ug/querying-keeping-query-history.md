# Keep your query history longer than 45 days

If you want to keep the query history longer than 45 days, you can retrieve the query
history and save it to a data store such as Amazon S3. To automate this process, you can use
Athena and Amazon S3 API actions and CLI commands. The following procedure summarizes these
steps.

###### To retrieve and save query history programmatically

1. Use Athena [ListQueryExecutions](../../../../reference/athena/latest/apireference/api-listqueryexecutions.md) API action or the [list-query-executions](../../../cli/latest/reference/athena/list-query-executions.md) CLI command to retrieve the query IDs.

2. Use the Athena [GetQueryExecution](../../../../reference/athena/latest/apireference/api-getqueryexecution.md) API action or the [get-query-execution](../../../cli/latest/reference/athena/get-query-execution.md) CLI command to retrieve information about each
    query based on its ID.

3. Use the Amazon S3 [PutObject](../../../s3/latest/api/api-putobject.md) API
    action or the [put-object](../../../cli/latest/reference/s3api/put-object.md) CLI command to save the information in Amazon S3.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configure recent query options

Find query output files in Amazon S3

All content copied from https://docs.aws.amazon.com/.
