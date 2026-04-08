# SalesforceConnectorProfileCredentials

The connector-specific profile credentials required when using Salesforce.

## Contents

**accessToken**

The credentials used to access protected Salesforce resources.

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

**jwtToken**

A JSON web token (JWT) that authorizes Amazon AppFlow to access your Salesforce
records.

Type: String

Length Constraints: Maximum length of 8000.

Pattern: `^([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_=]+)\.([a-zA-Z0-9_\-\+\/=]*)`

Required: No

**oAuth2GrantType**

Specifies the OAuth 2.0 grant type that Amazon AppFlow uses when it requests an
access token from Salesforce. Amazon AppFlow requires an access token each time it
attempts to access your Salesforce records.

You can specify one of the following values:

AUTHORIZATION\_CODE

Amazon AppFlow passes an authorization code when it requests the access token
from Salesforce. Amazon AppFlow receives the authorization code from Salesforce
after you log in to your Salesforce account and authorize Amazon AppFlow to access
your records.

JWT\_BEARER

Amazon AppFlow passes a JSON web token (JWT) when it requests the access token
from Salesforce. You provide the JWT to Amazon AppFlow when you define the
connection to your Salesforce account. When you use this grant type, you don't need to
log in to your Salesforce account to authorize Amazon AppFlow to access your
records.

###### Note

The CLIENT\_CREDENTIALS value is not supported for Salesforce.

Type: String

Valid Values: `CLIENT_CREDENTIALS | AUTHORIZATION_CODE | JWT_BEARER`

Required: No

**oAuthRequest**

The OAuth requirement needed to request security tokens from the connector endpoint.

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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/salesforceconnectorprofilecredentials.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/salesforceconnectorprofilecredentials.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/salesforceconnectorprofilecredentials.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3SourceProperties

SalesforceConnectorProfileProperties
