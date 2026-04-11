This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition VolumeFrom

Details on a data volume from another container in the same task definition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ReadOnly" : Boolean,
  "SourceContainer" : String
}

```

### YAML

```yaml

  ReadOnly: Boolean
  SourceContainer: String

```

## Properties

`ReadOnly`

If this value is `true`, the container has read-only access to the volume.
If this value is `false`, then the container can write to the volume. The
default value is `false`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceContainer`

The name of another container within the same task definition to mount volumes
from.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Volume

AWS::ECS::TaskSet

All content copied from https://docs.aws.amazon.com/.
