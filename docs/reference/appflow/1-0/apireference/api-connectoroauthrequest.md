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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/ConnectorOAuthRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/ConnectorOAuthRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/ConnectorOAuthRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectorMetadata

ConnectorOperator
