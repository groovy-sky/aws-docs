# S3AccessPointConfiguration

The configuration for an Amazon S3 access point or multi-region access point for the bucket.
You can propose up to 10 access points or multi-region access points per bucket. If the
proposed Amazon S3 access point configuration is for an existing bucket, the access preview uses
the proposed access point configuration in place of the existing access points. To propose
an access point without a policy, you can provide an empty string as the access point
policy. For more information, see [Creating access points](https://docs.aws.amazon.com/AmazonS3/latest/dev/creating-access-points.html).
For more information about access point policy limits, see [Access points\
restrictions and limitations](https://docs.aws.amazon.com/AmazonS3/latest/dev/access-points-restrictions-limitations.html).

## Contents

**accessPointPolicy**

The access point or multi-region access point policy.

Type: String

Required: No

**networkOrigin**

The proposed `Internet` and `VpcConfiguration` to apply to this
Amazon S3 access point. `VpcConfiguration` does not apply to multi-region access
points. If the access preview is for a new resource and neither is specified, the access
preview uses `Internet` for the network origin. If the access preview is for an
existing resource and neither is specified, the access preview uses the existing network
origin.

Type: [NetworkOriginConfiguration](api-networkoriginconfiguration.md) object

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: No

**publicAccessBlock**

The proposed `S3PublicAccessBlock` configuration to apply to this Amazon S3 access
point or multi-region access point.

Type: [S3PublicAccessBlockConfiguration](api-s3publicaccessblockconfiguration.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/accessanalyzer-2019-11-01/S3AccessPointConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/accessanalyzer-2019-11-01/S3AccessPointConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/accessanalyzer-2019-11-01/S3AccessPointConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResourceTypeDetails

S3BucketAclGrantConfiguration
