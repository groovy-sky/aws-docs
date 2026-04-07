# ObjectLambdaAccessPoint

An access point with an attached AWS Lambda function used to access transformed data from an Amazon S3
bucket.

## Contents

**Name**

The name of the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 45.

Pattern: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

Required: Yes

**Alias**

The alias of the Object Lambda Access Point.

Type: [ObjectLambdaAccessPointAlias](api-control-objectlambdaaccesspointalias.md) data type

Required: No

**ObjectLambdaAccessPointArn**

Specifies the ARN for the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[^:]+:s3-object-lambda:[^:]*:\d{12}:accesspoint/.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/ObjectLambdaAccessPoint)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/ObjectLambdaAccessPoint)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/ObjectLambdaAccessPoint)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ObjectEncryptionFilter

ObjectLambdaAccessPointAlias
