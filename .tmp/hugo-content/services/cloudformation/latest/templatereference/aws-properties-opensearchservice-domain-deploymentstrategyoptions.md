This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::OpenSearchService::Domain DeploymentStrategyOptions

Specifies the deployment strategy options for the domain.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DeploymentStrategy" : String
}

```

### YAML

```yaml

  DeploymentStrategy: String

```

## Properties

`DeploymentStrategy`

Specifies the deployment strategy for the domain. Valid values are `Default` and `CapacityOptimized`.

_Required_: No

_Type_: String

_Allowed values_: `Default | CapacityOptimized`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ColdStorageOptions

DomainEndpointOptions

All content copied from https://docs.aws.amazon.com/.
