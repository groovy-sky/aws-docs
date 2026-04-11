This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow RetrievalFlowNodeServiceConfiguration

Contains configurations for the service to use for retrieving data to return as the output from the node.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : RetrievalFlowNodeS3Configuration
}

```

### YAML

```yaml

  S3:
    RetrievalFlowNodeS3Configuration

```

## Properties

`S3`

Contains configurations for the Amazon S3 location from which to retrieve data to return as the output from the node.

_Required_: No

_Type_: [RetrievalFlowNodeS3Configuration](aws-properties-bedrock-flow-retrievalflownodes3configuration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetrievalFlowNodeS3Configuration

S3Location

All content copied from https://docs.aws.amazon.com/.
