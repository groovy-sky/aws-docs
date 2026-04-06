# PutTableBucketReplication

Creates or updates the replication configuration for a table bucket. This operation defines how tables in the source bucket are replicated to destination buckets. Replication helps ensure data availability and disaster recovery across regions or accounts.

Permissions

- You must have the `s3tables:PutTableBucketReplication` permission to use this operation. The IAM role specified in the configuration must have permissions to read from the source bucket and write permissions to all destination buckets.

- You must also have the following permissions:

- `s3tables:GetTable` permission on the source table.

- `s3tables:ListTables` permission on the bucket containing the table.

- `s3tables:CreateTable` permission for the destination.

- `s3tables:CreateNamespace` permission for the destination.

- `s3tables:GetTableMaintenanceConfig` permission for the source
bucket.

- `s3tables:PutTableMaintenanceConfig` permission for the destination
bucket.

- You must have `iam:PassRole` permission with condition allowing
roles to be passed to `replication.s3tables.amazonaws.com`.

## Request Syntax

```nohighlight

PUT /table-bucket-replication?tableBucketARN=tableBucketARN&versionToken=versionToken HTTP/1.1
Content-type: application/json

{
   "configuration": {
      "role": "string",
      "rules": [
         {
            "destinations": [
               {
                  "destinationTableBucketARN": "string"
               }
            ]
         }
      ]
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[tableBucketARN](#API_s3Buckets_PutTableBucketReplication_RequestSyntax)**

The Amazon Resource Name (ARN) of the source table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

**[versionToken](#API_s3Buckets_PutTableBucketReplication_RequestSyntax)**

A version token from a previous GetTableBucketReplication call. Use this token to ensure you're updating the expected version of the configuration.

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Request Body

The request accepts the following data in JSON format.

**[configuration](#API_s3Buckets_PutTableBucketReplication_RequestSyntax)**

The replication configuration to apply, including the IAM role and replication rules.

Type: [TableBucketReplicationConfiguration](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_TableBucketReplicationConfiguration.html) object

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "status": "string",
   "versionToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[status](#API_s3Buckets_PutTableBucketReplication_ResponseSyntax)**

The status of the replication configuration operation.

Type: String

**[versionToken](#API_s3Buckets_PutTableBucketReplication_ResponseSyntax)**

A new version token representing the updated replication configuration.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Errors

**AccessDeniedException**

The action cannot be performed because you do not have the required permission.

HTTP Status Code: 403

**BadRequestException**

The request is invalid or malformed.

HTTP Status Code: 400

**ConflictException**

The request failed because there is a conflict with a previous write. You can retry the
request.

HTTP Status Code: 409

**ForbiddenException**

The caller isn't authorized to make the request.

HTTP Status Code: 403

**InternalServerErrorException**

The request failed due to an internal server error.

HTTP Status Code: 500

**NotFoundException**

The request was rejected because the specified resource could not be found.

HTTP Status Code: 404

**TooManyRequestsException**

The limit on the number of requests per second was exceeded.

HTTP Status Code: 429

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3tables-2018-05-10/PutTableBucketReplication)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/PutTableBucketReplication)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutTableBucketPolicy

PutTableBucketStorageClass
