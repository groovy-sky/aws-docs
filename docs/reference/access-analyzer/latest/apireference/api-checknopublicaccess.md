# CheckNoPublicAccess

Checks whether a resource policy can grant public access to the specified resource
type.

## Request Syntax

```nohighlight

POST /policy/check-no-public-access HTTP/1.1
Content-type: application/json

{
   "policyDocument": "string",
   "resourceType": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[policyDocument](#API_CheckNoPublicAccess_RequestSyntax)**

The JSON policy document to evaluate for public access.

Type: String

Required: Yes

**[resourceType](#API_CheckNoPublicAccess_RequestSyntax)**

The type of resource to evaluate for public access. For example, to check for public
access to Amazon S3 buckets, you can choose `AWS::S3::Bucket` for the resource
type.

For resource types not supported as valid values, IAM Access Analyzer will return an
error.

Type: String

Valid Values: `AWS::DynamoDB::Table | AWS::DynamoDB::Stream | AWS::EFS::FileSystem | AWS::OpenSearchService::Domain | AWS::Kinesis::Stream | AWS::Kinesis::StreamConsumer | AWS::KMS::Key | AWS::Lambda::Function | AWS::S3::Bucket | AWS::S3::AccessPoint | AWS::S3Express::DirectoryBucket | AWS::S3::Glacier | AWS::S3Outposts::Bucket | AWS::S3Outposts::AccessPoint | AWS::SecretsManager::Secret | AWS::SNS::Topic | AWS::SQS::Queue | AWS::IAM::AssumeRolePolicyDocument | AWS::S3Tables::TableBucket | AWS::ApiGateway::RestApi | AWS::CodeArtifact::Domain | AWS::Backup::BackupVault | AWS::CloudTrail::Dashboard | AWS::CloudTrail::EventDataStore | AWS::S3Tables::Table | AWS::S3Express::AccessPoint`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "message": "string",
   "reasons": [
      {
         "description": "string",
         "statementId": "string",
         "statementIndex": number
      }
   ],
   "result": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[message](#API_CheckNoPublicAccess_ResponseSyntax)**

The message indicating whether the specified policy allows public access to
resources.

Type: String

**[reasons](#API_CheckNoPublicAccess_ResponseSyntax)**

A list of reasons why the specified resource policy grants public access for the
resource type.

Type: Array of [ReasonSummary](api-reasonsummary.md) objects

**[result](#API_CheckNoPublicAccess_ResponseSyntax)**

The result of the check for public access to the specified resource type. If the result
is `PASS`, the policy doesn't allow public access to the specified resource
type. If the result is `FAIL`, the policy might allow public access to the
specified resource type.

Type: String

Valid Values: `PASS | FAIL`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You do not have sufficient access to perform this action.

HTTP Status Code: 403

**InternalServerException**

Internal server error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 500

**InvalidParameterException**

The specified parameter is invalid.

HTTP Status Code: 400

**ThrottlingException**

Throttling limit exceeded error.

**retryAfterSeconds**

The seconds to wait to retry.

HTTP Status Code: 429

**UnprocessableEntityException**

The specified entity could not be processed.

HTTP Status Code: 422

**ValidationException**

Validation exception error.

**fieldList**

A list of fields that didn't validate.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/accessanalyzer-2019-11-01/CheckNoPublicAccess)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/accessanalyzer-2019-11-01/CheckNoPublicAccess)

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/accessanalyzer-2019-11-01/checknopublicaccess.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/accessanalyzer-2019-11-01/CheckNoPublicAccess)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/checknopublicaccess.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CheckNoNewAccess

CreateAccessPreview
