This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::RefreshSchedule

Creates a refresh schedule for a dataset in Quick.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::QuickSight::RefreshSchedule",
  "Properties" : {
      "AwsAccountId" : String,
      "DataSetId" : String,
      "Schedule" : RefreshScheduleMap
    }
}

```

### YAML

```yaml

Type: AWS::QuickSight::RefreshSchedule
Properties:
  AwsAccountId: String
  DataSetId: String
  Schedule:
    RefreshScheduleMap

```

## Properties

`AwsAccountId`

The AWS account ID of the account that you are creating a schedule in.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSetId`

The ID of the dataset that you are creating a refresh schedule for.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Schedule`

The refresh schedule of a dataset.

_Required_: No

_Type_: [RefreshScheduleMap](aws-properties-quicksight-refreshschedule-refreshschedulemap.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) for the refresh schedule.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

RefreshOnDay

All content copied from https://docs.aws.amazon.com/.
