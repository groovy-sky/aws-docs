This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTTwinMaker::Entity Relationship

The entity relationship.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RelationshipType" : String,
  "TargetComponentTypeId" : String
}

```

### YAML

```yaml

  RelationshipType: String
  TargetComponentTypeId: String

```

## Properties

`RelationshipType`

The relationship type.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetComponentTypeId`

the component type Id target.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z_\.\-0-9:]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PropertyGroup

RelationshipValue

All content copied from https://docs.aws.amazon.com/.
