# ServiceAccountCredentials

Describes the credentials for the service account used by the fleet or image builder to connect to the directory.

## Contents

**AccountName**

The user name of the account. This account must have the following privileges: create computer objects,
join computers to the domain, and change/reset the password on descendant computer objects for the
organizational units specified.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**AccountPassword**

The password for the account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 127.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/serviceaccountcredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/serviceaccountcredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/serviceaccountcredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScriptDetails

Session

All content copied from https://docs.aws.amazon.com/.
