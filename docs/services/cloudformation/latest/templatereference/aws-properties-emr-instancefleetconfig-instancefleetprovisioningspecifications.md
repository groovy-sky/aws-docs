This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig InstanceFleetProvisioningSpecifications

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

`InstanceTypeConfig` is a sub-property of `InstanceFleetConfig`. `InstanceTypeConfig` determines the EC2 instances that Amazon EMR attempts to provision to fulfill On-Demand and Spot target capacities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OnDemandSpecification" : OnDemandProvisioningSpecification,
  "SpotSpecification" : SpotProvisioningSpecification
}

```

### YAML

```yaml

  OnDemandSpecification:
    OnDemandProvisioningSpecification
  SpotSpecification:
    SpotProvisioningSpecification

```

## Properties

`OnDemandSpecification`

The launch specification for On-Demand Instances in the instance fleet, which
determines the allocation strategy and capacity reservation options.

###### Note

The instance fleet configuration is available only in Amazon EMR releases
4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is
available in Amazon EMR releases 5.12.1 and later.

_Required_: No

_Type_: [OnDemandProvisioningSpecification](aws-properties-emr-instancefleetconfig-ondemandprovisioningspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpotSpecification`

The launch specification for Spot instances in the fleet, which determines the allocation strategy, defined
duration, and provisioning timeout behavior.

_Required_: No

_Type_: [SpotProvisioningSpecification](aws-properties-emr-instancefleetconfig-spotprovisioningspecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbsConfiguration

InstanceFleetResizingSpecifications

All content copied from https://docs.aws.amazon.com/.
