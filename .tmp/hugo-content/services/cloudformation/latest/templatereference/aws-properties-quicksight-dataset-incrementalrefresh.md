This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSet IncrementalRefresh

The incremental refresh configuration for a dataset.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LookbackWindow" : LookbackWindow
}

```

### YAML

```yaml

  LookbackWindow:
    LookbackWindow

```

## Properties

`LookbackWindow`

The lookback window setup for an incremental refresh configuration.

_Required_: Yes

_Type_: [LookbackWindow](aws-properties-quicksight-dataset-lookbackwindow.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportTableOperationSource

IngestionWaitPolicy

All content copied from https://docs.aws.amazon.com/.
