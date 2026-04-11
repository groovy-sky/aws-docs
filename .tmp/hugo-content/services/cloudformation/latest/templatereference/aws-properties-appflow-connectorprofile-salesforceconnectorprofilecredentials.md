This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile SalesforceConnectorProfileCredentials

The connector-specific profile credentials required when using Salesforce.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "ClientCredentialsArn" : String,
  "ConnectorOAuthRequest" : ConnectorOAuthRequest,
  "JwtToken" : String,
  "OAuth2GrantType" : String,
  "RefreshToken" : String
}

```

### YAML

```yaml

  AccessToken: String
  ClientCredentialsArn: String
  ConnectorOAuthRequest:
    ConnectorOAuthRequest
  JwtToken: String
  OAuth2GrantType: String
  RefreshToken: String

```

## Properties

`AccessToken`

The credentials used to access protected Salesforce resources.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientCredentialsArn`

The secret manager ARN, which contains the client ID and client secret of the connected
app.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:secretsmanager:.*:[0-9]+:.*`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorOAuthRequest`

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

_Required_: No

_Type_: [ConnectorOAuthRequest](aws-properties-appflow-connectorprofile-connectoroauthrequest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtToken`

A JSON web token (JWT) that authorizes Amazon AppFlow to access your Salesforce
records.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9-_=]+\.[A-Za-z0-9-_=]+\.[A-Za-z0-9-_.+/=]*$`

_Maximum_: `8000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OAuth2GrantType`

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

_Required_: No

_Type_: String

_Allowed values_: `CLIENT_CREDENTIALS | AUTHORIZATION_CODE | JWT_BEARER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

The credentials used to acquire new access tokens.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [SalesforceConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-salesforceconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RedshiftConnectorProfileProperties

SalesforceConnectorProfileProperties

All content copied from https://docs.aws.amazon.com/.
