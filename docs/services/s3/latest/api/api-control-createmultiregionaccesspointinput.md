# CreateMultiRegionAccessPointInput

A container for the information associated with a [CreateMultiRegionAccessPoint](api-control-createmultiregionaccesspoint.md) request.

## Contents

**Name**

The name of the Multi-Region Access Point associated with this request.

Type: String

Length Constraints: Maximum length of 50.

Pattern: `^[a-z0-9][-a-z0-9]{1,48}[a-z0-9]$`

Required: Yes

**Regions**

The buckets in different Regions that are associated with the Multi-Region Access Point.

Type: Array of [Region](api-control-region.md) data types

Required: Yes

**PublicAccessBlock**

The `PublicAccessBlock` configuration that you want to apply to this Amazon S3
account. You can enable the configuration options in any combination. For more information
about when Amazon S3 considers a bucket or object public, see [The Meaning of "Public"](../dev/access-control-block-public-access.md#access-control-block-public-access-policy-status) in the _Amazon S3 User Guide_.

This data type is not supported for Amazon S3 on Outposts.

Type: [PublicAccessBlockConfiguration](api-control-publicaccessblockconfiguration.md) data type

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/createmultiregionaccesspointinput.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/createmultiregionaccesspointinput.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/createmultiregionaccesspointinput.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateBucketConfiguration

Credentials

All content copied from https://docs.aws.amazon.com/.
