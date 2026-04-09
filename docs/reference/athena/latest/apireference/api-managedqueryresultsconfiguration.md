# ManagedQueryResultsConfiguration

The configuration for storing results in Athena owned storage, which includes whether this feature is enabled; whether encryption configuration, if any, is used for encrypting query results.

## Contents

**Enabled**

If set to true, allows you to store query results in Athena owned storage. If set to
false, workgroup member stores query results in location specified under
`ResultConfiguration$OutputLocation`. The default is false. A workgroup
cannot have the `ResultConfiguration$OutputLocation` parameter when you set
this field to true.

Type: Boolean

Required: Yes

**EncryptionConfiguration**

If you encrypt query and calculation results in Athena owned storage, this field
indicates the encryption option (for example, SSE\_KMS or CSE\_KMS) and key
information.

Type: [ManagedQueryResultsEncryptionConfiguration](api-managedqueryresultsencryptionconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/managedqueryresultsconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/managedqueryresultsconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/managedqueryresultsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ManagedLoggingConfiguration

ManagedQueryResultsConfigurationUpdates

All content copied from https://docs.aws.amazon.com/.
