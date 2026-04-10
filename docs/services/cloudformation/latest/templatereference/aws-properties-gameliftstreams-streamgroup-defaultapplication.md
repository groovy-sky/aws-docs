This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLiftStreams::StreamGroup DefaultApplication

Represents the default Amazon GameLift Streams application that a stream group hosts.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "Id" : String
}

```

### YAML

```yaml

  Arn: String
  Id: String

```

## Properties

`Arn`

An
[Amazon Resource Name (ARN)](../../../iam/latest/userguide/reference-arns.md) that uniquely identifies the application resource. Example ARN: `arn:aws:gameliftstreams:us-west-2:111122223333:application/a-9ZY8X7Wv6`.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws:gameliftstreams:([^:
]*):([0-9]{12}):([^:
]*)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Id`

An
ID that uniquely identifies the application resource. Example ID: `a-9ZY8X7Wv6`.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9-]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLiftStreams::StreamGroup

LocationConfiguration

All content copied from https://docs.aws.amazon.com/.
