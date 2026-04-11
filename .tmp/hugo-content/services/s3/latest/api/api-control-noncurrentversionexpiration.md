# NoncurrentVersionExpiration

The container of the noncurrent version expiration.

## Contents

**NewerNoncurrentVersions**

Specifies how many noncurrent versions S3 on Outposts will retain. If there are this many
more recent noncurrent versions, S3 on Outposts will take the associated action. For more
information about noncurrent versions, see [Lifecycle configuration\
elements](../userguide/intro-lifecycle-rules.md) in the _Amazon S3 User Guide_.

Type: Integer

Required: No

**NoncurrentDays**

Specifies the number of days an object is noncurrent before Amazon S3 can perform the
associated action. For information about the noncurrent days calculations, see [How\
Amazon S3 Calculates When an Object Became Noncurrent](../dev/intro-lifecycle-rules.md#non-current-days-calculations) in the
_Amazon S3 User Guide_.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/noncurrentversionexpiration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/noncurrentversionexpiration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/noncurrentversionexpiration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MultiRegionAccessPointsAsyncResponse

NoncurrentVersionTransition

All content copied from https://docs.aws.amazon.com/.
