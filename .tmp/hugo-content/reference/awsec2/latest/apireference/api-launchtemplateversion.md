# LaunchTemplateVersion

Describes a launch template version.

## Contents

**createdBy**

The principal that created the version.

Type: String

Required: No

**createTime**

The time the version was created.

Type: Timestamp

Required: No

**defaultVersion**

Indicates whether the version is the default version.

Type: Boolean

Required: No

**launchTemplateData**

Information about the launch template.

Type: [ResponseLaunchTemplateData](api-responselaunchtemplatedata.md) object

Required: No

**launchTemplateId**

The ID of the launch template.

Type: String

Required: No

**launchTemplateName**

The name of the launch template.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `[a-zA-Z0-9\(\)\.\-/_]+`

Required: No

**operator**

The entity that manages the launch template.

Type: [OperatorResponse](api-operatorresponse.md) object

Required: No

**versionDescription**

The description for the version.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 255.

Required: No

**versionNumber**

The version number.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplateversion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplateversion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplateversion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplateTagSpecificationRequest

LicenseConfiguration
