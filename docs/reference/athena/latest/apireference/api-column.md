# Column

Contains metadata for a column in a table.

## Contents

**Name**

The name of the column.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Comment**

Optional information about the column.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: No

**Type**

The data type of the column.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 4096.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/column.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/column.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/column.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CloudWatchLoggingConfiguration

ColumnInfo

All content copied from https://docs.aws.amazon.com/.
