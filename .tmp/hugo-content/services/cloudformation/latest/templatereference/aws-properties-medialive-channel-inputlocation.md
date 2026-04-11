This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel InputLocation

The input location.

The parent of this entity is InputLossBehavior.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PasswordParam" : String,
  "Uri" : String,
  "Username" : String
}

```

### YAML

```yaml

  PasswordParam: String
  Uri: String
  Username: String

```

## Properties

`PasswordParam`

The password parameter that holds the password for accessing the downstream system. This applies only if the
downstream system requires credentials.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Uri`

The URI should be a path to a file that is accessible to the Live system (for example, an http:// URI) depending
on the output type. For example, an RTMP destination should have a URI similar to rtmp://fmsserver/live.

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

InputChannelLevel

InputLossBehavior

All content copied from https://docs.aws.amazon.com/.
