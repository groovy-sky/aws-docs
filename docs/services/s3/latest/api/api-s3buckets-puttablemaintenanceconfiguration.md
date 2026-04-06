# PutTableMaintenanceConfiguration

Creates a new maintenance configuration or replaces an existing maintenance configuration
for a table. For more information, see [S3 Tables maintenance](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance.html) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:PutTableMaintenanceConfiguration` permission to use this operation.

## Request Syntax

```nohighlight

PUT /tables/tableBucketARN/namespace/name/maintenance/type HTTP/1.1
Content-type: application/json

{
   "value": {
      "settings": { ... },
      "status": "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_s3Buckets_PutTableMaintenanceConfiguration_RequestSyntax)**

The name of the table.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**[namespace](#API_s3Buckets_PutTableMaintenanceConfiguration_RequestSyntax)**

The namespace of the table.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

Required: Yes

**[tableBucketARN](#API_s3Buckets_PutTableMaintenanceConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the table associated with the maintenance
configuration.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

**[type](#API_s3Buckets_PutTableMaintenanceConfiguration_RequestSyntax)**

The type of the maintenance configuration.

Valid Values: `icebergCompaction | icebergSnapshotManagement`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[value](#API_s3Buckets_PutTableMaintenanceConfiguration_RequestSyntax)**

Defines the values of the maintenance configuration for the table.

Type: [TableMaintenanceConfigurationValue](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_TableMaintenanceConfigurationValue.html) object

Required: Yes

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## Errors

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/PutTableMaintenanceConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutTableBucketStorageClass

PutTablePolicy
