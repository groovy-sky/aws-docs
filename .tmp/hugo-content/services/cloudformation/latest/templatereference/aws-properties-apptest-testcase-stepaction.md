This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppTest::TestCase StepAction

Specifies a step action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CompareAction" : CompareAction,
  "MainframeAction" : MainframeAction,
  "ResourceAction" : ResourceAction
}

```

### YAML

```yaml

  CompareAction:
    CompareAction
  MainframeAction:
    MainframeAction
  ResourceAction:
    ResourceAction

```

## Properties

`CompareAction`

The compare action of the step action.

_Required_: No

_Type_: [CompareAction](aws-properties-apptest-testcase-compareaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MainframeAction`

The mainframe action of the step action.

_Required_: No

_Type_: [MainframeAction](aws-properties-apptest-testcase-mainframeaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceAction`

The resource action of the step action.

_Required_: No

_Type_: [ResourceAction](aws-properties-apptest-testcase-resourceaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Step

TargetDatabaseMetadata

All content copied from https://docs.aws.amazon.com/.
