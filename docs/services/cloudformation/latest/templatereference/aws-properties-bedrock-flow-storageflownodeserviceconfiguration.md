This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow StorageFlowNodeServiceConfiguration

Contains configurations for the service to use for storing the input into the node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : StorageFlowNodeS3Configuration
}

```

### YAML

```yaml

  S3:
    StorageFlowNodeS3Configuration

```

## Properties

`S3`

Contains configurations for the Amazon S3 location in which to store the input into the node.

_Required_: No

_Type_: [StorageFlowNodeS3Configuration](aws-properties-bedrock-flow-storageflownodes3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StorageFlowNodeS3Configuration

TextPromptTemplateConfiguration

All content copied from https://docs.aws.amazon.com/.
