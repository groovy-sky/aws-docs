# OAuth2Defaults

Contains the default values required for OAuth 2.0 authentication.

## Contents

**authCodeUrls**

Auth code URLs that can be used for OAuth 2.0 authentication.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: No

**oauth2CustomProperties**

List of custom parameters required for OAuth 2.0 authentication.

Type: Array of [OAuth2CustomParameter](api-oauth2customparameter.md) objects

Required: No

**oauth2GrantTypesSupported**

OAuth 2.0 grant types supported by the connector.

Type: Array of strings

Valid Values: `CLIENT_CREDENTIALS | AUTHORIZATION_CODE | JWT_BEARER`

Required: No

**oauthScopes**

OAuth 2.0 scopes that the connector supports.

Type: Array of strings

Length Constraints: Maximum length of 128.

Pattern: `\S+`

Required: No

**tokenUrls**

Token URLs that can be used for OAuth 2.0 authentication.

Type: Array of strings

Length Constraints: Maximum length of 256.

Pattern: `^(https?)://[-a-zA-Z0-9+&@#/%?=~_|!:,.;]*[-a-zA-Z0-9+&@#/%=~_|]`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/oauth2defaults.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/oauth2defaults.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/oauth2defaults.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuth2CustomParameter

OAuth2Properties

All content copied from https://docs.aws.amazon.com/.
