This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster Logging

Enable or disable exporting the Kubernetes control plane logs for your cluster to
CloudWatch Logs. By default, cluster control plane logs aren't exported to CloudWatch
Logs. For more information, see [Amazon EKS Cluster control plane\
logs](../../../eks/latest/userguide/control-plane-logs.md) in the _Amazon EKS User Guide_.

###### Important

When updating a resource, you must include this `Logging` property if
the previous CloudFormation template of the resource had it.

###### Note

CloudWatch Logs ingestion, archive storage, and data scanning rates apply to
exported control plane logs. For more information, see [CloudWatch Pricing](https://aws.amazon.com/cloudwatch/pricing).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ClusterLogging" : ClusterLogging
}

```

### YAML

```yaml

  ClusterLogging:
    ClusterLogging

```

## Properties

`ClusterLogging`

The cluster control plane logging configuration for your cluster.

_Required_: No

_Type_: [ClusterLogging](aws-properties-eks-cluster-clusterlogging.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KubernetesNetworkConfig

LoggingTypeConfig

All content copied from https://docs.aws.amazon.com/.
