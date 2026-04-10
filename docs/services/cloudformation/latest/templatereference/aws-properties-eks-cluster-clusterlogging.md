This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Cluster ClusterLogging

The cluster control plane logging configuration for your cluster.

###### Important

When updating a resource, you must include this `ClusterLogging`
property if the previous CloudFormation template of the resource had it.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnabledTypes" : [ LoggingTypeConfig, ... ]
}

```

### YAML

```yaml

  EnabledTypes:
    - LoggingTypeConfig

```

## Properties

`EnabledTypes`

The enabled control plane logs for your cluster. All log types are disabled if the
array is empty.

###### Important

When updating a resource, you must include this `EnabledTypes` property
if the previous CloudFormation template of the resource had it.

_Required_: No

_Type_: Array of [LoggingTypeConfig](aws-properties-eks-cluster-loggingtypeconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BlockStorage

ComputeConfig

All content copied from https://docs.aws.amazon.com/.
