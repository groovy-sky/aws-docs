# ConnectorOAuthRequest

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

## Contents

**authCode**

The code provided by the connector when it has been authenticated via the connected app.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**redirectUri**

The URL to which the authentication server redirects the browser after authorization has
been granted.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `\S+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/connectoroauthrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/connectoroauthrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/connectoroauthrequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConnectorMetadata

ConnectorOperator

All content copied from https://docs.aws.amazon.com/.
