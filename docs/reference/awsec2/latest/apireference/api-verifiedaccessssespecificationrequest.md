# VerifiedAccessSseSpecificationRequest

AWS Verified Access provides server side encryption by default to data at rest using AWS-owned KMS keys. You also have the option of using customer managed KMS keys, which can be specified using the options below.

## Contents

**CustomerManagedKeyEnabled**

Enable or disable the use of customer managed KMS keys for server side encryption.

Valid values: `True` \| `False`

Type: Boolean

Required: No

**KmsKeyArn**

The ARN of the KMS key.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/VerifiedAccessSseSpecificationRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/VerifiedAccessSseSpecificationRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/VerifiedAccessSseSpecificationRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VerifiedAccessLogS3DestinationOptions

VerifiedAccessSseSpecificationResponse
