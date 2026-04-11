# StorageLensAwsOrg

The AWS organization for your S3 Storage Lens.

## Contents

**Arn**

A container for the Amazon Resource Name (ARN) of the AWS organization. This property
is read-only and follows the following format: `
               arn:aws:organizations:us-east-1:example-account-id:organization/o-ex2l495dck
                  `

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `arn:[a-z\-]+:organizations::\d{12}:organization\/o-[a-z0-9]{10,32}`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/storagelensawsorg.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/storagelensawsorg.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/storagelensawsorg.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SSES3Filter

StorageLensConfiguration

All content copied from https://docs.aws.amazon.com/.
