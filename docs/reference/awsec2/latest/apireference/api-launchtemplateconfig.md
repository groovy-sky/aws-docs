# LaunchTemplateConfig

Describes a launch template and overrides.

## Contents

**LaunchTemplateSpecification** (request), **launchTemplateSpecification** (response)

The launch template to use. Make sure that the launch template does not contain the
`NetworkInterfaceId` parameter because you can't specify a network interface
ID in a Spot Fleet.

Type: [FleetLaunchTemplateSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FleetLaunchTemplateSpecification.html) object

Required: No

**Overrides.N**

Any parameters that you specify override the same parameters in the launch
template.

Type: Array of [LaunchTemplateOverrides](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchTemplateOverrides.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/LaunchTemplateConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/LaunchTemplateConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/LaunchTemplateConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateCapacityReservationSpecificationResponse

LaunchTemplateCpuOptions
