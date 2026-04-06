# S3PublicAccessBlockConfiguration

The `PublicAccessBlock` configuration to apply to this Amazon S3 bucket. If the
proposed configuration is for an existing Amazon S3 bucket and the configuration is not
specified, the access preview uses the existing setting. If the proposed configuration is
for a new bucket and the configuration is not specified, the access preview uses
`false`. If the proposed configuration is for a new access point or
multi-region access point and the access point BPA configuration is not specified, the
access preview uses `true`. For more information, see [PublicAccessBlockConfiguration](../../../../services/cloudformation/latest/userguide/aws-properties-s3-bucket-publicaccessblockconfiguration.md).

## Contents

**ignorePublicAcls**

Specifies whether Amazon S3 should ignore public ACLs for this bucket and objects in this
bucket.

Type: Boolean

Required: Yes

**restrictPublicBuckets**

Specifies whether Amazon S3 should restrict public bucket policies for this bucket.

Type: Boolean

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/s3publicaccessblockconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/s3publicaccessblockconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/s3publicaccessblockconfiguration.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

S3ExpressDirectoryBucketConfiguration

SecretsManagerSecretConfiguration
