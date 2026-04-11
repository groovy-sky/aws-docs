# IdcAuthConfiguration

Information about the AWS IAM Identity Center Application used to configure authentication for a plugin.

## Contents

**idcApplicationArn**

The Amazon Resource Name (ARN) of the AWS IAM Identity Center Application used to configure authentication.

Type: String

Length Constraints: Minimum length of 10. Maximum length of 1224.

Pattern: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso::\d{12}:application/(sso)?ins-[a-zA-Z0-9-.]{16}/apl-[a-zA-Z0-9]{16}`

Required: Yes

**roleArn**

The Amazon Resource Name (ARN) of the IAM role with permissions to perform actions on AWS services
on your behalf.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 1284.

Pattern: `arn:[a-z0-9-\.]{1,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[a-z0-9-\.]{0,63}:[^/].{0,1023}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qbusiness-2023-11-27/idcauthconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qbusiness-2023-11-27/idcauthconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qbusiness-2023-11-27/idcauthconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HookConfiguration

IdentityProviderConfiguration

All content copied from https://docs.aws.amazon.com/.
