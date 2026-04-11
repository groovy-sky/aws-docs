# AppliedExtension

An extension that was invoked during a deployment.

## Contents

**ExtensionAssociationId**

The system-generated ID for the association.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**ExtensionId**

The system-generated ID of the extension.

Type: String

Pattern: `[a-z0-9]{4,7}`

Required: No

**Parameters**

One or more parameters for the actions called by the extension.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 10 items.

Key Pattern: `^[^\/#:\n]{1,64}$`

Value Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**VersionNumber**

The extension version number.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appconfig-2019-10-09/appliedextension.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appconfig-2019-10-09/appliedextension.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appconfig-2019-10-09/appliedextension.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Application

BadRequestDetails

All content copied from https://docs.aws.amazon.com/.
