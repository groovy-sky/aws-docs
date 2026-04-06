# IdentityCenterConfiguration

Specifies whether the workgroup is IAM Identity Center supported.

## Contents

**EnableIdentityCenter**

Specifies whether the workgroup is IAM Identity Center supported.

Type: Boolean

Required: No

**IdentityCenterInstanceArn**

The IAM Identity Center instance ARN that the workgroup associates to.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 255.

Pattern: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/IdentityCenterConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/IdentityCenterConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/IdentityCenterConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FilterDefinition

ListNamedQueriesInput
