This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Replicator AmazonMskCluster

Details of an Amazon MSK Cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MskClusterArn" : String
}

```

### YAML

```yaml

  MskClusterArn: String

```

## Properties

`MskClusterArn`

The Amazon Resource Name (ARN) of an Amazon MSK cluster.

_Required_: Yes

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn):kafka:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MSK::Replicator

ConsumerGroupReplication

All content copied from https://docs.aws.amazon.com/.
