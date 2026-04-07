# NamedQuery

A query, where `QueryString` contains the SQL statements that make up the
query.

## Contents

**Database**

The database to which the query belongs.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**Name**

The query name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**QueryString**

The SQL statements that make up the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 262144.

Required: Yes

**Description**

The query description.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**NamedQueryId**

The unique identifier of the query.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `\S+`

Required: No

**WorkGroup**

The name of the workgroup that contains the named query.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/NamedQuery)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/NamedQuery)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/NamedQuery)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MonitoringConfiguration

NotebookMetadata
