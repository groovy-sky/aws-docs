This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster OnDemandResizingSpecification

The resize specification for On-Demand Instances in the instance fleet, which contains
the resize timeout period.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "CapacityReservationOptions" : OnDemandCapacityReservationOptions,
  "TimeoutDurationMinutes" : Integer
}

```

### YAML

```yaml

  AllocationStrategy: String
  CapacityReservationOptions:
    OnDemandCapacityReservationOptions
  TimeoutDurationMinutes: Integer

```

## Properties

`AllocationStrategy`

Specifies the allocation strategy to use to launch On-Demand instances during a resize. The default is `lowest-price`.

_Required_: No

_Type_: String

_Allowed values_: `lowest-price | prioritized`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CapacityReservationOptions`

Property description not available.

_Required_: No

_Type_: [OnDemandCapacityReservationOptions](aws-properties-emr-cluster-ondemandcapacityreservationoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutDurationMinutes`

On-Demand resize timeout in minutes. If On-Demand Instances are not provisioned within
this time, the resize workflow stops. The minimum value is 5 minutes, and the maximum value
is 10,080 minutes (7 days). The timeout applies to all resize workflows on the Instance
Fleet. The resize could be triggered by Amazon EMR Managed Scaling or by the
customer (via Amazon EMR Console, Amazon EMR CLI modify-instance-fleet or
Amazon EMR SDK ModifyInstanceFleet API) or by Amazon EMR due to Amazon EC2 Spot Reclamation.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

OnDemandProvisioningSpecification

PlacementGroupConfig

All content copied from https://docs.aws.amazon.com/.
