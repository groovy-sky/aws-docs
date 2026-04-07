This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::CloudWatchAlarmTemplate

The `AWS::MediaLive::CloudWatchAlarmTemplate` resource Property description not available. for MediaLive.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MediaLive::CloudWatchAlarmTemplate",
  "Properties" : {
      "ComparisonOperator" : String,
      "DatapointsToAlarm" : Number,
      "Description" : String,
      "EvaluationPeriods" : Number,
      "GroupIdentifier" : String,
      "MetricName" : String,
      "Name" : String,
      "Period" : Number,
      "Statistic" : String,
      "Tags" : {Key: Value, ...},
      "TargetResourceType" : String,
      "Threshold" : Number,
      "TreatMissingData" : String
    }
}

```

### YAML

```yaml

Type: AWS::MediaLive::CloudWatchAlarmTemplate
Properties:
  ComparisonOperator: String
  DatapointsToAlarm: Number
  Description: String
  EvaluationPeriods: Number
  GroupIdentifier: String
  MetricName: String
  Name: String
  Period: Number
  Statistic: String
  Tags:
    Key: Value
  TargetResourceType: String
  Threshold: Number
  TreatMissingData: String

```

## Properties

`ComparisonOperator`

The comparison operator used to compare the specified statistic and the threshold.

_Required_: Yes

_Type_: String

_Allowed values_: `GreaterThanOrEqualToThreshold | GreaterThanThreshold | LessThanThreshold | LessThanOrEqualToThreshold`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DatapointsToAlarm`

The number of datapoints within the evaluation period that must be breaching to trigger the alarm.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A resource's optional description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EvaluationPeriods`

The number of periods over which data is compared to the specified threshold.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupIdentifier`

A cloudwatch alarm template group's identifier. Can be either be its id or current name.

_Required_: No

_Type_: String

_Pattern_: `^[^\s]+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the metric associated with the alarm. Must be compatible with targetResourceType.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

A resource's name. Names must be unique within the scope of a resource type in a specific region.

_Required_: Yes

_Type_: String

_Pattern_: `^[^\s]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Period`

The period, in seconds, over which the specified statistic is applied.

_Required_: Yes

_Type_: Number

_Minimum_: `10`

_Maximum_: `86400`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Statistic`

The statistic to apply to the alarm's metric data.

_Required_: Yes

_Type_: String

_Allowed values_: `SampleCount | Average | Sum | Minimum | Maximum`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `.+`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetResourceType`

The resource type this template should dynamically generate CloudWatch metric alarms for.

_Required_: Yes

_Type_: String

_Allowed values_: `CLOUDFRONT_DISTRIBUTION | MEDIALIVE_MULTIPLEX | MEDIALIVE_CHANNEL | MEDIALIVE_INPUT_DEVICE | MEDIAPACKAGE_CHANNEL | MEDIAPACKAGE_ORIGIN_ENDPOINT | MEDIACONNECT_FLOW | MEDIATAILOR_PLAYBACK_CONFIGURATION | S3_BUCKET`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Threshold`

The threshold value to compare with the specified statistic.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatMissingData`

Specifies how missing data points are treated when evaluating the alarm's condition.

_Required_: Yes

_Type_: String

_Allowed values_: `notBreaching | breaching | ignore | missing`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

A cloudwatch alarm template's ARN (Amazon Resource Name)

`CreatedAt`

The date and time of resource creation.

`GroupId`

A CloudWatch alarm template group's id. Amazon Web Services provided template groups have ids that start with <code>\`aws-\`</code>

`Id`

A cloudwatch alarm template's id. Amazon Web Services provided templates have ids that start with <code>\`aws-\`</code>.

`Identifier`

Property description not available.

`ModifiedAt`

The date and time of latest resource modification.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tags

AWS::MediaLive::CloudWatchAlarmTemplateGroup
