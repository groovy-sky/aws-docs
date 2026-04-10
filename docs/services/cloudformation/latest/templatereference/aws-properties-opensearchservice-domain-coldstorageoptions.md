This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain ColdStorageOptions

Container for the parameters required to enable cold storage for an OpenSearch Service
domain. For more information, see [Cold storage for\
Amazon OpenSearch Service](../../../opensearch-service/latest/developerguide/cold-storage.md).

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

Whether to enable or disable cold storage on the domain. You must enable UltraWarm
storage to enable cold storage.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CognitoOptions

DeploymentStrategyOptions

All content copied from https://docs.aws.amazon.com/.
