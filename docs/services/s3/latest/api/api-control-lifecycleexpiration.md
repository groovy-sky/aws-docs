# LifecycleExpiration

The container of the Outposts bucket lifecycle expiration.

## Contents

**Date**

Indicates at what date the object is to be deleted. Should be in GMT ISO 8601
format.

Type: Timestamp

Required: No

**Days**

Indicates the lifetime, in days, of the objects that are subject to the rule. The value
must be a non-zero positive integer.

Type: Integer

Required: No

**ExpiredObjectDeleteMarker**

Indicates whether Amazon S3 will remove a delete marker with no noncurrent versions. If set
to true, the delete marker will be expired. If set to false, the policy takes no action.
This cannot be specified with Days or Date in a Lifecycle Expiration Policy. To learn more about delete markers, see [Working with delete markers](../userguide/deletemarker.md).

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/lifecycleexpiration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/lifecycleexpiration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/lifecycleexpiration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LifecycleConfiguration

LifecycleRule

All content copied from https://docs.aws.amazon.com/.
