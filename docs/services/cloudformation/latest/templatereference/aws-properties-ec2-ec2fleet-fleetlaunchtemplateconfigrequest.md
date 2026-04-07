This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::EC2Fleet FleetLaunchTemplateConfigRequest

Specifies a launch template and overrides for an EC2 Fleet.

`FleetLaunchTemplateConfigRequest` is a property of the [AWS::EC2::EC2Fleet](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-ec2fleet.html) resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateSpecification" : FleetLaunchTemplateSpecificationRequest,
  "Overrides" : [ FleetLaunchTemplateOverridesRequest, ... ]
}

```

### YAML

```yaml

  LaunchTemplateSpecification:
    FleetLaunchTemplateSpecificationRequest
  Overrides:
    - FleetLaunchTemplateOverridesRequest

```

## Properties

`LaunchTemplateSpecification`

The launch template to use. You must specify either the launch template ID or launch
template name in the request.

_Required_: No

_Type_: [FleetLaunchTemplateSpecificationRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ec2fleet-fleetlaunchtemplatespecificationrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Overrides`

Any parameters that you specify override the same parameters in the launch
template.

For fleets of type `request` and `maintain`, a maximum of 300
items is allowed across all launch templates.

_Required_: No

_Type_: Array of [FleetLaunchTemplateOverridesRequest](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-ec2fleet-fleetlaunchtemplateoverridesrequest.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## See also

- [FleetLaunchTemplateConfigRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FleetLaunchTemplateConfigRequest.html) in the _Amazon EC2 API_
_Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EbsBlockDevice

FleetLaunchTemplateOverridesRequest
