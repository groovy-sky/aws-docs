This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::RealtimeLogConfig

A real-time log configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudFront::RealtimeLogConfig",
  "Properties" : {
      "EndPoints" : [ EndPoint, ... ],
      "Fields" : [ String, ... ],
      "Name" : String,
      "SamplingRate" : Number
    }
}

```

### YAML

```yaml

Type: AWS::CloudFront::RealtimeLogConfig
Properties:
  EndPoints:
    - EndPoint
  Fields:
    - String
  Name: String
  SamplingRate: Number

```

## Properties

`EndPoints`

Contains information about the Amazon Kinesis data stream where you are sending real-time
log data for this real-time log configuration.

_Required_: Yes

_Type_: Array of [EndPoint](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cloudfront-realtimelogconfig-endpoint.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Fields`

A list of fields that are included in each real-time log record. In an API response,
the fields are provided in the same order in which they are sent to the Amazon Kinesis data
stream.

For more information about fields, see [Real-time log configuration fields](../../../amazoncloudfront/latest/developerguide/real-time-logs.md#understand-real-time-log-config-fields) in the
_Amazon CloudFront Developer Guide_.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The unique name of this real-time log configuration.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SamplingRate`

The sampling rate for this real-time log configuration. The sampling rate determines
the percentage of viewer requests that are represented in the real-time log data. The
sampling rate is an integer between 1 and 100, inclusive.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the real-time log
configuration. For example:
`arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the real-time log configuration. For example:
`arn:aws:cloudfront::111122223333:realtime-log-config/ExampleNameForRealtimeLogConfig`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PublicKeyConfig

EndPoint
