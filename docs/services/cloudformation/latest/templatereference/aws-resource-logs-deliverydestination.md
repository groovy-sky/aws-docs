This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::DeliveryDestination

This structure contains information about one _delivery destination_ in
your account. A delivery destination is an AWS resource that represents an
AWS service that logs can be sent to. CloudWatch Logs, Amazon S3,
Firehose, and X-Ray are supported as delivery destinations.

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the resource that
is actually sending the logs. For more information, see [PutDeliverySource](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverysource.md).

- Create a _delivery destination_, which is a logical object that
represents the actual delivery destination.

- If you are delivering logs cross-account, you must use [PutDeliveryDestinationPolicy](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverydestinationpolicy.md) in the destination account to assign an IAM policy to the destination. This policy allows delivery to that destination.

- Create a _delivery_ by pairing exactly one delivery source and one
delivery destination. For more information, see [CreateDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createdelivery.md).

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure multiple
delivery sources to send logs to the same delivery destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::DeliveryDestination",
  "Properties" : {
      "DeliveryDestinationPolicy" : DestinationPolicy,
      "DeliveryDestinationType" : String,
      "DestinationResourceArn" : String,
      "Name" : String,
      "OutputFormat" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::DeliveryDestination
Properties:
  DeliveryDestinationPolicy:
    DestinationPolicy
  DeliveryDestinationType: String
  DestinationResourceArn: String
  Name: String
  OutputFormat: String
  Tags:
    - Tag

```

## Properties

`DeliveryDestinationPolicy`

An IAM policy that grants permissions to CloudWatch Logs to deliver
logs cross-account to a specified destination in this account. For examples of this
policy, see [Examples](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverydestinationpolicy.md#API_PutDeliveryDestinationPolicy_Examples) in the CloudWatch Logs API Reference.

_Required_: No

_Type_: [DestinationPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-deliverydestination-destinationpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryDestinationType`

Displays whether this delivery destination is CloudWatch Logs, Amazon S3,
Firehose, or X-Ray.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z]+$`

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationResourceArn`

The ARN of the AWS destination that this delivery destination
represents. That AWS destination can be a log group in CloudWatch Logs, an Amazon S3 bucket, or a Firehose stream.

_Required_: No

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `16`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of this delivery destination.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]*$`

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OutputFormat`

The format of the logs that are sent to this delivery destination.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z]+$`

_Minimum_: `1`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to the delivery destination.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-deliverydestination-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) that uniquely identifies this delivery
destination.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DestinationPolicy
