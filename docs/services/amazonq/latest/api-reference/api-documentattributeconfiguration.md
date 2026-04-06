# DocumentAttributeConfiguration

Configuration information for document attributes. Document attributes are metadata or
fields associated with your documents. For example, the company department name
associated with each document.

For more information, see [Understanding document attributes](https://docs.aws.amazon.com/amazonq/latest/business-use-dg/doc-attributes.html).

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/DocumentAttributeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/DocumentAttributeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/DocumentAttributeConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DocumentAttributeCondition

DocumentAttributeTarget
