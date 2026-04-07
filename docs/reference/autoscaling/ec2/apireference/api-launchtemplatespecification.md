# LaunchTemplateSpecification

Describes the launch template and the version of the launch template that Amazon EC2 Auto Scaling
uses to launch Amazon EC2 instances. For more information about launch templates, see [Launch\
templates](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-templates.html) in the _Amazon EC2 Auto Scaling User Guide_.

## Contents

**LaunchTemplateId**

The ID of the launch template. To get the template ID, use the Amazon EC2 [DescribeLaunchTemplates](../../../awsec2/latest/apireference/api-describelaunchtemplates.md) API operation. New launch templates can be created
using the Amazon EC2 [CreateLaunchTemplate](../../../awsec2/latest/apireference/api-createlaunchtemplate.md) API.

Conditional: You must specify either a `LaunchTemplateId` or a
`LaunchTemplateName`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

**LaunchTemplateName**

The name of the launch template. To get the template name, use the Amazon EC2 [DescribeLaunchTemplates](../../../awsec2/latest/apireference/api-describelaunchtemplates.md) API operation. New launch templates can be created
using the Amazon EC2 [CreateLaunchTemplate](../../../awsec2/latest/apireference/api-createlaunchtemplate.md) API.

Conditional: You must specify either a `LaunchTemplateId` or a
`LaunchTemplateName`.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `[a-zA-Z0-9\(\)\.\-/_]+`

Required: No

**Version**

The version number, `$Latest`, or `$Default`. To get the version
number, use the Amazon EC2 [DescribeLaunchTemplateVersions](../../../awsec2/latest/apireference/api-describelaunchtemplateversions.md) API operation. New launch template versions
can be created using the Amazon EC2 [CreateLaunchTemplateVersion](../../../awsec2/latest/apireference/api-createlaunchtemplateversion.md) API. If the value is `$Latest`,
Amazon EC2 Auto Scaling selects the latest version of the launch template when launching instances. If
the value is `$Default`, Amazon EC2 Auto Scaling selects the default version of the launch
template when launching instances. The default value is `$Default`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/autoscaling-2011-01-01/LaunchTemplateSpecification)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/autoscaling-2011-01-01/LaunchTemplateSpecification)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/autoscaling-2011-01-01/LaunchTemplateSpecification)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LaunchTemplateOverrides

LifecycleHook
