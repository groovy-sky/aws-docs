This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile PardotConnectorProfileCredentials

The connector-specific profile credentials required when using Salesforce Pardot.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "ClientCredentialsArn" : String,
  "ConnectorOAuthRequest" : ConnectorOAuthRequest,
  "RefreshToken" : String
}

```

### YAML

```yaml

  AccessToken: String
  ClientCredentialsArn: String
  ConnectorOAuthRequest:
    ConnectorOAuthRequest
  RefreshToken: String

```

## Properties

`AccessToken`

The credentials used to access protected Salesforce Pardot resources.

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

Property description not available.

_Required_: No

_Type_: [ConnectorOAuthRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-connectoroauthrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RefreshToken`

The credentials used to acquire new access tokens.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OAuthProperties

PardotConnectorProfileProperties
