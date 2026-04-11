This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Cluster ClusterSettings

The settings to use when creating a cluster. This parameter is used to turn on
CloudWatch Container Insights with enhanced observability or CloudWatch Container
Insights for a cluster.

Container Insights with enhanced observability provides all the Container Insights
metrics, plus additional task and container metrics. This version supports enhanced
observability for Amazon ECS clusters using the Amazon EC2 and Fargate launch types.
After you configure Container Insights with enhanced observability on Amazon ECS,
Container Insights auto-collects detailed infrastructure telemetry from the cluster
level down to the container level in your environment and displays these critical
performance data in curated dashboards removing the heavy lifting in observability
set-up.

For more information, see [Monitor\
Amazon ECS containers using Container Insights with enhanced observability](../../../amazonecs/latest/developerguide/cloudwatch-container-insights.md)
in the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the cluster setting. The value is `containerInsights`.

_Required_: No

_Type_: String

_Allowed values_: `containerInsights`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value to set for the cluster setting. The supported values are
`enhanced`, `enabled`, and `disabled`.

To use Container Insights with enhanced observability, set the
`containerInsights` account setting to `enhanced`.

To use Container Insights, set the `containerInsights` account setting to
`enabled`.

If a cluster value is specified, it will override the `containerInsights`
value set with [PutAccountSetting](../../../../reference/amazonecs/latest/apireference/api-putaccountsetting.md) or [PutAccountSettingDefault](../../../../reference/amazonecs/latest/apireference/api-putaccountsettingdefault.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Define an empty cluster with CloudWatch Container Insights enabled and defined\
tags](../userguide/aws-resource-ecs-cluster.md#ws-resource-ecs-cluster--examples--Define_an_empty_cluster_with_CloudWatch_Container_Insights_enabled_and_defined_tags)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterConfiguration

ExecuteCommandConfiguration

All content copied from https://docs.aws.amazon.com/.
