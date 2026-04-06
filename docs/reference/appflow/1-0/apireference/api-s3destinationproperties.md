# S3DestinationProperties

The properties that are applied when Amazon S3 is used as a destination.

## Contents

**bucketName**

The Amazon S3 bucket name in which Amazon AppFlow places the transferred
data.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `\S+`

Required: Yes

**bucketPrefix**

The object key for the destination bucket in which Amazon AppFlow places the files.

Type: String

Length Constraints: Maximum length of 512.

Pattern: `.*`

Required: No

**s3OutputFormatConfig**

The configuration that determines how Amazon AppFlow should format the flow output
data when Amazon S3 is used as the destination.

Type: [S3OutputFormatConfig](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_S3OutputFormatConfig.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/S3DestinationProperties)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/S3DestinationProperties)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/S3DestinationProperties)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RegistrationOutput

S3InputFormatConfig
