This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector LogDelivery

Details about log delivery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "WorkerLogDelivery" : WorkerLogDelivery
}

```

### YAML

```yaml

  WorkerLogDelivery:
    WorkerLogDelivery

```

## Properties

`WorkerLogDelivery`

The workers can send worker logs to different destination types. This configuration
specifies the details of these destinations.

_Required_: Yes

_Type_: [WorkerLogDelivery](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kafkaconnect-connector-workerlogdelivery.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KafkaClusterEncryptionInTransit

Plugin
