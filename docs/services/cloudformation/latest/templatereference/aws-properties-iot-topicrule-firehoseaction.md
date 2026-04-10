This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule FirehoseAction

Describes an action that writes data to an Amazon Kinesis Firehose stream.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchMode" : Boolean,
  "DeliveryStreamName" : String,
  "RoleArn" : String,
  "Separator" : String
}

```

### YAML

```yaml

  BatchMode: Boolean
  DeliveryStreamName: String
  RoleArn: String
  Separator: String

```

## Properties

`BatchMode`

Whether to deliver the Kinesis Data Firehose stream as a batch by using [`PutRecordBatch`](../../../../reference/firehose/latest/apireference/api-putrecordbatch.md). The default value is `false`.

When `batchMode` is `true` and the rule's SQL statement
evaluates to an Array, each Array element forms one record in the [`PutRecordBatch`](../../../../reference/firehose/latest/apireference/api-putrecordbatch.md) request. The resulting array can't have more than 500 records.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DeliveryStreamName`

The delivery stream name.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The IAM role that grants access to the Amazon Kinesis Firehose stream.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Separator`

A character separator that will be used to separate records written to the Firehose
stream. Valid values are: '\\n' (newline), '\\t' (tab), '\\r\\n' (Windows newline), ','
(comma).

_Required_: No

_Type_: String

_Pattern_: `([\n\t])|(\r\n)|(,)`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ElasticsearchAction

HttpAction

All content copied from https://docs.aws.amazon.com/.
