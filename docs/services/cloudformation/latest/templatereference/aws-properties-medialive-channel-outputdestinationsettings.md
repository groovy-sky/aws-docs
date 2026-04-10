This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel OutputDestinationSettings

The configuration information for this output.

The parent of this entity is OutputDestination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordParam" : String,
  "StreamName" : String,
  "Url" : String,
  "Username" : String
}

```

### YAML

```yaml

  PasswordParam: String
  StreamName: String
  Url: String
  Username: String

```

## Properties

`PasswordParam`

The password parameter that holds the password for accessing the downstream system. This password parameter
applies only if the downstream system requires credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StreamName`

The stream name for the content. This applies only to RTMP outputs.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Url`

The URL for the destination.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Username`

The user name to connect to the downstream system. This applies only if the downstream system requires
credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OutputDestination

OutputGroup

All content copied from https://docs.aws.amazon.com/.
