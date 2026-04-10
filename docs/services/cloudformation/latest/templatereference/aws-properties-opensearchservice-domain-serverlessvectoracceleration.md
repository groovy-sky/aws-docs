This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain ServerlessVectorAcceleration

Configuration for serverless vector acceleration, which provides [GPU-accelerated](../../../opensearch-service/latest/developerguide/gpu-acceleration-vector-index.md) vector
search capabilities for improved performance on vector workloads.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean

```

## Properties

`Enabled`

Specifies whether serverless vector acceleration is enabled for the domain.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SAMLOptions

ServiceSoftwareOptions

All content copied from https://docs.aws.amazon.com/.
