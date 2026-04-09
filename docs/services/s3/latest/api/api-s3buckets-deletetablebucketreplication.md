# DeleteTableBucketReplication

Deletes the replication configuration for a table bucket. After deletion, new table updates will no longer be replicated to destination buckets, though existing replicated tables will remain in destination buckets.

Permissions

You must have the `s3tables:DeleteTableBucketReplication` permission to use this operation.

## Request Syntax

```nohighlight

DELETE /table-bucket-replication?tableBucketARN=tableBucketARN&versionToken=versionToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[tableBucketARN](#API_s3Buckets_DeleteTableBucketReplication_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

**[versionToken](#API_s3Buckets_DeleteTableBucketReplication_RequestSyntax)**

A version token from a previous GetTableBucketReplication call. Use this token to ensure you're deleting the expected version of the configuration.

Length Constraints: Minimum length of 1. Maximum length of 2048.

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/deletetablebucketreplication.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/deletetablebucketreplication.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteTableBucketPolicy

DeleteTablePolicy

All content copied from https://docs.aws.amazon.com/.
