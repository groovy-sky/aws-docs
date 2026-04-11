# LaunchTemplateConfig

Describes a launch template and overrides.

## Contents

**LaunchTemplateSpecification** (request), **launchTemplateSpecification** (response)

The launch template to use. Make sure that the launch template does not contain the
`NetworkInterfaceId` parameter because you can't specify a network interface
ID in a Spot Fleet.

Type: [FleetLaunchTemplateSpecification](api-fleetlaunchtemplatespecification.md) object

Required: No

**Overrides.N**

Any parameters that you specify override the same parameters in the launch
template.

Type: Array of [LaunchTemplateOverrides](api-launchtemplateoverrides.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateCapacityReservationSpecificationResponse

LaunchTemplateCpuOptions
