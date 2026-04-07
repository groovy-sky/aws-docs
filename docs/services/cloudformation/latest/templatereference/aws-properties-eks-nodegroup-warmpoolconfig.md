This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::Nodegroup WarmPoolConfig

The `WarmPoolConfig` property type specifies Property description not available. for an [AWS::EKS::Nodegroup](aws-resource-eks-nodegroup.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Enabled" : Boolean,
  "MaxGroupPreparedCapacity" : Integer,
  "MinSize" : Integer,
  "PoolState" : String,
  "ReuseOnScaleIn" : Boolean
}

```

### YAML

```yaml

  Enabled: Boolean
  MaxGroupPreparedCapacity: Integer
  MinSize: Integer
  PoolState: String
  ReuseOnScaleIn: Boolean

```

## Properties

`Enabled`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxGroupPreparedCapacity`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinSize`

Property description not available.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PoolState`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReuseOnScaleIn`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

UpdateConfig

AWS::EKS::PodIdentityAssociation
