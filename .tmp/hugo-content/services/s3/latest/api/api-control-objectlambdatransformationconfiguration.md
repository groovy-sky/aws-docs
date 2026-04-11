# ObjectLambdaTransformationConfiguration

A configuration used when creating an Object Lambda Access Point transformation.

## Contents

**Actions**

A container for the action of an Object Lambda Access Point configuration. Valid inputs are
`GetObject`, `ListObjects`, `HeadObject`, and
`ListObjectsV2`.

Type: Array of strings

Valid Values: `GetObject | HeadObject | ListObjects | ListObjectsV2`

Required: Yes

**ContentTransformation**

A container for the content transformation of an Object Lambda Access Point configuration.

Type: [ObjectLambdaContentTransformation](api-control-objectlambdacontenttransformation.md) data type

**Note:** This object is a Union. Only one member of this object can be specified or returned.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/objectlambdatransformationconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/objectlambdatransformationconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/objectlambdatransformationconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ObjectLambdaContentTransformation

PolicyStatus

All content copied from https://docs.aws.amazon.com/.
