# ImageUsageResourceTypeRequest

A resource type to include in the report. Associated options can also be specified if the
resource type is a launch template.

## Contents

**ResourceType**

The resource type.

Valid values: `ec2:Instance` \| `ec2:LaunchTemplate`

Type: String

Required: No

**ResourceTypeOption.N**

The options that affect the scope of the report. Valid only when `ResourceType`
is `ec2:LaunchTemplate`.

Type: Array of [ImageUsageResourceTypeOptionRequest](api-imageusageresourcetypeoptionrequest.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/imageusageresourcetyperequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/imageusageresourcetyperequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/imageusageresourcetyperequest.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImageUsageResourceTypeOptionRequest

ImportImageLicenseConfigurationRequest
