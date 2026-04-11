# DeleteMountTarget

Deletes the specified mount target. This operation is irreversible.

## Request Syntax

```nohighlight

DELETE /mount-targets/mountTargetId HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[mountTargetId](#API_S3Files_DeleteMountTarget_RequestSyntax)**

The ID of the mount target to delete.

Length Constraints: Minimum length of 22. Maximum length of 45.

Pattern: `fsmt-[0-9a-f]{17,40}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/deletemounttarget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/deletemounttarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteFileSystemPolicy

GetAccessPoint

All content copied from https://docs.aws.amazon.com/.
