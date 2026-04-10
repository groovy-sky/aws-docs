This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup AvailabilityZoneImpairmentPolicy

Describes an Availability Zone impairment policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ImpairedZoneHealthCheckBehavior" : String,
  "ZonalShiftEnabled" : Boolean
}

```

### YAML

```yaml

  ImpairedZoneHealthCheckBehavior: String
  ZonalShiftEnabled: Boolean

```

## Properties

`ImpairedZoneHealthCheckBehavior`

Specifies the health check behavior for the impaired Availability Zone in an active zonal shift. If you select `Replace unhealthy`, instances that appear unhealthy will be replaced in all Availability Zones.
If you select `Ignore unhealthy`, instances will not be replaced in the Availability Zone with the active zonal shift. For more information, see [Auto Scaling group zonal shift](../../../autoscaling/ec2/userguide/ec2-auto-scaling-zonal-shift.md)
in the _Amazon EC2 Auto Scaling User Guide_.

_Required_: Yes

_Type_: String

_Allowed values_: `IgnoreUnhealthy | ReplaceUnhealthy`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ZonalShiftEnabled`

If `true`, enable zonal shift for your Auto Scaling group.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailabilityZoneDistribution

BaselineEbsBandwidthMbpsRequest

All content copied from https://docs.aws.amazon.com/.
