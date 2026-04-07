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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/UnprocessedQueryExecutionId)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/UnprocessedQueryExecutionId)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/UnprocessedQueryExecutionId)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UnprocessedPreparedStatementName

WorkGroup
