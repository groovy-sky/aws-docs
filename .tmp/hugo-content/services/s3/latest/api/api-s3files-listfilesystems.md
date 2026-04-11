# ListFileSystems

Returns a list of all S3 File Systems owned by the account with optional filtering by
bucket.

## Request Syntax

```nohighlight

GET /file-systems?bucket=bucket&maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[bucket](#API_S3Files_ListFileSystems_RequestSyntax)**

Optional filter to list only file systems associated with the specified S3 bucket
Amazon Resource Name (ARN). If provided, only file systems that provide access to this bucket will be
returned in the response.

Pattern: `(arn:aws[a-zA-Z0-9-]*:s3:::.+)`

**[maxResults](#API_S3Files_ListFileSystems_RequestSyntax)**

The maximum number of file systems to return in a single response. If not specified,
up to 100 file systems are returned.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_S3Files_ListFileSystems_RequestSyntax)**

A pagination token returned from a previous call to continue listing file
systems.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "fileSystems": [
      {
         "bucket": "string",
         "creationTime": number,
         "fileSystemArn": "string",
         "fileSystemId": "string",
         "name": "string",
         "ownerId": "string",
         "roleArn": "string",
         "status": "string",
         "statusMessage": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[fileSystems](#API_S3Files_ListFileSystems_ResponseSyntax)**

An array of file system descriptions.

Type: Array of [ListFileSystemsDescription](api-s3files-listfilesystemsdescription.md) objects

**[nextToken](#API_S3Files_ListFileSystems_ResponseSyntax)**

A pagination token to use in a subsequent request if more results are
available.

Type: String

## Errors

**InternalServerException**

An internal server error occurred. Retry your request.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 500

**ThrottlingException**

The request was throttled. Retry your request using exponential backoff.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 429

**ValidationException**

The input parameters are not valid. Check the parameter values and try again.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/listfilesystems.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listfilesystems.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListAccessPoints

ListMountTargets

All content copied from https://docs.aws.amazon.com/.
