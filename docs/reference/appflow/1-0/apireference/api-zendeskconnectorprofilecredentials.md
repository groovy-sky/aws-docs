# ZendeskConnectorProfileCredentials

The connector-specific profile credentials required when using Zendesk.

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

The credentials used to access protected Zendesk resources.

Type: String

Length Constraints: Maximum length of 4096.

Pattern: `\S+`

Required: No

**oAuthRequest**

The OAuth requirement needed to request security tokens from the connector endpoint.

Type: [ConnectorOAuthRequest](api-connectoroauthrequest.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/zendeskconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/zendeskconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/zendeskconnectorprofilecredentials.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

VeevaSourceProperties

ZendeskConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
