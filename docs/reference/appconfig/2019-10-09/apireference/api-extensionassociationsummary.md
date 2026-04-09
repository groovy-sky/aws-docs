# ExtensionAssociationSummary

Information about an association between an extension and an AWS AppConfig
resource such as an application, environment, or configuration profile. Call
`GetExtensionAssociation` to get more information about an
association.

## Contents

**ExtensionArn**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

Required: No

**Id**

The extension association ID. This ID is used to call other
`ExtensionAssociation` API actions such as
`GetExtensionAssociation` or `DeleteExtensionAssociation`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**ResourceArn**

The ARNs of applications, configuration profiles, or environments defined in the
association.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/extensionassociationsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/extensionassociationsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/extensionassociationsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Environment

ExtensionSummary

All content copied from https://docs.aws.amazon.com/.
