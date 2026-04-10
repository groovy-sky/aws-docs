This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Cluster SpotResizingSpecification

The resize specification for Spot Instances in the instance fleet, which contains the
resize timeout period.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllocationStrategy" : String,
  "TimeoutDurationMinutes" : Integer
}

```

### YAML

```yaml

  AllocationStrategy: String
  TimeoutDurationMinutes: Integer

```

## Properties

`AllocationStrategy`

Specifies the allocation strategy to use to launch Spot instances during a resize. If you run Amazon EMR releases 6.9.0 or higher,
the default is `price-capacity-optimized`. If you run Amazon EMR releases 6.8.0 or lower, the default is
`capacity-optimized`.

_Required_: No

_Type_: String

_Allowed values_: `capacity-optimized | price-capacity-optimized | lowest-price | diversified | capacity-optimized-prioritized`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutDurationMinutes`

Spot resize timeout in minutes. If Spot Instances are not provisioned within this time,
the resize workflow will stop provisioning of Spot instances. Minimum value is 5 minutes
and maximum value is 10,080 minutes (7 days). The timeout applies to all resize workflows
on the Instance Fleet. The resize could be triggered by Amazon EMR Managed Scaling
or by the customer (via Amazon EMR Console, Amazon EMR CLI
modify-instance-fleet or Amazon EMR SDK ModifyInstanceFleet API) or by Amazon EMR due to Amazon EC2 Spot Reclamation.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SpotProvisioningSpecification

StepConfig

All content copied from https://docs.aws.amazon.com/.
