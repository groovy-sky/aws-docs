# PardotConnectorProfileCredentials

The connector-specific profile credentials required when using Salesforce Pardot.

## Contents

**accessToken**

The credentials used to access protected Salesforce Pardot resources.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**clientCredentialsArn**

The secret manager ARN, which contains the client ID and client secret of the connected
app.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `arn:aws:secretsmanager:.*:[0-9]+:.*`

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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/pardotconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/pardotconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/pardotconnectorprofilecredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OAuthProperties

PardotConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
