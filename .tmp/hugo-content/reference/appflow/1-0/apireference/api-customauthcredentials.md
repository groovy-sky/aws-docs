# CustomAuthCredentials

The custom credentials required for custom authentication.

## Contents

**customAuthenticationType**

The custom authentication type that the connector uses.

Type: String

Length Constraints: Maximum length of 256.

Pattern: `\S+`

Required: Yes

**credentialsMap**

A map that holds custom authentication credentials.

Type: String to string map

Map Entries: Minimum number of 0 items. Maximum number of 50 items.

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `[\w]+`

Value Length Constraints: Maximum length of 2048.

Value Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/customauthcredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/customauthcredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/customauthcredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomAuthConfig

CustomConnectorDestinationProperties

All content copied from https://docs.aws.amazon.com/.
