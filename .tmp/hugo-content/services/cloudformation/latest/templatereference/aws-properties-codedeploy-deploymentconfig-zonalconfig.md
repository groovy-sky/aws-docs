This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig ZonalConfig

Configure the `ZonalConfig` object if you want AWS CodeDeploy to
deploy your application to one [Availability Zone](../../../ec2/latest/userguide/using-regions-availability-zones.md#concepts-availability-zones) at a time, within an AWS Region. By
deploying to one Availability Zone at a time, you can expose your deployment to a
progressively larger audience as confidence in the deployment's performance and
viability grows. If you don't configure the `ZonalConfig` object, CodeDeploy deploys your application to a random selection of hosts across a
Region.

For more information about the zonal configuration feature, see [zonal configuration](../../../codedeploy/latest/userguide/deployment-configurations-create.md#zonal-config) in the _CodeDeploy User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FirstZoneMonitorDurationInSeconds" : Integer,
  "MinimumHealthyHostsPerZone" : MinimumHealthyHostsPerZone,
  "MonitorDurationInSeconds" : Integer
}

```

### YAML

```yaml

  FirstZoneMonitorDurationInSeconds: Integer
  MinimumHealthyHostsPerZone:
    MinimumHealthyHostsPerZone
  MonitorDurationInSeconds: Integer

```

## Properties

`FirstZoneMonitorDurationInSeconds`

The period of time, in seconds, that CodeDeploy must wait after completing a
deployment to the _first_ Availability Zone. CodeDeploy will
wait this amount of time before starting a deployment to the second Availability Zone.
You might set this option if you want to allow extra bake time for the first
Availability Zone. If you don't specify a value for
`firstZoneMonitorDurationInSeconds`, then CodeDeploy uses the
`monitorDurationInSeconds` value for the first Availability Zone.

For more information about the zonal configuration feature, see [zonal configuration](../../../codedeploy/latest/userguide/deployment-configurations-create.md#zonal-config) in the _CodeDeploy User_
_Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MinimumHealthyHostsPerZone`

The number or percentage of instances that must remain available per Availability Zone
during a deployment. This option works in conjunction with the
`MinimumHealthyHosts` option. For more information, see [About the minimum number of healthy hosts per Availability Zone](../../../codedeploy/latest/userguide/instances-health.md#minimum-healthy-hosts-az) in the
_CodeDeploy User Guide_.

If you don't specify the `minimumHealthyHostsPerZone` option, then CodeDeploy uses a default value of `0` percent.

For more information about the zonal configuration feature, see [zonal configuration](../../../codedeploy/latest/userguide/deployment-configurations-create.md#zonal-config) in the _CodeDeploy User_
_Guide_.

_Required_: No

_Type_: [MinimumHealthyHostsPerZone](aws-properties-codedeploy-deploymentconfig-minimumhealthyhostsperzone.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MonitorDurationInSeconds`

The period of time, in seconds, that CodeDeploy must wait after completing a
deployment to an Availability Zone. CodeDeploy will wait this amount of time
before starting a deployment to the next Availability Zone. Consider adding a monitor
duration to give the deployment some time to prove itself (or 'bake') in one
Availability Zone before it is released in the next zone. If you don't specify a
`monitorDurationInSeconds`, CodeDeploy starts deploying to the
next Availability Zone immediately.

For more information about the zonal configuration feature, see [zonal configuration](../../../codedeploy/latest/userguide/deployment-configurations-create.md#zonal-config) in the _CodeDeploy User_
_Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TrafficRoutingConfig

AWS::CodeDeploy::DeploymentGroup

All content copied from https://docs.aws.amazon.com/.
