This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::EventStream

Creates a new event stream for an application or updates the settings of an
existing event stream for an application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::EventStream",
  "Properties" : {
      "ApplicationId" : String,
      "DestinationStreamArn" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::EventStream
Properties:
  ApplicationId: String
  DestinationStreamArn: String
  RoleArn: String

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application that you want to export data
from.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationStreamArn`

The Amazon Resource Name (ARN) of the Amazon Kinesis Data Stream or Amazon Data Firehose delivery stream that you want to publish event data
to.

For a Kinesis Data Stream, the ARN format is:
`arn:aws:kinesis:region:account-id:stream/stream_name`

For a Firehose delivery stream, the ARN format is:
`arn:aws:firehose:region:account-id:deliverystream/stream_name`

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The AWS Identity and Access Management (IAM) role that authorizes Amazon Pinpoint to publish
event data to the stream in your AWS account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that the event stream is associated with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::EmailTemplate

AWS::Pinpoint::GCMChannel
