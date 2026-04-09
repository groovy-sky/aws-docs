# Database

Contains metadata information for a database in a data catalog.

## Contents

**Name**

The name of the database.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Required: Yes

**Description**

An optional description of the database.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**Parameters**

A set of custom key/value pairs.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 255.

Key Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Value Length Constraints: Maximum length of 51200.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/database.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/database.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/database.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomerContentEncryptionConfiguration

DataCatalog

All content copied from https://docs.aws.amazon.com/.
