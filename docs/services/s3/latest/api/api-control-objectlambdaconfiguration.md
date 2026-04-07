# ObjectLambdaConfiguration

A configuration used when creating an Object Lambda Access Point.

## Contents

**SupportingAccessPoint**

Standard access point associated with the Object Lambda Access Point.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `arn:[^:]+:s3:[^:]*:\d{12}:accesspoint/.*`

Required: Yes

**TransformationConfigurations**

A container for transformation configurations for an Object Lambda Access Point.

Type: Array of [ObjectLambdaTransformationConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ObjectLambdaTransformationConfiguration.html) data types

Required: Yes

**AllowedFeatures**

A container for allowed features. Valid inputs are `GetObject-Range`,
`GetObject-PartNumber`, `HeadObject-Range`, and
`HeadObject-PartNumber`.

Type: Array of strings

Valid Values: `GetObject-Range | GetObject-PartNumber | HeadObject-Range | HeadObject-PartNumber`

Required: No

**CloudWatchMetricsEnabled**

A container for whether the CloudWatch metrics configuration is enabled.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/ObjectLambdaConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/ObjectLambdaConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/ObjectLambdaConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ObjectLambdaAccessPointAlias

ObjectLambdaContentTransformation
