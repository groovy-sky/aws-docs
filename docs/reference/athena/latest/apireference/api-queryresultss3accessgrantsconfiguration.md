# QueryResultsS3AccessGrantsConfiguration

Specifies whether Amazon S3 access grants are enabled for query
results.

## Contents

**AuthenticationType**

The authentication type used for Amazon S3 access grants. Currently, only
`DIRECTORY_IDENTITY` is supported.

Type: String

Valid Values: `DIRECTORY_IDENTITY`

Required: Yes

**EnableS3AccessGrants**

Specifies whether Amazon S3 access grants are enabled for query
results.

Type: Boolean

Required: Yes

**CreateUserLevelPrefix**

When enabled, appends the user ID as an Amazon S3 path prefix to the query
result output location.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/queryresultss3accessgrantsconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/queryresultss3accessgrantsconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/queryresultss3accessgrantsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryExecutionStatus

QueryRuntimeStatistics

All content copied from https://docs.aws.amazon.com/.
