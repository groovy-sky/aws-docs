# LaunchTemplateSpecification

Describes the launch template to use.

## Contents

**LaunchTemplateId**

The ID of the launch template.

You must specify either the launch template ID or the
launch template name, but not both.

Type: String

Required: No

**LaunchTemplateName**

The name of the launch template.

You must specify either the launch template ID or the
launch template name, but not both.

Type: String

Required: No

**Version**

The launch template version number, `$Latest`, or
`$Default`.

A value of `$Latest` uses the latest version of the launch template.

A value of `$Default` uses the default version of the launch template.

Default: The default version of the launch template.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/launchtemplatespecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/launchtemplatespecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/launchtemplatespecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

LaunchTemplatesMonitoringRequest

LaunchTemplateSpotMarketOptions
