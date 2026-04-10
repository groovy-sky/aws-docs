This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::ConnectorProfile ConnectorOAuthRequest

Used by select connectors for which the OAuth workflow is supported, such as Salesforce,
Google Analytics, Marketo, Zendesk, and Slack.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthCode" : String,
  "RedirectUri" : String
}

```

### YAML

```yaml

  AuthCode: String
  RedirectUri: String

```

## Properties

`AuthCode`

The code provided by the connector when it has been authenticated via the connected app.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedirectUri`

The URL to which the authentication server redirects the browser after authorization has
been granted.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ConnectorOAuthRequest](../../../../reference/appflow/1-0/apireference/api-connectoroauthrequest.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthCredentials

ConnectorProfileConfig

All content copied from https://docs.aws.amazon.com/.
