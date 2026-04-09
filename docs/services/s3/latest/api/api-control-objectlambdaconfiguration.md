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

Type: Array of [ObjectLambdaTransformationConfiguration](api-control-objectlambdatransformationconfiguration.md) data types

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

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/objectlambdaconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/objectlambdaconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/objectlambdaconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectLambdaAccessPointAlias

ObjectLambdaContentTransformation

All content copied from https://docs.aws.amazon.com/.
