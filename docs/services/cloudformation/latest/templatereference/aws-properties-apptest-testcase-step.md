This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase Step

Defines a step.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Action" : StepAction,
  "Description" : String,
  "Name" : String
}

```

### YAML

```yaml

  Action:
    StepAction
  Description: String
  Name: String

```

## Properties

`Action`

The action of the step.

_Required_: Yes

_Type_: [StepAction](aws-properties-apptest-testcase-stepaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the step.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the step.

_Required_: Yes

_Type_: String

_Pattern_: `^[A-Za-z][A-Za-z0-9_\-]{1,59}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceDatabaseMetadata

StepAction

All content copied from https://docs.aws.amazon.com/.
