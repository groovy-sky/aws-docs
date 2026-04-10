This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Cluster ClusterCapacityRequirements

Defines the instance capacity requirements for an instance group,
including configurations for both Spot and On-Demand capacity types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnDemand" : Json,
  "Spot" : Json
}

```

### YAML

```yaml

  OnDemand: Json
  Spot: Json

```

## Properties

`OnDemand`

Configuration options specific to On-Demand instances.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Spot`

Configuration options specific to Spot instances.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClusterAutoScalingConfig

ClusterEbsVolumeConfig

All content copied from https://docs.aws.amazon.com/.
