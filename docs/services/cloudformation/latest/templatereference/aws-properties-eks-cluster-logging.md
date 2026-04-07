This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster Logging

Enable or disable exporting the Kubernetes control plane logs for your cluster to
CloudWatch Logs. By default, cluster control plane logs aren't exported to CloudWatch
Logs. For more information, see [Amazon EKS Cluster control plane\
logs](https://docs.aws.amazon.com/eks/latest/userguide/control-plane-logs.html) in the _Amazon EKS User Guide_.

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

_Type_: [ClusterLogging](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-eks-cluster-clusterlogging.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KubernetesNetworkConfig

LoggingTypeConfig
