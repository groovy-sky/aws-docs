This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase CloudFormationAction

Specifies the CloudFormation action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionType" : String,
  "Resource" : String
}

```

### YAML

```yaml

  ActionType: String
  Resource: String

```

## Properties

`ActionType`

The action type of the CloudFormation action.

_Required_: No

_Type_: String

_Allowed values_: `Create | Delete`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

The resource of the CloudFormation action.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Batch

CompareAction

All content copied from https://docs.aws.amazon.com/.
