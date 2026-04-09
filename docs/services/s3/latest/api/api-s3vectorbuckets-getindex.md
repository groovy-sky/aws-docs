# GetIndex

Returns vector index attributes. To specify the vector index, you can either use both
the vector bucket name and the vector index name, or use the vector index Amazon Resource
Name (ARN).

Permissions

You must have the `s3vectors:GetIndex` permission to use this
operation.

## Request Syntax

```nohighlight

POST /GetIndex HTTP/1.1
Content-type: application/json

{
   "indexArn": "string",
   "indexName": "string",
   "vectorBucketName": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[indexArn](#API_S3VectorBuckets_GetIndex_RequestSyntax)**

The ARN of the vector index.

Type: String

Pattern: `arn:aws[-a-z0-9]*:s3vectors:[a-z0-9-]+:[0-9]{12}:bucket/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]/index/[a-z0-9][a-z0-9-.]{1,61}[a-z0-9]`

Required: No

**[indexName](#API_S3VectorBuckets_GetIndex_RequestSyntax)**

The name of the vector index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: No

**[vectorBucketName](#API_S3VectorBuckets_GetIndex_RequestSyntax)**

The name of the vector bucket that contains the vector index.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 63.

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "index": {
      "creationTime": number,
      "dataType": "string",
      "dimension": number,
      "distanceMetric": "string",
      "encryptionConfiguration": {
         "kmsKeyArn": "string",
         "sseType": "string"
      },
      "indexArn": "string",
      "indexName": "string",
      "metadataConfiguration": {
         "nonFilterableMetadataKeys": [ "string" ]
      },
      "vectorBucketName": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[index](#API_S3VectorBuckets_GetIndex_ResponseSyntax)**

The attributes of the vector index.

Type: [Index](api-s3vectorbuckets-index.md) object

## Errors

**AccessDeniedException**

Access denied.

HTTP Status Code: 403

**InternalServerException**

The request failed due to an internal server error.

HTTP Status Code: 500

**NotFoundException**

The request was rejected because the specified resource can't be found.

HTTP Status Code: 404

**RequestTimeoutException**

The request timed out. Retry your request.

HTTP Status Code: 408

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for Python](../../../goto/boto3/s3vectors-2025-07-15/getindex.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3vectors-2025-07-15/getindex.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteVectors

GetVectorBucket

All content copied from https://docs.aws.amazon.com/.
