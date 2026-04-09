# VersioningConfiguration

Describes the versioning state of an Amazon S3 bucket. For more information, see [PUT Bucket\
versioning](restbucketputversioningstatus.md) in the _Amazon S3 API Reference_.

## Contents

**MFADelete**

Specifies whether MFA delete is enabled in the bucket versioning configuration. This element is only
returned if the bucket has been configured with MFA delete. If the bucket has never been so configured,
this element is not returned.

Type: String

Valid Values: `Enabled | Disabled`

Required: No

**Status**

The versioning state of the bucket.

Type: String

Valid Values: `Enabled | Suspended`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/versioningconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/versioningconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/versioningconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Transition

WebsiteConfiguration

All content copied from https://docs.aws.amazon.com/.
