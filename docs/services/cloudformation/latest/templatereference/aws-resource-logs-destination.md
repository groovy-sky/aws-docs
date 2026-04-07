This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Destination

The AWS::Logs::Destination resource specifies a CloudWatch Logs destination. A
destination encapsulates a physical resource (such as an Amazon Kinesis data stream) and
enables you to subscribe that resource to a stream of log events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Logs::Destination",
  "Properties" : {
      "DestinationName" : String,
      "DestinationPolicy" : String,
      "RoleArn" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Logs::Destination
Properties:
  DestinationName: String
  DestinationPolicy: String
  RoleArn: String
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`DestinationName`

The name of the destination.

_Required_: Yes

_Type_: String

_Pattern_: `^[^:*]{1,512}$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationPolicy`

An IAM policy document that governs which AWS accounts can create
subscription filters against this destination.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of an IAM role that permits CloudWatch Logs to send data to the
specified AWS resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags that have been assigned to this delivery destination.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-destination-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The Amazon Resource Name (ARN) of the physical target where the log events are
delivered (for example, a Kinesis stream).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name, such as
`TestDestination`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the CloudWatch Logs destination, such as
`arn:aws:logs:us-west-1:123456789012:destination:MyDestination`.

## Examples

### Create a Destination

In the following example, the target stream ( `TestStream`) can
receive log events from CloudWatch Logs. CloudWatch Logs can use only the
`PutSubscriptionFilter` action against the
`TestDestination` destination.

#### JSON

```json

"DestinationWithName" : {
  "Type" : "AWS::Logs::Destination",
  "Properties" : {
    "DestinationName": "TestDestination",
    "RoleArn": "arn:aws:iam::123456789012:role/LogKinesisRole",
    "TargetArn": "arn:aws:kinesis:us-east-1:123456789012:stream/TestStream",
    "DestinationPolicy": "{ \"Version\": \"2012-10-17\", \"Statement\": [{ \"Effect\": \"Allow\", \"Principal\": { \"AWS\": \"logs.amazonaws.com\"}, \"Action\": \"logs:PutSubscriptionFilter\",\"Resource\": \"arn:aws:logs:us-east-1:123456789012:destination:TestDestination\"}]}"
  }
}
```

#### YAML

```yaml

DestinationWithName:
  Type: AWS::Logs::Destination
  Properties:
    DestinationName: "TestDestination"
    RoleArn: "arn:aws:iam::123456789012:role/LogKinesisRole"
    TargetArn: "arn:aws:kinesis:us-east-1:123456789012:stream/TestStream"
    DestinationPolicy: >
      {"Version" : "2012-10-17","Statement" : [{"Effect" : "Allow", "Principal" : {"AWS" : "logs.amazonaws.com"}, "Action" : "logs:PutSubscriptionFilter", "Resource" : "arn:aws:logs:us-east-1:123456789012:destination:TestDestination"}]}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
