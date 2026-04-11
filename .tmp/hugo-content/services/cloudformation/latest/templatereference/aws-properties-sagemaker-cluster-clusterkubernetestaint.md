This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterKubernetesTaint

A Kubernetes taint that can be applied to cluster nodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Effect" : String,
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Effect: String
  Key: String
  Value: String

```

## Properties

`Effect`

The effect of the taint. Valid values are
`NoSchedule`, `PreferNoSchedule`,
and `NoExecute`.

_Required_: Yes

_Type_: String

_Allowed values_: `NoSchedule | PreferNoSchedule | NoExecute`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The key of the taint.

_Required_: Yes

_Type_: String

_Pattern_: `([a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*/)?[A-Za-z0-9]([-A-Za-z0-9_.]*[A-Za-z0-9])?`

_Minimum_: `1`

_Maximum_: `317`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the taint.

_Required_: No

_Type_: String

_Pattern_: `(([A-Za-z0-9][-A-Za-z0-9_.]*)?[A-Za-z0-9])?`

_Minimum_: `1`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterKubernetesConfig

ClusterLifeCycleConfig

All content copied from https://docs.aws.amazon.com/.
