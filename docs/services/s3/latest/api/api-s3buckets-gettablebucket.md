# GetTableBucket

Gets details on a table bucket. For more information, see [Viewing details about an Amazon S3 table bucket](../userguide/s3-tables-buckets-details.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:GetTableBucket` permission to use this operation.

## Request Syntax

```nohighlight

GET /buckets/tableBucketARN HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[tableBucketARN](#API_s3Buckets_GetTableBucket_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "arn": "string",
   "createdAt": "string",
   "name": "string",
   "ownerAccountId": "string",
   "tableBucketId": "string",
   "type": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[arn](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Type: String

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

**[createdAt](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The date and time the table bucket was created.

Type: Timestamp

**[name](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The name of the table bucket

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Pattern: `[0-9a-z-]*`

**[ownerAccountId](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The ID of the account that owns the table bucket.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `[0-9].*`

**[tableBucketId](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The unique identifier of the table bucket.

Type: String

**[type](#API_s3Buckets_GetTableBucket_ResponseSyntax)**

The type of the table bucket.

Type: String

Valid Values: `customer | aws`

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/gettablebucket.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/gettablebucket.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetTable

GetTableBucketEncryption

All content copied from https://docs.aws.amazon.com/.
