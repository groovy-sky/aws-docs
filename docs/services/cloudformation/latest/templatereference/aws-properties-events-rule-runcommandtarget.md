This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Rule RunCommandTarget

Information about the EC2 instances that are to be sent the command, specified as
key-value pairs. Each `RunCommandTarget` block can include only one key, but this
key may specify multiple values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Key: String
  Values:
    - String

```

## Properties

`Key`

Can be either `tag:` _tag-key_ or
`InstanceIds`.

_Required_: Yes

_Type_: String

_Pattern_: `^[\p{L}\p{Z}\p{N}_.:/=+\-@]*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

If `Key` is `tag:` _tag-key_, `Values`
is a list of tag values. If `Key` is `InstanceIds`, `Values`
is a list of Amazon EC2 instance IDs.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RunCommandParameters

SageMakerPipelineParameter

All content copied from https://docs.aws.amazon.com/.
