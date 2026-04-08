# UnprocessedQueryExecutionId

Describes a query execution that failed to process.

## Contents

**ErrorCode**

The error code returned when the query execution failed to process, if
applicable.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**ErrorMessage**

The error message returned when the query execution failed to process, if
applicable.

Type: String

Required: No

**QueryExecutionId**

The unique identifier of the query execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/unprocessedqueryexecutionid.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/unprocessedqueryexecutionid.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/unprocessedqueryexecutionid.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnprocessedPreparedStatementName

WorkGroup
