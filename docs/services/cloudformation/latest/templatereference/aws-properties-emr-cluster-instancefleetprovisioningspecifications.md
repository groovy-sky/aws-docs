This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster InstanceFleetProvisioningSpecifications

`InstanceFleetProvisioningSpecification` is a subproperty of `InstanceFleetConfig`. `InstanceFleetProvisioningSpecification` defines the launch specification for Spot instances in an instance fleet, which determines the defined duration and provisioning timeout behavior for Spot instances.

###### Note

The instance fleet configuration is available only in Amazon EMR versions 4.8.0 and later, excluding 5.0.x versions.

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

_Type_: [OnDemandProvisioningSpecification](aws-properties-emr-cluster-ondemandprovisioningspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotSpecification`

The launch specification for Spot instances in the fleet, which determines the allocation strategy, defined
duration, and provisioning timeout behavior.

_Required_: No

_Type_: [SpotProvisioningSpecification](aws-properties-emr-cluster-spotprovisioningspecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceFleetConfig

InstanceFleetResizingSpecifications

All content copied from https://docs.aws.amazon.com/.
