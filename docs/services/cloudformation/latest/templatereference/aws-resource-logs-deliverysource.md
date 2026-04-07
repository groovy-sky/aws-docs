This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::DeliverySource

Creates or updates one _delivery source_ in your account. A
delivery source is an AWS resource that sends logs to an AWS destination. The destination can be CloudWatch Logs, Amazon S3, or Firehose.

Only some AWS services support being configured as a delivery source.
These services are listed as **Supported \[V2 Permissions\]**
in the table at [Enabling\
logging from AWS services.](../../../amazoncloudwatch/latest/logs/aws-logs-and-resource-policy.md)

To configure logs delivery between a supported AWS service and a
destination, you must do the following:

- Create a delivery source, which is a logical object that represents the
resource that is actually sending the logs.

- Create a _delivery destination_, which is a logical object
that represents the actual delivery destination. For more information, see
[AWS::Logs::DeliveryDestination](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-deliverydestination.html) or [PutDeliveryDestination](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-putdeliverydestination.md).

- Create a _delivery_ by pairing exactly one delivery source
and one delivery destination. For more information, see [AWS::Logs::Delivery](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-delivery.html) or [CreateDelivery](../../../../reference/amazoncloudwatchlogs/latest/apireference/api-createdelivery.md).

You can configure a single delivery source to send logs to multiple destinations by
creating multiple deliveries. You can also create multiple deliveries to configure
multiple delivery sources to send logs to the same delivery destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::DeliverySource",
  "Properties" : {
      "LogType" : String,
      "Name" : String,
      "ResourceArn" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Logs::DeliverySource
Properties:
  LogType: String
  Name: String
  ResourceArn: String
  Tags:
    - Tag

```

## Properties

`LogType`

The type of log that the source is sending. For valid values for this parameter, see the
documentation for the source service.

_Required_: No

_Type_: String

_Pattern_: `[\w-]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of the delivery source.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]*$`

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceArn`

The ARN of the AWS resource that is generating and sending logs. For
example,
`arn:aws:workmail:us-east-1:123456789012:organization/m-1234EXAMPLEabcd1234abcd1234abcd1234`

_Required_: No

_Type_: String

_Pattern_: `[\w#+=/:,.@-]*\*?`

_Minimum_: `16`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to the delivery source.

For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-deliverysource-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) that uniquely identifies this delivery source.

`ResourceArns`

This array contains the ARN of the AWS resource that sends logs and is
represented by this delivery source. Currently, only one ARN can be in the array.

`Service`

The AWS service that is sending logs.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
