This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase MainframeAction

Specifies the mainframe action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionType" : MainframeActionType,
  "Properties" : MainframeActionProperties,
  "Resource" : String
}

```

### YAML

```yaml

  ActionType:
    MainframeActionType
  Properties:
    MainframeActionProperties
  Resource: String

```

## Properties

`ActionType`

The action type of the mainframe action.

_Required_: Yes

_Type_: [MainframeActionType](aws-properties-apptest-testcase-mainframeactiontype.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

The properties of the mainframe action.

_Required_: No

_Type_: [MainframeActionProperties](aws-properties-apptest-testcase-mainframeactionproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resource`

The resource of the mainframe action.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

M2NonManagedApplicationAction

MainframeActionProperties

All content copied from https://docs.aws.amazon.com/.
