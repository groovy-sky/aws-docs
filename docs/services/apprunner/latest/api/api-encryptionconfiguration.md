# EncryptionConfiguration

Describes a custom encryption key that AWS App Runner uses to encrypt copies of the source repository and service logs.

## Contents

**KmsKey**

The ARN of the KMS key that's used for encryption.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `arn:aws(-[\w]+)*:kms:[a-z\-]+-[0-9]{1}:[0-9]{12}:key\/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/encryptionconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/encryptionconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/encryptionconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EgressConfiguration

HealthCheckConfiguration

All content copied from https://docs.aws.amazon.com/.
