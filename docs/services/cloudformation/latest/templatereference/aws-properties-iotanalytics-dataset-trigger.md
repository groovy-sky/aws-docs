This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset Trigger

The "DatasetTrigger"
that specifies when the data set is automatically updated.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Schedule" : Schedule,
  "TriggeringDataset" : TriggeringDataset
}

```

### YAML

```yaml

  Schedule:
    Schedule
  TriggeringDataset:
    TriggeringDataset

```

## Properties

`Schedule`

The "Schedule" when the trigger is initiated.

_Required_: No

_Type_: [Schedule](aws-properties-iotanalytics-dataset-schedule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggeringDataset`

Information about the data set whose content generation triggers the new data set content
generation.

_Required_: No

_Type_: [TriggeringDataset](aws-properties-iotanalytics-dataset-triggeringdataset.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TriggeringDataset

All content copied from https://docs.aws.amazon.com/.
