# PutTableBucketPolicy

Creates a new table bucket policy or replaces an existing table bucket policy for a
table bucket. For more information, see [Adding a table bucket policy](../userguide/s3-tables-bucket-policy.md#table-bucket-policy-add) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:PutTableBucketPolicy` permission to use this operation.

## Request Syntax

```nohighlight

PUT /buckets/tableBucketARN/policy HTTP/1.1
Content-type: application/json

{
   "resourcePolicy": "string"
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[tableBucketARN](#API_s3Buckets_PutTableBucketPolicy_RequestSyntax)**

The Amazon Resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[resourcePolicy](#API_s3Buckets_PutTableBucketPolicy_RequestSyntax)**

The `JSON` that defines the policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 20480.

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for Python](../../../goto/boto3/s3tables-2018-05-10/puttablebucketpolicy.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3tables-2018-05-10/puttablebucketpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutTableBucketMetricsConfiguration

PutTableBucketReplication

All content copied from https://docs.aws.amazon.com/.
