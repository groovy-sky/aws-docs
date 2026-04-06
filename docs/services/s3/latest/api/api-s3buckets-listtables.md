# ListTables

List tables in the given table bucket. For more information, see [S3 Tables](../userguide/s3-tables-tables.md) in the _Amazon Simple Storage Service User Guide_.

Permissions

You must have the `s3tables:ListTables` permission to use this operation.

## Request Syntax

```nohighlight

GET /tables/tableBucketARN?continuationToken=continuationToken&maxTables=maxTables&namespace=namespace&prefix=prefix HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[continuationToken](#API_s3Buckets_ListTables_RequestSyntax)**

`ContinuationToken` indicates to Amazon S3 that the list is being continued on
this bucket with a token. `ContinuationToken` is obfuscated and is not a real key.
You can use this `ContinuationToken` for pagination of the list results.

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[maxTables](#API_s3Buckets_ListTables_RequestSyntax)**

The maximum number of tables to return.

Valid Range: Minimum value of 1. Maximum value of 1000.

**[namespace](#API_s3Buckets_ListTables_RequestSyntax)**

The namespace of the tables.

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `[0-9a-z_]*`

**[prefix](#API_s3Buckets_ListTables_RequestSyntax)**

The prefix of the tables.

Length Constraints: Minimum length of 1. Maximum length of 255.

**[tableBucketARN](#API_s3Buckets_ListTables_RequestSyntax)**

The Amazon resource Name (ARN) of the table bucket.

Pattern: `(arn:aws[-a-z0-9]*:[a-z0-9]+:[-a-z0-9]*:[0-9]{12}:bucket/[a-z0-9_-]{3,63})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "continuationToken": "string",
   "tables": [
      {
         "createdAt": "string",
         "managedByService": "string",
         "modifiedAt": "string",
         "name": "string",
         "namespace": [ "string" ],
         "namespaceId": "string",
         "tableARN": "string",
         "tableBucketId": "string",
         "type": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[continuationToken](#API_s3Buckets_ListTables_ResponseSyntax)**

You can use this `ContinuationToken` for pagination of the list results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[tables](#API_s3Buckets_ListTables_ResponseSyntax)**

A list of tables.

Type: Array of [TableSummary](https://docs.aws.amazon.com/AmazonS3/latest/API/API_s3Buckets_TableSummary.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3tables-2018-05-10/ListTables)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3tables-2018-05-10/ListTables)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3tables-2018-05-10/ListTables)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3tables-2018-05-10/ListTables)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3tables-2018-05-10/ListTables)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3tables-2018-05-10/ListTables)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3tables-2018-05-10/ListTables)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3tables-2018-05-10/ListTables)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3tables-2018-05-10/ListTables)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3tables-2018-05-10/ListTables)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ListTableBuckets

ListTagsForResource
