This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppSync::ChannelNamespace AuthMode

Describes an authorization configuration. Use `AuthMode` to specify the
publishing and subscription authorization configuration for an Event API.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthType" : String
}

```

### YAML

```yaml

  AuthType: String

```

## Properties

`AuthType`

The authorization type.

_Required_: No

_Type_: String

_Allowed values_: `AMAZON_COGNITO_USER_POOLS | AWS_IAM | API_KEY | OPENID_CONNECT | AWS_LAMBDA`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::AppSync::ChannelNamespace

HandlerConfig

All content copied from https://docs.aws.amazon.com/.
