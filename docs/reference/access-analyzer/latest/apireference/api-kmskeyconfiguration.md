# KmsKeyConfiguration

Proposed access control configuration for a KMS key. You can propose a configuration
for a new KMS key or an existing KMS key that you own by specifying the key policy and
AWS KMS grant configuration. If the configuration is for an existing key and you do not
specify the key policy, the access preview uses the existing policy for the key. If the
access preview is for a new resource and you do not specify the key policy, then the access
preview uses the default key policy. The proposed key policy cannot be an empty string. For
more information, see [Default key\
policy](../../../../services/kms/latest/developerguide/key-policies.md#key-policy-default). For more information about key policy limits, see [Resource\
quotas](../../../../services/kms/latest/developerguide/resource-limits.md).

## Contents

**grants**

A list of proposed grant configurations for the KMS key. If the proposed grant
configuration is for an existing key, the access preview uses the proposed list of grant
configurations in place of the existing grants. Otherwise, the access preview uses the
existing grants for the key.

Type: Array of [KmsGrantConfiguration](api-kmsgrantconfiguration.md) objects

Required: No

**keyPolicies**

Resource policy configuration for the KMS key. The only valid value for the name of
the key policy is `default`. For more information, see [Default key\
policy](../../../../services/kms/latest/developerguide/key-policies.md#key-policy-default).

Type: String to string map

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/kmskeyconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/kmskeyconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/kmskeyconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

KmsGrantConstraints

Location
