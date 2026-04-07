# WorkGroupSummary

The summary information for the workgroup, which includes its name, state,
description, and the date and time it was created.

## Contents

**CreationTime**

The workgroup creation date and time.

Type: Timestamp

Required: No

**Description**

The workgroup description.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1024.

Required: No

**EngineVersion**

The engine version setting for all queries on the workgroup. Queries on the
`AmazonAthenaPreviewFunctionality` workgroup run on the preview engine
regardless of this setting.

Type: [EngineVersion](api-engineversion.md) object

Required: No

**IdentityCenterApplicationArn**

The ARN of the IAM Identity Center enabled application associated with the
workgroup.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 255.

Pattern: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}$`

Required: No

**Name**

The name of the workgroup.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

**State**

The state of the workgroup.

Type: String

Valid Values: `ENABLED | DISABLED`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/WorkGroupSummary)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/WorkGroupSummary)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/WorkGroupSummary)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WorkGroupConfigurationUpdates

Common Parameters
