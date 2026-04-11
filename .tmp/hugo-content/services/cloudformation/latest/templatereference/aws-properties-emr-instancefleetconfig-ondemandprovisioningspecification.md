This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::InstanceFleetConfig OnDemandProvisioningSpecification

The launch specification for On-Demand Instances in the instance fleet, which
determines the allocation strategy.

###### Note

The instance fleet configuration is available only in Amazon EMR releases
4.8.0 and later, excluding 5.0.x versions. On-Demand Instances allocation strategy is
available in Amazon EMR releases 5.12.1 and later.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "CapacityReservationOptions" : OnDemandCapacityReservationOptions
}

```

### YAML

```yaml

  AllocationStrategy: String
  CapacityReservationOptions:
    OnDemandCapacityReservationOptions

```

## Properties

`AllocationStrategy`

Specifies the strategy to use in launching On-Demand instance fleets. Available
options are `lowest-price` and `prioritized`. `lowest-price`
specifies to launch the instances with the lowest price first, and `prioritized` specifies
that Amazon EMR should launch the instances with the highest priority first. The default is
`lowest-price`.

_Required_: Yes

_Type_: String

_Allowed values_: `lowest-price | prioritized`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CapacityReservationOptions`

The launch specification for On-Demand instances in the instance fleet, which determines
the allocation strategy.

_Required_: No

_Type_: [OnDemandCapacityReservationOptions](aws-properties-emr-instancefleetconfig-ondemandcapacityreservationoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnDemandCapacityReservationOptions

OnDemandResizingSpecification

All content copied from https://docs.aws.amazon.com/.
