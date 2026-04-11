# LifecycleRule

The container for the Outposts bucket lifecycle rule.

## Contents

**Status**

If 'Enabled', the rule is currently being applied. If 'Disabled', the rule is not
currently being applied.

Type: String

Valid Values: `Enabled | Disabled`

Required: Yes

**AbortIncompleteMultipartUpload**

Specifies the days since the initiation of an incomplete multipart upload that Amazon S3
waits before permanently removing all parts of the upload. For more information, see [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration](../dev/mpuoverview.md#mpu-abort-incomplete-mpu-lifecycle-config) in
the _Amazon S3 User Guide_.

Type: [AbortIncompleteMultipartUpload](api-control-abortincompletemultipartupload.md) data type

Required: No

**Expiration**

Specifies the expiration for the lifecycle of the object in the form of date, days and,
whether the object has a delete marker.

Type: [LifecycleExpiration](api-control-lifecycleexpiration.md) data type

Required: No

**Filter**

The container for the filter of lifecycle rule.

Type: [LifecycleRuleFilter](api-control-lifecyclerulefilter.md) data type

Required: No

**ID**

Unique identifier for the rule. The value cannot be longer than 255 characters.

Type: String

Required: No

**NoncurrentVersionExpiration**

The noncurrent version expiration of the lifecycle rule.

Type: [NoncurrentVersionExpiration](api-control-noncurrentversionexpiration.md) data type

Required: No

**NoncurrentVersionTransitions**

Specifies the transition rule for the lifecycle rule that describes when noncurrent
objects transition to a specific storage class. If your bucket is versioning-enabled (or
versioning is suspended), you can set this action to request that Amazon S3 transition
noncurrent object versions to a specific storage class at a set period in the object's
lifetime.

###### Note

This is not supported by Amazon S3 on Outposts buckets.

Type: Array of [NoncurrentVersionTransition](api-control-noncurrentversiontransition.md) data types

Required: No

**Transitions**

Specifies when an Amazon S3 object transitions to a specified storage class.

###### Note

This is not supported by Amazon S3 on Outposts buckets.

Type: Array of [Transition](api-control-transition.md) data types

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/lifecyclerule.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/lifecyclerule.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/lifecyclerule.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleExpiration

LifecycleRuleAndOperator

All content copied from https://docs.aws.amazon.com/.
