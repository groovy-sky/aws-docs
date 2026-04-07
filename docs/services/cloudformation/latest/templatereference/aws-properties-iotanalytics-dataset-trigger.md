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

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TriggeringDataset`

Information about the data set whose content generation triggers the new data set content
generation.

_Required_: No

_Type_: [TriggeringDataset](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotanalytics-dataset-triggeringdataset.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

TriggeringDataset
