This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain NodeOption

Configuration settings for defining the node type within a cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "NodeConfig" : NodeConfig,
  "NodeType" : String
}

```

### YAML

```yaml

  NodeConfig:
    NodeConfig
  NodeType: String

```

## Properties

`NodeConfig`

Configuration options for defining the setup of any node type.

_Required_: No

_Type_: [NodeConfig](aws-properties-opensearchservice-domain-nodeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NodeType`

Defines the type of node, such as coordinating nodes.

_Required_: No

_Type_: String

_Allowed values_: `coordinator`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NodeConfig

NodeToNodeEncryptionOptions

All content copied from https://docs.aws.amazon.com/.
