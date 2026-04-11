# DocumentAttribute

A document attribute or metadata field.

## Contents

**name**

The identifier for the attribute.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 200.

Pattern: `[a-zA-Z0-9_][a-zA-Z0-9_-]*`

Required: Yes

**value**

The value of the attribute.

Type: [DocumentAttributeValue](api-documentattributevalue.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/documentattribute.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/documentattribute.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/documentattribute.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAclUser

DocumentAttributeBoostingConfiguration

All content copied from https://docs.aws.amazon.com/.
