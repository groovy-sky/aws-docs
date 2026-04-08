# SecretsManagerSecretConfiguration

The configuration for a Secrets Manager secret. For more information, see [CreateSecret](../../../secretsmanager/latest/apireference/api-createsecret.md).

You can propose a configuration for a new secret or an existing secret that you own by
specifying the secret policy and optional AWS KMS encryption key. If the configuration is for
an existing secret and you do not specify the secret policy, the access preview uses the
existing policy for the secret. If the access preview is for a new resource and you do not
specify the policy, the access preview assumes a secret without a policy. To propose
deletion of an existing policy, you can specify an empty string. If the proposed
configuration is for a new secret and you do not specify the KMS key ID, the access
preview uses the AWS managed key `aws/secretsmanager`. If you specify an empty
string for the KMS key ID, the access preview uses the AWS managed key of the
AWS account. For more information about secret policy limits, see [Quotas\
for AWS Secrets Manager.](../../../../services/secretsmanager/latest/userguide/reference-limits.md).

## Contents

**kmsKeyId**

The proposed ARN, key ID, or alias of the KMS key.

Type: String

Required: No

**secretPolicy**

The proposed resource policy defining who can access or manage the secret.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/secretsmanagersecretconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/secretsmanagersecretconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/secretsmanagersecretconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3PublicAccessBlockConfiguration

SnsTopicConfiguration
