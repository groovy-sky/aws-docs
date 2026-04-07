# FleetLaunchTemplateConfigRequest

Describes a launch template and overrides.

## Contents

**LaunchTemplateSpecification**

The launch template to use. You must specify either the launch template ID or launch
template name in the request.

Type: [FleetLaunchTemplateSpecificationRequest](api-fleetlaunchtemplatespecificationrequest.md) object

Required: No

**Overrides.N**

Any parameters that you specify override the same parameters in the launch
template.

For fleets of type `request` and `maintain`, a maximum of 300
items is allowed across all launch templates.

Type: Array of [FleetLaunchTemplateOverridesRequest](api-fleetlaunchtemplateoverridesrequest.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/fleetlaunchtemplateconfigrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/fleetlaunchtemplateconfigrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/fleetlaunchtemplateconfigrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FleetLaunchTemplateConfig

FleetLaunchTemplateOverrides
