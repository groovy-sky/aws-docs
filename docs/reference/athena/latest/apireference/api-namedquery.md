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

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/namedquery.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/namedquery.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/namedquery.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringConfiguration

NotebookMetadata

All content copied from https://docs.aws.amazon.com/.
