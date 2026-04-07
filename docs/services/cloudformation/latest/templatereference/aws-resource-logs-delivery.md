This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Delivery

This structure contains information about one _delivery_ in your
account.

A delivery is a connection between a logical _delivery source_ and a
logical _delivery destination_.

For more information, see [CreateDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createdelivery.md).

To update an existing delivery configuration, use [UpdateDeliveryConfiguration](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_UpdateDeliveryConfiguration.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::Delivery",
  "Properties" : {
      "DeliveryDestinationArn" : String,
      "DeliverySourceName" : String,
      "FieldDelimiter" : String,
      "RecordFields" : [ String, ... ],
      "S3EnableHiveCompatiblePath" : Boolean,
      "S3SuffixPath" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::Delivery
Properties:
  DeliveryDestinationArn: String
  DeliverySourceName: String
  FieldDelimiter: String
  RecordFields:
    - String
  S3EnableHiveCompatiblePath: Boolean
  S3SuffixPath: String
  Tags:
    - Tag

```

## Properties

`DeliveryDestinationArn`

The ARN of the delivery destination that is associated with this delivery.

_Required_: Yes

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `16`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeliverySourceName`

The name of the delivery source that is associated with this delivery.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]*$`

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FieldDelimiter`

The field delimiter that is used between record fields when the final output format of
a delivery is in `Plain`, `W3C`, or `Raw`
format.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RecordFields`

The list of record fields to be delivered to the destination, in order. If the
delivery's log source has mandatory fields, they must be included in this list.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3EnableHiveCompatiblePath`

Use this parameter to cause the S3 objects that contain delivered logs to use a prefix
structure that allows for integration with Apache Hive.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3SuffixPath`

Use this to reconfigure the S3 object prefix to contain either static or variable
sections. The valid variables to use in the suffix path will vary by each log source. To
find the values supported for the suffix path for each log source, use the [DescribeConfigurationTemplates](https://docs.aws.amazon.com/AmazonCloudWatchLogs/latest/APIReference/API_DescribeConfigurationTemplates.html) operation and check the
`allowedSuffixPathFields` field in the response.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the delivery.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-delivery-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) that uniquely identifies this delivery.

`DeliveryDestinationType`

Displays whether the delivery destination associated with this delivery is CloudWatch Logs, Amazon S3, or Firehose.

`DeliveryId`

The unique ID that identifies this delivery in your account.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Logs::AccountPolicy

Tag
