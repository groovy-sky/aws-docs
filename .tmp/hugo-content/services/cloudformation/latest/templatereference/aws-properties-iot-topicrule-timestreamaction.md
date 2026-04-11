This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule TimestreamAction

Describes an action that writes records into an Amazon Timestream table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DatabaseName" : String,
  "Dimensions" : [ TimestreamDimension, ... ],
  "RoleArn" : String,
  "TableName" : String,
  "Timestamp" : TimestreamTimestamp
}

```

### YAML

```yaml

  DatabaseName: String
  Dimensions:
    - TimestreamDimension
  RoleArn: String
  TableName: String
  Timestamp:
    TimestreamTimestamp

```

## Properties

`DatabaseName`

The name of an Amazon Timestream database that has the table to write records
into.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dimensions`

Metadata attributes of the time series that are written in each measure record.

_Required_: Yes

_Type_: Array of [TimestreamDimension](aws-properties-iot-topicrule-timestreamdimension.md)

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the role that grants AWS IoT permission
to write to the Timestream database table.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TableName`

The table where the message data will be written.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Timestamp`

The value to use for the entry's timestamp. If blank, the time that the entry was
processed is used.

_Required_: No

_Type_: [TimestreamTimestamp](aws-properties-iot-topicrule-timestreamtimestamp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Timestamp

TimestreamDimension

All content copied from https://docs.aws.amazon.com/.
