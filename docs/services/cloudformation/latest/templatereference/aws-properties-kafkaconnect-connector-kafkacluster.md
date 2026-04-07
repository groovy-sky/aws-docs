This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector KafkaCluster

The details of the Apache Kafka cluster to which the connector is connected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApacheKafkaCluster" : ApacheKafkaCluster
}

```

### YAML

```yaml

  ApacheKafkaCluster:
    ApacheKafkaCluster

```

## Properties

`ApacheKafkaCluster`

The Apache Kafka cluster to which the connector is connected.

_Required_: Yes

_Type_: [ApacheKafkaCluster](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-kafkaconnect-connector-apachekafkacluster.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FirehoseLogDelivery

KafkaClusterClientAuthentication
