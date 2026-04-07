# LaunchTemplate

Describes a launch template.

## Contents

**createdBy**

The principal that created the launch template.

Type: String

Required: No

**createTime**

The time launch template was created.

Type: Timestamp

Required: No

**defaultVersionNumber**

The version number of the default version of the launch template.

Type: Long

Required: No

**latestVersionNumber**

The version number of the latest version of the launch template.

Type: Long

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

**TagSet.N**

The tags for the launch template.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplate.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchSpecification

LaunchTemplateAndOverridesResponse
