This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector WorkerConfiguration

The configuration of the workers, which are the processes that run the connector
logic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Revision" : Integer,
  "WorkerConfigurationArn" : String
}

```

### YAML

```yaml

  Revision: Integer
  WorkerConfigurationArn: String

```

## Properties

`Revision`

The revision of the worker configuration.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`WorkerConfigurationArn`

The Amazon Resource Name (ARN) of the worker configuration.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):kafkaconnect:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Vpc

WorkerLogDelivery
