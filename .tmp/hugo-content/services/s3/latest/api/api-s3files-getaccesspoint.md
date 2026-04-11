# GetAccessPoint

Returns resource information for an S3 File System Access Point.

## Request Syntax

```nohighlight

GET /access-points/accessPointId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[accessPointId](#API_S3Files_GetAccessPoint_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the access point to retrieve information for.

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}|fsap-[0-9a-f]{17,40})`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "accessPointArn": "string",
   "accessPointId": "string",
   "clientToken": "string",
   "fileSystemId": "string",
   "name": "string",
   "ownerId": "string",
   "posixUser": {
      "gid": number,
      "secondaryGids": [ number ],
      "uid": number
   },
   "rootDirectory": {
      "creationPermissions": {
         "ownerGid": number,
         "ownerUid": number,
         "permissions": "string"
      },
      "path": "string"
   },
   "status": "string",
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[accessPointArn](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The ARN of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}`

**[accessPointId](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The ID of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}|fsap-[0-9a-f]{17,40})`

**[clientToken](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The client token used for idempotency when the access point was created.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `(.+)`

**[fileSystemId](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The ID of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[name](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The name of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

**[ownerId](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The AWS account ID of the access point owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

**[posixUser](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The POSIX identity configured for this access point.

Type: [PosixUser](api-s3files-posixuser.md) object

**[rootDirectory](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The root directory configuration for this access point.

Type: [RootDirectory](api-s3files-rootdirectory.md) object

**[status](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The current status of the access point.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

**[tags](#API_S3Files_GetAccessPoint_ResponseSyntax)**

The tags associated with the access point.

Type: Array of [Tag](api-s3files-tag.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/getaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/getaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMountTarget

GetFileSystem

All content copied from https://docs.aws.amazon.com/.
