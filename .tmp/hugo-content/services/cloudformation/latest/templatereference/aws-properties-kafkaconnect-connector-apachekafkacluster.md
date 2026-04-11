This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector ApacheKafkaCluster

The details of the Apache Kafka cluster to which the connector is connected.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BootstrapServers" : String,
  "Vpc" : Vpc
}

```

### YAML

```yaml

  BootstrapServers: String
  Vpc:
    Vpc

```

## Properties

`BootstrapServers`

The bootstrap servers of the cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Vpc`

Details of an Amazon VPC which has network connectivity to the Apache Kafka
cluster.

_Required_: Yes

_Type_: [Vpc](aws-properties-kafkaconnect-connector-vpc.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::KafkaConnect::Connector

AutoScaling

All content copied from https://docs.aws.amazon.com/.
