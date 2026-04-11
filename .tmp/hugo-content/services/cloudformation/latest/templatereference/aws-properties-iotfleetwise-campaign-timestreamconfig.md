This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign TimestreamConfig

The Amazon Timestream table where the AWS IoT FleetWise campaign sends data. Timestream stores and organizes data to optimize query processing time and to reduce storage costs. For more information, see [Data modeling](../../../timestream/latest/developerguide/data-modeling.md) in the _Amazon Timestream Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionRoleArn" : String,
  "TimestreamTableArn" : String
}

```

### YAML

```yaml

  ExecutionRoleArn: String
  TimestreamTableArn: String

```

## Properties

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the task execution role that grants AWS IoT FleetWise permission to deliver data to the Amazon Timestream table.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):iam::(\d{12})?:(role((\u002F)|(\u002F[\u0021-\u007F]+\u002F))[\w+=,.@-]+)$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TimestreamTableArn`

The Amazon Resource Name (ARN) of the Amazon Timestream table.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z0-9-]*):timestream:[a-zA-Z0-9-]+:[0-9]{12}:database\/[a-zA-Z0-9_.-]+\/table\/[a-zA-Z0-9_.-]+$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeBasedSignalFetchConfig

AWS::IoTFleetWise::DecoderManifest

All content copied from https://docs.aws.amazon.com/.
