This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pipes::Pipe PipeSourceActiveMQBrokerParameters

The parameters for using an Active MQ broker as a source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BatchSize" : Integer,
  "Credentials" : MQBrokerAccessCredentials,
  "MaximumBatchingWindowInSeconds" : Integer,
  "QueueName" : String
}

```

### YAML

```yaml

  BatchSize: Integer
  Credentials:
    MQBrokerAccessCredentials
  MaximumBatchingWindowInSeconds: Integer
  QueueName: String

```

## Properties

`BatchSize`

The maximum number of records to include in each batch.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Credentials`

The credentials needed to access the resource.

_Required_: Yes

_Type_: [MQBrokerAccessCredentials](aws-properties-pipes-pipe-mqbrokeraccesscredentials.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumBatchingWindowInSeconds`

The maximum length of a time to wait for events.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `300`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueueName`

The name of the destination queue to consume.

_Required_: Yes

_Type_: String

_Pattern_: `^[\s\S]*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PipeLogConfiguration

PipeSourceDynamoDBStreamParameters

All content copied from https://docs.aws.amazon.com/.
