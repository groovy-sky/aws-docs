# IamRoleConfiguration

The proposed access control configuration for an IAM role. You can propose a
configuration for a new IAM role or an existing IAM role that you own by specifying the
trust policy. If the configuration is for a new IAM role, you must specify the trust
policy. If the configuration is for an existing IAM role that you own and you do not
propose the trust policy, the access preview uses the existing trust policy for the role.
The proposed trust policy cannot be an empty string. For more information about role trust
policy limits, see [IAM and AWS STS\
quotas](../../../../services/iam/latest/userguide/reference-iam-quotas.md).

## Contents

**trustPolicy**

The proposed trust policy for the IAM role.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/iamroleconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/iamroleconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/iamroleconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GeneratedPolicyResult

InlineArchiveRule
