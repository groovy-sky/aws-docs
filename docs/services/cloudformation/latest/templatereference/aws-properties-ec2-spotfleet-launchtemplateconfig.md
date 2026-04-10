This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet LaunchTemplateConfig

Specifies a launch template and overrides.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateSpecification" : FleetLaunchTemplateSpecification,
  "Overrides" : [ LaunchTemplateOverrides, ... ]
}

```

### YAML

```yaml

  LaunchTemplateSpecification:
    FleetLaunchTemplateSpecification
  Overrides:
    - LaunchTemplateOverrides

```

## Properties

`LaunchTemplateSpecification`

The launch template to use. Make sure that the launch template does not contain the
`NetworkInterfaceId` parameter because you can't specify a network interface
ID in a Spot Fleet.

_Required_: No

_Type_: [FleetLaunchTemplateSpecification](aws-properties-ec2-spotfleet-fleetlaunchtemplatespecification.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Overrides`

Any parameters that you specify override the same parameters in the launch
template.

_Required_: No

_Type_: Array of [LaunchTemplateOverrides](aws-properties-ec2-spotfleet-launchtemplateoverrides.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceRequirementsRequest

LaunchTemplateOverrides

All content copied from https://docs.aws.amazon.com/.
