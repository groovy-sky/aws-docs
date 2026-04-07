# FleetLaunchTemplateConfigRequest

Describes a launch template and overrides.

## Contents

**LaunchTemplateSpecification**

The launch template to use. You must specify either the launch template ID or launch
template name in the request.

Type: [FleetLaunchTemplateSpecificationRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FleetLaunchTemplateSpecificationRequest.html) object

Required: No

**Overrides.N**

Any parameters that you specify override the same parameters in the launch
template.

For fleets of type `request` and `maintain`, a maximum of 300
items is allowed across all launch templates.

Type: Array of [FleetLaunchTemplateOverridesRequest](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_FleetLaunchTemplateOverridesRequest.html) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/FleetLaunchTemplateConfigRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/FleetLaunchTemplateConfigRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/FleetLaunchTemplateConfigRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FleetLaunchTemplateConfig

FleetLaunchTemplateOverrides
