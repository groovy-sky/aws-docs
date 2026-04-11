# LoggingEnabled

Describes where logs are stored and the prefix that Amazon S3 assigns to all log object keys for a
bucket. For more information, see [PUT Bucket logging](restbucketputlogging.md) in the
_Amazon S3 API Reference_.

## Contents

**TargetBucket**

Specifies the bucket where you want Amazon S3 to store server access logs. You can have your logs
delivered to any bucket that you own, including the same bucket that is being logged. You can also
configure multiple buckets to deliver their logs to the same target bucket. In this case, you should
choose a different `TargetPrefix` for each source bucket so that the delivered log files can
be distinguished by key.

Type: String

Required: Yes

**TargetPrefix**

A prefix for all log object keys. If you store log files from multiple Amazon S3 buckets in a single
bucket, you can use a prefix to distinguish which log files came from which bucket.

Type: String

Required: Yes

**TargetGrants**

Container for granting information.

Buckets that use the bucket owner enforced setting for Object Ownership don't support target grants.
For more information, see [Permissions for server access log delivery](../userguide/enable-server-access-logging.md#grant-log-delivery-permissions-general) in the
_Amazon S3 User Guide_.

Type: Array of [TargetGrant](api-targetgrant.md) data types

Required: No

**TargetObjectKeyFormat**

Amazon S3 key format for log objects.

Type: [TargetObjectKeyFormat](api-targetobjectkeyformat.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/loggingenabled.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/loggingenabled.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/loggingenabled.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LocationInfo

MetadataConfiguration

All content copied from https://docs.aws.amazon.com/.
