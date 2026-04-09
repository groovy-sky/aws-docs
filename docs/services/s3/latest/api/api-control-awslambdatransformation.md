# AwsLambdaTransformation

AWS Lambda function used to transform objects through an Object Lambda Access Point.

## Contents

**FunctionArn**

The Amazon Resource Name (ARN) of the AWS Lambda function.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `(arn:(aws[a-zA-Z-]*)?:lambda:)?([a-z]{2}((-gov)|(-iso([a-z]?)))?-[a-z]+-\d{1}:)?(\d{12}:)?(function:)?([a-zA-Z0-9-_]+)(:(\$LATEST|[a-zA-Z0-9-_]+))?`

Required: Yes

**FunctionPayload**

Additional JSON that provides supplemental data to the Lambda function used to transform
objects.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/awslambdatransformation.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/awslambdatransformation.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/awslambdatransformation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AsyncResponseDetails

BucketLevel

All content copied from https://docs.aws.amazon.com/.
