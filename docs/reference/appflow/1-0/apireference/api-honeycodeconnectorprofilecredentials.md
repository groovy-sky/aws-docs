# HoneycodeConnectorProfileCredentials

The connector-specific credentials required when using Amazon Honeycode.

## Contents

**accessToken**

The credentials used to access protected Amazon Honeycode resources.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**oAuthRequest**

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

Type: [ConnectorOAuthRequest](api-connectoroauthrequest.md) object

Required: No

**refreshToken**

The credentials used to acquire new access tokens.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/HoneycodeConnectorProfileCredentials)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/HoneycodeConnectorProfileCredentials)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/HoneycodeConnectorProfileCredentials)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GoogleAnalyticsSourceProperties

HoneycodeConnectorProfileProperties
