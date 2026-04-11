This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow StorageFlowNodeS3Configuration

Contains configurations for the Amazon S3 location in which to store the input into the node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String
}

```

### YAML

```yaml

  BucketName: String

```

## Properties

`BucketName`

The name of the Amazon S3 bucket in which to store the input into the node.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-z0-9][\.\-a-z0-9]{1,61}[a-z0-9]$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageFlowNodeConfiguration

StorageFlowNodeServiceConfiguration

All content copied from https://docs.aws.amazon.com/.
