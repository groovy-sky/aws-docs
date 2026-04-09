# CreateAccessPoint

Creates an S3 File System Access Point for application-specific access with POSIX user
identity and root directory enforcement. Access points provide a way to manage access to
shared datasets in multi-tenant scenarios.

## Request Syntax

```nohighlight

PUT /access-points HTTP/1.1
Content-type: application/json

{
   "clientToken": "string",
   "fileSystemId": "string",
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
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[clientToken](#API_S3Files_CreateAccessPoint_RequestSyntax)**

A unique, case-sensitive identifier to ensure that the operation completes no more
than one time. If this token matches a previous request, AWS ignores the request, but
does not return an error.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `(.+)`

Required: No

**[fileSystemId](#API_S3Files_CreateAccessPoint_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

**[posixUser](#API_S3Files_CreateAccessPoint_RequestSyntax)**

The POSIX identity with uid, gid, and secondary group IDs for user enforcement when
accessing the file system through this access point.

Type: [PosixUser](api-s3files-posixuser.md) object

Required: No

**[rootDirectory](#API_S3Files_CreateAccessPoint_RequestSyntax)**

The root directory path for the access point, with optional creation permissions for
newly created directories.

Type: [RootDirectory](api-s3files-rootdirectory.md) object

Required: No

**[tags](#API_S3Files_CreateAccessPoint_RequestSyntax)**

An array of key-value pairs to apply to the access point for resource tagging.

Type: Array of [Tag](api-s3files-tag.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

Required: No

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

**[accessPointArn](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The Amazon Resource Name (ARN) of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}`

**[accessPointId](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The ID of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}|fsap-[0-9a-f]{17,40})`

**[clientToken](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The client token that was provided in the request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `(.+)`

**[fileSystemId](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The ID of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[name](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The name of the access point.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

**[ownerId](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The AWS account ID of the access point owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

**[posixUser](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The POSIX identity configured for this access point.

Type: [PosixUser](api-s3files-posixuser.md) object

**[rootDirectory](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The root directory configuration for this access point.

Type: [RootDirectory](api-s3files-rootdirectory.md) object

**[status](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The current status of the access point.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

**[tags](#API_S3Files_CreateAccessPoint_ResponseSyntax)**

The tags associated with the access point.

Type: Array of [Tag](api-s3files-tag.md) objects

Array Members: Minimum number of 1 item. Maximum number of 50 items.

## Errors

**ConflictException**

The request conflicts with the current state of the resource. This can occur when trying to create a resource that already exists or delete a resource that is in use.

**errorCode**

The error code associated with the exception.

**resourceId**

The identifier of the resource that caused the conflict.

**resourceType**

The type of the resource that caused the conflict.

HTTP Status Code: 409

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

**ServiceQuotaExceededException**

The request would exceed a service quota. Review your service quotas and either delete resources or request a quota increase.

**errorCode**

The error code associated with the exception.

HTTP Status Code: 402

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/createaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/createaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon S3 Files

CreateFileSystem

All content copied from https://docs.aws.amazon.com/.
