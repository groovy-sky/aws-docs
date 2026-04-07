# FleetLaunchTemplateSpecificationRequest

The Amazon EC2 launch template that can be used by
an EC2 Fleet to configure Amazon EC2 instances. You must specify either the ID or name of the launch template in the request, but not both.

For information about launch templates, see [Launch\
an instance from a launch template](../../../../services/ec2/latest/userguide/ec2-launch-templates.md) in the
_Amazon EC2 User Guide_.

## Contents

**LaunchTemplateId**

The ID of the launch template.

You must specify the `LaunchTemplateId` or the `LaunchTemplateName`, but not both.

Type: String

Required: No

**LaunchTemplateName**

The name of the launch template.

You must specify the `LaunchTemplateName` or the `LaunchTemplateId`, but not both.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `[a-zA-Z0-9\(\)\.\-/_]+`

Required: No

**Version**

The launch template version number, `$Latest`, or `$Default`. You must specify a value, otherwise the request fails.

If the value is `$Latest`, Amazon EC2 uses the latest version of the launch template.

If the value is `$Default`, Amazon EC2 uses the default version of the launch template.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/fleetlaunchtemplatespecificationrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/fleetlaunchtemplatespecificationrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/fleetlaunchtemplatespecificationrequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FleetLaunchTemplateSpecification

FleetSpotCapacityRebalance
