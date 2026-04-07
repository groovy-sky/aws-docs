This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Cluster

Creates a new MSK cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::Cluster",
  "Properties" : {
      "BrokerNodeGroupInfo" : BrokerNodeGroupInfo,
      "ClientAuthentication" : ClientAuthentication,
      "ClusterName" : String,
      "ConfigurationInfo" : ConfigurationInfo,
      "EncryptionInfo" : EncryptionInfo,
      "EnhancedMonitoring" : String,
      "KafkaVersion" : String,
      "LoggingInfo" : LoggingInfo,
      "NumberOfBrokerNodes" : Integer,
      "OpenMonitoring" : OpenMonitoring,
      "Rebalancing" : Rebalancing,
      "StorageMode" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::MSK::Cluster
Properties:
  BrokerNodeGroupInfo:
    BrokerNodeGroupInfo
  ClientAuthentication:
    ClientAuthentication
  ClusterName: String
  ConfigurationInfo:
    ConfigurationInfo
  EncryptionInfo:
    EncryptionInfo
  EnhancedMonitoring: String
  KafkaVersion: String
  LoggingInfo:
    LoggingInfo
  NumberOfBrokerNodes: Integer
  OpenMonitoring:
    OpenMonitoring
  Rebalancing:
    Rebalancing
  StorageMode: String
  Tags:
    Key: Value

```

## Properties

`BrokerNodeGroupInfo`

Information about the broker nodes in the cluster.

_Required_: Yes

_Type_: [BrokerNodeGroupInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-brokernodegroupinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientAuthentication`

Includes all client authentication related information.

_Required_: No

_Type_: [ClientAuthentication](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-clientauthentication.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClusterName`

The name of the cluster.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationInfo`

Represents the configuration that you want MSK to use for the cluster.

_Required_: No

_Type_: [ConfigurationInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-configurationinfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EncryptionInfo`

Includes all encryption-related information.

_Required_: No

_Type_: [EncryptionInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-encryptioninfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnhancedMonitoring`

Specifies the level of monitoring for the MSK cluster.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | PER_BROKER | PER_TOPIC_PER_BROKER | PER_TOPIC_PER_PARTITION`

_Minimum_: `7`

_Maximum_: `23`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KafkaVersion`

The version of Apache Kafka. You can use Amazon MSK to create clusters that use [supported Apache Kafka versions](https://docs.aws.amazon.com/msk/latest/developerguide/supported-kafka-versions.html).

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LoggingInfo`

Logging info details for the cluster.

_Required_: No

_Type_: [LoggingInfo](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-logginginfo.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfBrokerNodes`

The number of broker nodes in the cluster.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OpenMonitoring`

The settings for open monitoring.

_Required_: No

_Type_: [OpenMonitoring](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-openmonitoring.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Rebalancing`

Property description not available.

_Required_: No

_Type_: [Rebalancing](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-msk-cluster-rebalancing.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageMode`

This controls storage mode for supported storage tiers.

_Required_: No

_Type_: String

_Allowed values_: `LOCAL | TIERED`

_Minimum_: `5`

_Maximum_: `6`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An arbitrary set of tags (key-value pairs) for the cluster.

_Required_: No

_Type_: Object of String

_Pattern_: `^([\p{L}\p{Z}\p{N}_.:/=+\-@]*)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, `Ref` returns the ARN of the created MSK cluster. For example, `arn:aws:kafka:us-east-1:123456789012:cluster/myCluster/abcd1234-abcd-dcba-4321-a1b2abcd9f9f-2`.

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`Arn`

The Amazon Resource Name (ARN) of the MSK cluster.

`CurrentVersion`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MSK::BatchScramSecret

BrokerLogs
