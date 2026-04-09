# GetFileSystemPolicy

Returns the IAM resource policy of an S3 File System.

## Request Syntax

```nohighlight

GET /file-systems/fileSystemId/policy HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[fileSystemId](#API_S3Files_GetFileSystemPolicy_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the S3 File System whose resource policy to retrieve.

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "fileSystemId": "string",
   "policy": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[fileSystemId](#API_S3Files_GetFileSystemPolicy_ResponseSyntax)**

The ID of the file system.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[policy](#API_S3Files_GetFileSystemPolicy_ResponseSyntax)**

The JSON-formatted resource policy for the file system.

Type: String

## Errors

**InternalServerException**

An internal server error occurred. Retry your request.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 500

**ResourceNotFoundException**

The specified resource was not found. Verify that the resource exists and that you have permission to access it.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 404

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/getfilesystempolicy.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/getfilesystempolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetFileSystem

GetMountTarget

All content copied from https://docs.aws.amazon.com/.
