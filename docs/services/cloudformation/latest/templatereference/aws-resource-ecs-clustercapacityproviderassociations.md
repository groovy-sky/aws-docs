This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::ClusterCapacityProviderAssociations

The `AWS::ECS::ClusterCapacityProviderAssociations` resource associates one
or more capacity providers and a default capacity provider strategy with a cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECS::ClusterCapacityProviderAssociations",
  "Properties" : {
      "CapacityProviders" : [ String, ... ],
      "Cluster" : String,
      "DefaultCapacityProviderStrategy" : [ CapacityProviderStrategy, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ECS::ClusterCapacityProviderAssociations
Properties:
  CapacityProviders:
    - String
  Cluster: String
  DefaultCapacityProviderStrategy:
    - CapacityProviderStrategy

```

## Properties

`CapacityProviders`

The capacity providers to associate with the cluster.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Cluster`

The cluster the capacity provider association is the target of.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DefaultCapacityProviderStrategy`

The default capacity provider strategy to associate with the cluster.

_Required_: Yes

_Type_: Array of [CapacityProviderStrategy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-clustercapacityproviderassociations-capacityproviderstrategy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the cluster name.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

CapacityProviderStrategy
