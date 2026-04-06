# OAuth2Credentials

The OAuth 2.0 credentials required for OAuth 2.0 authentication.

## Contents

**accessToken**

The access token used to access the connector on your behalf.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**clientId**

The identifier for the desired client.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**clientSecret**

The client secret used by the OAuth client to authenticate to the authorization
server.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

**oAuthRequest**

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

Type: [ConnectorOAuthRequest](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_ConnectorOAuthRequest.html) object

Required: No

**refreshToken**

The refresh token used to refresh an expired access token.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/OAuth2Credentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/OAuth2Credentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/OAuth2Credentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetadataCatalogDetail

OAuth2CustomParameter
