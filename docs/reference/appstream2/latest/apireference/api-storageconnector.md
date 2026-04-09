# StorageConnector

Describes a connector that enables persistent storage for users.

## Contents

**ConnectorType**

The type of storage connector.

Type: String

Valid Values: `HOMEFOLDERS | GOOGLE_DRIVE | ONE_DRIVE`

Required: Yes

**Domains**

The names of the domains for the account.

Type: Array of strings

Array Members: Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**DomainsRequireAdminConsent**

The OneDrive for Business domains where you require admin consent when users try to link their OneDrive account to WorkSpaces Applications. The attribute can only be specified when ConnectorType=ONE\_DRIVE.

Type: Array of strings

Array Members: Maximum number of 50 items.

Length Constraints: Minimum length of 1. Maximum length of 64.

Required: No

**ResourceIdentifier**

The ARN of the storage connector.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/storageconnector.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/storageconnector.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/storageconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StackError

StreamingExperienceSettings

All content copied from https://docs.aws.amazon.com/.
