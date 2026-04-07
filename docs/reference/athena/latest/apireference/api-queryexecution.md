# QueryExecution

Information about a single instance of a query execution.

## Contents

**EngineVersion**

The engine version that executed the query.

Type: [EngineVersion](api-engineversion.md) object

Required: No

**ExecutionParameters**

A list of values for the parameters in a query. The values are applied sequentially to
the parameters in the query in the order in which the parameters occur. The list of
parameters is not returned in the response.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**ManagedQueryResultsConfiguration**

The configuration for storing results in Athena owned storage, which includes whether
this feature is enabled; whether encryption configuration, if any, is used for
encrypting query results.

Type: [ManagedQueryResultsConfiguration](api-managedqueryresultsconfiguration.md) object

Required: No

**Query**

The SQL query statements which the query execution ran.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 262144.

Required: No

**QueryExecutionContext**

The database in which the query execution occurred.

Type: [QueryExecutionContext](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionContext.html) object

Required: No

**QueryExecutionId**

The unique identifier for each query execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: No

**QueryResultsS3AccessGrantsConfiguration**

Specifies whether Amazon S3 access grants are enabled for query
results.

Type: [QueryResultsS3AccessGrantsConfiguration](api-queryresultss3accessgrantsconfiguration.md) object

Required: No

**ResultConfiguration**

The location in Amazon S3 where query and calculation results are stored and
the encryption option, if any, used for query results. These are known as "client-side
settings". If workgroup settings override client-side settings, then the query uses the
location for the query results and the encryption configuration that are specified for
the workgroup.

Type: [ResultConfiguration](api-resultconfiguration.md) object

Required: No

**ResultReuseConfiguration**

Specifies the query result reuse behavior that was used for the query.

Type: [ResultReuseConfiguration](https://docs.aws.amazon.com/athena/latest/APIReference/API_ResultReuseConfiguration.html) object

Required: No

**StatementType**

The type of query statement that was run. `DDL` indicates DDL query
statements. `DML` indicates DML (Data Manipulation Language) query
statements, such as `CREATE TABLE AS SELECT`. `UTILITY` indicates
query statements other than DDL and DML, such as `SHOW CREATE TABLE`,
`EXPLAIN`, `DESCRIBE`, or `SHOW TABLES`.

Type: String

Valid Values: `DDL | DML | UTILITY`

Required: No

**Statistics**

Query execution statistics, such as the amount of data scanned, the amount of time
that the query took to process, and the type of statement that was run.

Type: [QueryExecutionStatistics](api-queryexecutionstatistics.md) object

Required: No

**Status**

The completion date, current state, submission time, and state change reason (if
applicable) for the query execution.

Type: [QueryExecutionStatus](https://docs.aws.amazon.com/athena/latest/APIReference/API_QueryExecutionStatus.html) object

Required: No

**SubstatementType**

The kind of query statement that was run.

Type: String

Required: No

**WorkGroup**

The name of the workgroup in which the query ran.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/QueryExecution)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/QueryExecution)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/QueryExecution)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PreparedStatementSummary

QueryExecutionContext
