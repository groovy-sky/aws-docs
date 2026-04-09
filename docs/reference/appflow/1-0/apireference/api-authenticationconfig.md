# AuthenticationConfig

Contains information about the authentication config that the connector supports.

## Contents

**customAuthConfigs**

Contains information required for custom authentication.

Type: Array of [CustomAuthConfig](api-customauthconfig.md) objects

Required: No

**isApiKeyAuthSupported**

Indicates whether API key authentication is supported by the connector

Type: Boolean

Required: No

**isBasicAuthSupported**

Indicates whether basic authentication is supported by the connector.

Type: Boolean

Required: No

**isCustomAuthSupported**

Indicates whether custom authentication is supported by the connector

Type: Boolean

Required: No

**isOAuth2Supported**

Indicates whether OAuth 2.0 authentication is supported by the connector.

Type: Boolean

Required: No

**oAuth2Defaults**

Contains the default values required for OAuth 2.0 authentication.

Type: [OAuth2Defaults](api-oauth2defaults.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/authenticationconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/authenticationconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/authenticationconfig.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ApiKeyCredentials

AuthParameter

All content copied from https://docs.aws.amazon.com/.
