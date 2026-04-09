# CalculationResult

Contains information about an application-specific calculation result.

## Contents

**ResultS3Uri**

The Amazon S3 location of the folder for the calculation results.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `^(https|s3|S3)://([^/]+)/?(.*)$`

Required: No

**ResultType**

The data format of the calculation result.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `\w+\/[-+.\w]+`

Required: No

**StdErrorS3Uri**

The Amazon S3 location of the `stderr` error messages file for the
calculation.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `^(https|s3|S3)://([^/]+)/?(.*)$`

Required: No

**StdOutS3Uri**

The Amazon S3 location of the `stdout` file for the calculation.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `^(https|s3|S3)://([^/]+)/?(.*)$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/calculationresult.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/calculationresult.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/calculationresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CalculationConfiguration

CalculationStatistics

All content copied from https://docs.aws.amazon.com/.
