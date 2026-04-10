This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::Fleet ServiceManagedEc2FleetConfiguration

The configuration details for a service managed EC2 fleet.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoScalingConfiguration" : ServiceManagedEc2AutoScalingConfiguration,
  "InstanceCapabilities" : ServiceManagedEc2InstanceCapabilities,
  "InstanceMarketOptions" : ServiceManagedEc2InstanceMarketOptions,
  "StorageProfileId" : String,
  "VpcConfiguration" : VpcConfiguration
}

```

### YAML

```yaml

  AutoScalingConfiguration:
    ServiceManagedEc2AutoScalingConfiguration
  InstanceCapabilities:
    ServiceManagedEc2InstanceCapabilities
  InstanceMarketOptions:
    ServiceManagedEc2InstanceMarketOptions
  StorageProfileId: String
  VpcConfiguration:
    VpcConfiguration

```

## Properties

`AutoScalingConfiguration`

Property description not available.

_Required_: No

_Type_: [ServiceManagedEc2AutoScalingConfiguration](aws-properties-deadline-fleet-servicemanagedec2autoscalingconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceCapabilities`

The instance capabilities for the service managed EC2 fleet.

_Required_: Yes

_Type_: [ServiceManagedEc2InstanceCapabilities](aws-properties-deadline-fleet-servicemanagedec2instancecapabilities.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceMarketOptions`

The instance market options for the service managed EC2 fleet.

_Required_: Yes

_Type_: [ServiceManagedEc2InstanceMarketOptions](aws-properties-deadline-fleet-servicemanagedec2instancemarketoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageProfileId`

The storage profile ID for the service managed EC2 fleet.

_Required_: No

_Type_: String

_Pattern_: `^sp-[0-9a-f]{32}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcConfiguration`

The VPC configuration for the service managed EC2 fleet.

_Required_: No

_Type_: [VpcConfiguration](aws-properties-deadline-fleet-vpcconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceManagedEc2AutoScalingConfiguration

ServiceManagedEc2InstanceCapabilities

All content copied from https://docs.aws.amazon.com/.
