# ManagedQueryResultsEncryptionConfiguration

If you encrypt query and calculation results in Athena owned storage, this field
indicates the encryption option (for example, SSE\_KMS or CSE\_KMS) and key
information.

## Contents

**KmsKey**

The ARN of an AWS KMS key for encrypting managed query results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `^arn:aws[a-z\-]*:kms:([a-z0-9\-]+):\d{12}:key/?[a-zA-Z_0-9+=,.@\-_/]+$|^arn:aws[a-z\-]*:kms:([a-z0-9\-]+):\d{12}:alias/?[a-zA-Z_0-9+=,.@\-_/]+$|^alias/[a-zA-Z0-9/_-]+$|[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/managedqueryresultsencryptionconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/managedqueryresultsencryptionconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/managedqueryresultsencryptionconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ManagedQueryResultsConfigurationUpdates

MonitoringConfiguration
