This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow RetrievalFlowNodeConfiguration

Contains configurations for a Retrieval node in a flow. This node retrieves data from the Amazon S3 location that you specify and returns it as the output.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ServiceConfiguration" : RetrievalFlowNodeServiceConfiguration
}

```

### YAML

```yaml

  ServiceConfiguration:
    RetrievalFlowNodeServiceConfiguration

```

## Properties

`ServiceConfiguration`

Contains configurations for the service to use for retrieving data to return as the output from the node.

_Required_: Yes

_Type_: [RetrievalFlowNodeServiceConfiguration](aws-properties-bedrock-flow-retrievalflownodeserviceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RerankingMetadataSelectiveModeConfiguration

RetrievalFlowNodeS3Configuration

All content copied from https://docs.aws.amazon.com/.
