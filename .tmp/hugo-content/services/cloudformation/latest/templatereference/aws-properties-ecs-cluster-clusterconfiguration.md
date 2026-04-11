This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ClusterConfiguration

The execute command and managed storage configuration for the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecuteCommandConfiguration" : ExecuteCommandConfiguration,
  "ManagedStorageConfiguration" : ManagedStorageConfiguration
}

```

### YAML

```yaml

  ExecuteCommandConfiguration:
    ExecuteCommandConfiguration
  ManagedStorageConfiguration:
    ManagedStorageConfiguration

```

## Properties

`ExecuteCommandConfiguration`

The details of the execute command configuration.

_Required_: No

_Type_: [ExecuteCommandConfiguration](aws-properties-ecs-cluster-executecommandconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedStorageConfiguration`

The details of the managed storage configuration.

_Required_: No

_Type_: [ManagedStorageConfiguration](aws-properties-ecs-cluster-managedstorageconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Define a cluster with the AWS Fargate capacity providers and\
a default capacity provider strategy defined](../userguide/aws-resource-ecs-cluster.md#aws-resource-ecs-cluster--examples--Define_a_cluster_with_the__capacity_providers_and_a_default_capacity_provider_strategy_defined)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CapacityProviderStrategyItem

ClusterSettings

All content copied from https://docs.aws.amazon.com/.
