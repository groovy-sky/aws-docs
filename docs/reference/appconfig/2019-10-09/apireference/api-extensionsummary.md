# ExtensionSummary

Information about an extension. Call `GetExtension` to get more information
about an extension.

## Contents

**Arn**

The system-generated Amazon Resource Name (ARN) for the extension.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:(aws[a-zA-Z-]*)?:[a-z]+:((eusc-)?[a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1})?:(\d{12})?:[a-zA-Z0-9-_/:.]+`

Required: No

**Description**

Information about the extension.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**Id**

The system-generated ID of the extension.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Name**

The extension name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**VersionNumber**

The extension version number.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/extensionsummary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/extensionsummary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/extensionsummary.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExtensionAssociationSummary

HostedConfigurationVersionSummary

All content copied from https://docs.aws.amazon.com/.
