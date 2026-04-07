This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ZendeskConnectorProfileCredentials

The connector-specific profile credentials required when using Zendesk.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessToken" : String,
  "ClientId" : String,
  "ClientSecret" : String,
  "ConnectorOAuthRequest" : ConnectorOAuthRequest
}

```

### YAML

```yaml

  AccessToken: String
  ClientId: String
  ClientSecret: String
  ConnectorOAuthRequest:
    ConnectorOAuthRequest

```

## Properties

`AccessToken`

The credentials used to access protected Zendesk resources.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The identifier for the desired client.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientSecret`

The client secret used by the OAuth client to authenticate to the authorization server.

_Required_: Yes

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorOAuthRequest`

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

_Required_: No

_Type_: [ConnectorOAuthRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connectorprofile-connectoroauthrequest.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ZendeskConnectorProfileCredentials](../../../../reference/appflow/1-0/apireference/api-zendeskconnectorprofilecredentials.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VeevaConnectorProfileProperties

ZendeskConnectorProfileProperties
