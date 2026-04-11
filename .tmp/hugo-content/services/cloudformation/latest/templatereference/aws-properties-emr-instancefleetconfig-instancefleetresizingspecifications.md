This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig InstanceFleetResizingSpecifications

The resize specification for On-Demand and Spot Instances in the fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnDemandResizeSpecification" : OnDemandResizingSpecification,
  "SpotResizeSpecification" : SpotResizingSpecification
}

```

### YAML

```yaml

  OnDemandResizeSpecification:
    OnDemandResizingSpecification
  SpotResizeSpecification:
    SpotResizingSpecification

```

## Properties

`OnDemandResizeSpecification`

The resize specification for On-Demand Instances in the instance fleet, which contains
the allocation strategy, capacity reservation options, and the resize timeout period.

_Required_: No

_Type_: [OnDemandResizingSpecification](aws-properties-emr-instancefleetconfig-ondemandresizingspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotResizeSpecification`

The resize specification for Spot Instances in the instance fleet, which contains the
allocation strategy and the resize timeout period.

_Required_: No

_Type_: [SpotResizingSpecification](aws-properties-emr-instancefleetconfig-spotresizingspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceFleetProvisioningSpecifications

InstanceTypeConfig

All content copied from https://docs.aws.amazon.com/.
