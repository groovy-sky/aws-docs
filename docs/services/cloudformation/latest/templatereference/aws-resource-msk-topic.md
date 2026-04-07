This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Topic

Creates a new MSK topic.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MSK::Topic",
  "Properties" : {
      "ClusterArn" : String,
      "Configs" : String,
      "PartitionCount" : Integer,
      "ReplicationFactor" : Integer,
      "TopicName" : String
    }
}

```

### YAML

```yaml

Type: AWS::MSK::Topic
Properties:
  ClusterArn: String
  Configs: String
  PartitionCount: Integer
  ReplicationFactor: Integer
  TopicName: String

```

## Properties

`ClusterArn`

The Amazon Resource Name (ARN) that uniquely identifies the cluster.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Configs`

Topic configurations encoded as a Base64 string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartitionCount`

The number of partitions for the topic.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationFactor`

The replication factor for the topic.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TopicName`

The name of the topic to create.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic function, `Ref` returns the topic ARN.

### Fn::GetAtt

`Fn::GetAtt` returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

`TopicArn`

The Amazon Resource Name (ARN) of the topic.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VpcConfig

AWS::MSK::VpcConnection
