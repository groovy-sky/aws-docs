# PreparedStatement

A prepared SQL statement for use with Athena.

## Contents

**Description**

The description of the prepared statement.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**LastModifiedTime**

The last modified time of the prepared statement.

Type: Timestamp

Required: No

**QueryStatement**

The query string for the prepared statement.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 262144.

Required: No

**StatementName**

The name of the prepared statement.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[a-zA-Z_][a-zA-Z0-9_@:]{1,256}`

Required: No

**WorkGroupName**

The name of the workgroup to which the prepared statement belongs.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/PreparedStatement)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/PreparedStatement)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/PreparedStatement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotebookSessionSummary

PreparedStatementSummary
