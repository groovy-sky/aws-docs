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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/qbusiness-2023-11-27/DocumentAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/qbusiness-2023-11-27/DocumentAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/qbusiness-2023-11-27/DocumentAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DocumentAclUser

DocumentAttributeBoostingConfiguration
