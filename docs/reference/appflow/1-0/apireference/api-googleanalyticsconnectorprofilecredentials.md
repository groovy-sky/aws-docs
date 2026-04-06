# GoogleAnalyticsConnectorProfileCredentials

The connector-specific profile credentials required by Google Analytics.

## Contents

**clientId**

The identifier for the desired client.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**clientSecret**

The client secret used by the OAuth client to authenticate to the authorization server.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: Yes

**accessToken**

The credentials used to access protected Google Analytics resources.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**oAuthRequest**

The OAuth requirement needed to request security tokens from the connector endpoint.

Type: [ConnectorOAuthRequest](api-connectoroauthrequest.md) object

Required: No

**refreshToken**

The credentials used to acquire new access tokens. This is required only for OAuth2
access tokens, and is not required for OAuth1 access tokens.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/GoogleAnalyticsConnectorProfileCredentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/GoogleAnalyticsConnectorProfileCredentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/GoogleAnalyticsConnectorProfileCredentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GlueDataCatalogConfig

GoogleAnalyticsConnectorProfileProperties
