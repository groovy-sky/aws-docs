This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy IncludeResources

Specifies how the lifecycle policy should apply actions to selected resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amis" : Boolean,
  "Containers" : Boolean,
  "Snapshots" : Boolean
}

```

### YAML

```yaml

  Amis: Boolean
  Containers: Boolean
  Snapshots: Boolean

```

## Properties

`Amis`

Specifies whether the lifecycle action should apply to distributed AMIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Containers`

Specifies whether the lifecycle action should apply to distributed containers.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Snapshots`

Specifies whether the lifecycle action should apply to snapshots associated with distributed AMIs.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

LastLaunched

All content copied from https://docs.aws.amazon.com/.
