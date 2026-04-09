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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/honeycodeconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/honeycodeconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/honeycodeconnectorprofilecredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GoogleAnalyticsSourceProperties

HoneycodeConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
