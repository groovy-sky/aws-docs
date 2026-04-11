This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterKubernetesConfig

Kubernetes configuration that specifies labels and taints to be applied
to cluster nodes in an instance group.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Labels" : {Key: Value, ...},
  "Taints" : [ ClusterKubernetesTaint, ... ]
}

```

### YAML

```yaml

  Labels:
    Key: Value
  Taints:
    - ClusterKubernetesTaint

```

## Properties

`Labels`

Key-value pairs of labels to be applied to cluster nodes.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Taints`

List of taints to be applied to cluster nodes.

_Required_: No

_Type_: Array of [ClusterKubernetesTaint](aws-properties-sagemaker-cluster-clusterkubernetestaint.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterInstanceStorageConfig

ClusterKubernetesTaint

All content copied from https://docs.aws.amazon.com/.
