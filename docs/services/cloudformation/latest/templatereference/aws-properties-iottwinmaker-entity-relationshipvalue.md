This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity RelationshipValue

The entity relationship.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetComponentName" : String,
  "TargetEntityId" : String
}

```

### YAML

```yaml

  TargetComponentName: String
  TargetEntityId: String

```

## Properties

`TargetComponentName`

The target component name.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z_\-0-9]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetEntityId`

The target entity Id.

_Required_: No

_Type_: String

_Pattern_: `[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|^[a-zA-Z0-9][a-zA-Z_\-0-9.:]*[a-zA-Z0-9]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Relationship

Status

All content copied from https://docs.aws.amazon.com/.
