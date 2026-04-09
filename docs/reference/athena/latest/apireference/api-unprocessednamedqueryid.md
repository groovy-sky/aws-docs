# UnprocessedNamedQueryId

Information about a named query ID that could not be processed.

## Contents

**ErrorCode**

The error code returned when the processing request for the named query failed, if
applicable.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: No

**ErrorMessage**

The error message returned when the processing request for the named query failed, if
applicable.

Type: String

Required: No

**NamedQueryId**

The unique identifier of the named query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/unprocessednamedqueryid.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/unprocessednamedqueryid.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/unprocessednamedqueryid.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

UnprocessedPreparedStatementName

All content copied from https://docs.aws.amazon.com/.
