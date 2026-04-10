This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::QuickConnect UserQuickConnectConfig

Contains information about the quick connect configuration settings for a user. The contact flow must be of type
Transfer to Agent.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContactFlowArn" : String,
  "UserArn" : String
}

```

### YAML

```yaml

  ContactFlowArn: String
  UserArn: String

```

## Properties

`ContactFlowArn`

The Amazon Resource Name (ARN) of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/contact-flow/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserArn`

The Amazon Resource Name (ARN) of the user.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent/[-a-zA-Z0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::Connect::RoutingProfile

All content copied from https://docs.aws.amazon.com/.
