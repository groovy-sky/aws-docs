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

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/identitycenterconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/identitycenterconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/identitycenterconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterDefinition

ListNamedQueriesInput

All content copied from https://docs.aws.amazon.com/.
