This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KafkaConnect::Connector

Creates a connector using the specified properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::KafkaConnect::Connector",
  "Properties" : {
      "Capacity" : Capacity,
      "ConnectorConfiguration" : {Key: Value, ...},
      "ConnectorDescription" : String,
      "ConnectorName" : String,
      "KafkaCluster" : KafkaCluster,
      "KafkaClusterClientAuthentication" : KafkaClusterClientAuthentication,
      "KafkaClusterEncryptionInTransit" : KafkaClusterEncryptionInTransit,
      "KafkaConnectVersion" : String,
      "LogDelivery" : LogDelivery,
      "NetworkType" : String,
      "Plugins" : [ Plugin, ... ],
      "ServiceExecutionRoleArn" : String,
      "Tags" : [ Tag, ... ],
      "WorkerConfiguration" : WorkerConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::KafkaConnect::Connector
Properties:
  Capacity:
    Capacity
  ConnectorConfiguration:
    Key: Value
  ConnectorDescription: String
  ConnectorName: String
  KafkaCluster:
    KafkaCluster
  KafkaClusterClientAuthentication:
    KafkaClusterClientAuthentication
  KafkaClusterEncryptionInTransit:
    KafkaClusterEncryptionInTransit
  KafkaConnectVersion: String
  LogDelivery:
    LogDelivery
  NetworkType: String
  Plugins:
    - Plugin
  ServiceExecutionRoleArn: String
  Tags:
    - Tag
  WorkerConfiguration:
    WorkerConfiguration

```

## Properties

`Capacity`

The connector's compute capacity settings.

_Required_: Yes

_Type_: [Capacity](aws-properties-kafkaconnect-connector-capacity.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorConfiguration`

The configuration of the connector.

_Required_: Yes

_Type_: Object of String

_Pattern_: `.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConnectorDescription`

The description of the connector.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConnectorName`

The name of the connector.

The connector name must be unique and can include up to 128 characters. Valid characters you can include in a connector name are: a-z, A-Z, 0-9, and -.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KafkaCluster`

The details of the Apache Kafka cluster to which the connector is connected.

_Required_: Yes

_Type_: [KafkaCluster](aws-properties-kafkaconnect-connector-kafkacluster.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KafkaClusterClientAuthentication`

The type of client authentication used to connect to the Apache Kafka cluster. The value
is NONE when no client authentication is used.

_Required_: Yes

_Type_: [KafkaClusterClientAuthentication](aws-properties-kafkaconnect-connector-kafkaclusterclientauthentication.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KafkaClusterEncryptionInTransit`

Details of encryption in transit to the Apache Kafka cluster.

_Required_: Yes

_Type_: [KafkaClusterEncryptionInTransit](aws-properties-kafkaconnect-connector-kafkaclusterencryptionintransit.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KafkaConnectVersion`

The version of Kafka Connect. It has to be compatible with both the Apache Kafka
cluster's version and the plugins.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogDelivery`

The settings for delivering connector logs to Amazon CloudWatch Logs.

_Required_: No

_Type_: [LogDelivery](aws-properties-kafkaconnect-connector-logdelivery.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NetworkType`

The network type of the connector. It gives connectors connectivity to either IPv4 (IPV4) or IPv4 and IPv6 (DUAL) destinations. Defaults to IPV4.

_Required_: No

_Type_: String

_Allowed values_: `IPV4 | DUAL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Plugins`

Specifies which plugin to use for the connector. You must specify a single-element list. Amazon MSK Connect does not currently support specifying multiple plugins.

_Required_: Yes

_Type_: Array of [Plugin](aws-properties-kafkaconnect-connector-plugin.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceExecutionRoleArn`

The Amazon Resource Name (ARN) of the IAM role used by the connector to access Amazon
Web Services resources.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):iam:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-kafkaconnect-connector-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerConfiguration`

The worker configurations that are in use with the connector.

_Required_: No

_Type_: [WorkerConfiguration](aws-properties-kafkaconnect-connector-workerconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`ConnectorArn`

The Amazon Resource Name (ARN) of the newly created connector.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Managed Streaming for Apache Kafka Connect

ApacheKafkaCluster

All content copied from https://docs.aws.amazon.com/.
