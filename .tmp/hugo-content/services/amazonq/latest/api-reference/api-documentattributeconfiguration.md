# DocumentAttributeConfiguration

Configuration information for document attributes. Document attributes are metadata or
fields associated with your documents. For example, the company department name
associated with each document.

For more information, see [Understanding document attributes](../business-use-dg/doc-attributes.md).

## Contents

**name**

The name of the document attribute.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 30.

Pattern: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

Required: No

**search**

Information about whether the document attribute can be used by an end user to search
for information on their web experience.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

**type**

The type of document attribute.

Type: String

Valid Values: `STRING | STRING_LIST | NUMBER | DATE`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentattributeconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentattributeconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentattributeconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttributeCondition

DocumentAttributeTarget

All content copied from https://docs.aws.amazon.com/.
