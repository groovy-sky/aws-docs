# UnprocessedPreparedStatementName

The name of a prepared statement that could not be returned.

## Contents

**ErrorCode**

The error code returned when the request for the prepared statement failed.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**ErrorMessage**

The error message containing the reason why the prepared statement could not be
returned. The following error messages are possible:

- `INVALID_INPUT` \- The name of the prepared statement that was
provided is not valid (for example, the name is too long).

- `STATEMENT_NOT_FOUND` \- A prepared statement with the name provided
could not be found.

- `UNAUTHORIZED` \- The requester does not have permission to access
the workgroup that contains the prepared statement.

Type: String

Required: No

**StatementName**

The name of a prepared statement that could not be returned due to an error.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[a-zA-Z_][a-zA-Z0-9_@:]{1,256}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/unprocessedpreparedstatementname.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/unprocessedpreparedstatementname.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/unprocessedpreparedstatementname.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnprocessedNamedQueryId

UnprocessedQueryExecutionId
