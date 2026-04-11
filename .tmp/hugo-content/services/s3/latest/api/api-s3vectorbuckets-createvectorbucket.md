# CreateVectorBucket

Creates a vector bucket in the AWS Region that you want your bucket to be in.

Permissions

You must have the `s3vectors:CreateVectorBucket` permission to use
this operation.

You must have the `s3vectors:TagResource` permission in addition to
`s3vectors:CreateVectorBucket` permission to create a vector bucket
with tags.

## Request Syntax

```nohighlight

POST /CreateVectorBucket HTTP/1.1
Content-type: application/json

{
   "encryptionConfiguration": {
      "kmsKeyArn": "string",
      "sseType": "string"
   },
   "tags": {
      "string" : "string"
   },
   "vectorBucketName": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[encryptionConfiguration](#API_S3VectorBuckets_CreateVectorBucket_RequestSyntax)**

The encryption configuration for the vector bucket. By default, if you don't specify,
all new vectors in Amazon S3 vector buckets use server-side encryption with Amazon S3
managed keys (SSE-S3), specifically `AES256`.

Type: [EncryptionConfiguration](api-s3vectorbuckets-encryptionconfiguration.md) object

Required: No

**[tags](#API_S3VectorBuckets_CreateVectorBucket_RequestSyntax)**

An array of user-defined tags that you would like to apply to the vector bucket that you
are creating. A tag is a key-value pair that you apply to your resources. Tags can help you
organize and control access to resources. For more information, see [Tagging for cost\
allocation or attribute-based access control (ABAC)](../userguide/tagging.md).

###### Note

You must have the `s3vectors:TagResource` permission in addition to
`s3vectors:CreateVectorBucket` permission to create a vector bucket with
tags.

Type: String to string map

Key Length Constraints: Minimum length of 1. Maximum length of 128.

Key Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Value Length Constraints: Minimum length of 0. Maximum length of 256.

Value Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

Required: No

**[vectorBucketName](#API_S3VectorBuckets_CreateVectorBucket_RequestSyntax)**

The name of the vector bucket to create.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "vectorBucketArn": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[vectorBucketArn](#API_S3VectorBuckets_CreateVectorBucket_ResponseSyntax)**

The Amazon Resource Name (ARN) of the newly created vector bucket.

Type: String

Pattern: `arn:aws[-a-z0-9]*:s3vectors:[a-z0-9-]+:[0-9]{12}:bucket/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]`

## Errors

**AccessDeniedException**

Access denied.

HTTP Status Code: 403

**ConflictException**

The request failed because a vector bucket name or a vector index name already exists.
Vector bucket names must be unique within your AWS account for each AWS Region. Vector
index names must be unique within your vector bucket. Choose a different vector bucket name
or vector index name, and try again.

HTTP Status Code: 409

**InternalServerException**

The request failed due to an internal server error.

HTTP Status Code: 500

**RequestTimeoutException**

The request timed out. Retry your request.

HTTP Status Code: 408

**ServiceQuotaExceededException**

Your request exceeds a service quota.

HTTP Status Code: 402

**ServiceUnavailableException**

The service is unavailable. Wait briefly and retry your request. If it continues to
fail, increase your waiting time between retries.

HTTP Status Code: 503

**TooManyRequestsException**

The request was denied due to request throttling.

HTTP Status Code: 429

**ValidationException**

The requested action isn't valid.

**fieldList**

A list of specific validation failures that are encountered during input processing. Each entry
in the list contains a path to the field that failed validation and a detailed message that
explains why the validation failed. To satisfy multiple constraints, a field can appear multiple times in this list if it
failed. You can use the information to identify and fix
the specific validation issues in your request.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for Python](../../../goto/boto3/s3vectors-2025-07-15/createvectorbucket.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/createvectorbucket.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateIndex

DeleteIndex

All content copied from https://docs.aws.amazon.com/.
