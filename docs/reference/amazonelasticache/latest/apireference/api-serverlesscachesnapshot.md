# ServerlessCacheSnapshot

The resource representing a serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

## Contents

###### Note

In the following list, the required parameters are described first.

**ARN**

The Amazon Resource Name (ARN) of a serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**BytesUsedForCache**

The total size of a serverless cache snapshot, in bytes. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**CreateTime**

The date and time that the source serverless cache's metadata and cache data set was obtained for
the snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: Timestamp

Required: No

**ExpiryTime**

The time that the serverless cache snapshot will expire. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: Timestamp

Required: No

**KmsKeyId**

The ID of the AWS Key Management Service (KMS) key of a serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**ServerlessCacheConfiguration**

The configuration of the serverless cache, at the time the snapshot was taken. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: [ServerlessCacheConfiguration](api-serverlesscacheconfiguration.md) object

Required: No

**ServerlessCacheSnapshotName**

The identifier of a serverless cache snapshot. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**SnapshotType**

The type of snapshot of serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

**Status**

The current status of the serverless cache. Available for Valkey, Redis OSS and Serverless Memcached only.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/elasticache-2015-02-02/serverlesscachesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/elasticache-2015-02-02/serverlesscachesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/elasticache-2015-02-02/serverlesscachesnapshot.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServerlessCacheConfiguration

ServiceUpdate

All content copied from https://docs.aws.amazon.com/.
