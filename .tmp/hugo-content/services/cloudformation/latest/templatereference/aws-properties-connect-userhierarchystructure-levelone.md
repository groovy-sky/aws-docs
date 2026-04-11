This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::UserHierarchyStructure LevelOne

The `LevelOne` property type specifies Property description not available. for an [AWS::Connect::UserHierarchyStructure](aws-resource-connect-userhierarchystructure.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HierarchyLevelArn" : String,
  "HierarchyLevelId" : String,
  "Name" : String
}

```

### YAML

```yaml

  HierarchyLevelArn: String
  HierarchyLevelId: String
  Name: String

```

## Properties

`HierarchyLevelArn`

The Amazon Resource Name (ARN) of the hierarchy level.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*/agent-group-level/[-0-9]*$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HierarchyLevelId`

The identifier of the hierarchy level.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the hierarchy level.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LevelFour

LevelThree

All content copied from https://docs.aws.amazon.com/.
