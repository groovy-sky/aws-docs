This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Analysis FilterOperationTargetVisualsConfiguration

The configuration of target visuals that you want to be filtered.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SameSheetTargetVisualConfiguration" : SameSheetTargetVisualConfiguration
}

```

### YAML

```yaml

  SameSheetTargetVisualConfiguration:
    SameSheetTargetVisualConfiguration

```

## Properties

`SameSheetTargetVisualConfiguration`

The configuration of the same-sheet target visuals that you want to be filtered.

_Required_: No

_Type_: [SameSheetTargetVisualConfiguration](aws-properties-quicksight-analysis-samesheettargetvisualconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterOperationSelectedFieldsConfiguration

FilterRelativeDateTimeControl

All content copied from https://docs.aws.amazon.com/.
