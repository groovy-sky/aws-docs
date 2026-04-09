# ListMountTargets

Returns resource information for all mount targets with optional filtering by file
system, access point, and VPC.

## Request Syntax

```nohighlight

GET /mount-targets?accessPointId=accessPointId&fileSystemId=fileSystemId&maxResults=maxResults&nextToken=nextToken HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[accessPointId](#API_S3Files_ListMountTargets_RequestSyntax)**

Optional filter to list only mount targets associated with the specified access point
ID or Amazon Resource Name (ARN).

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}/access-point/fsap-[0-9a-f]{17,40}|fsap-[0-9a-f]{17,40})`

**[fileSystemId](#API_S3Files_ListMountTargets_RequestSyntax)**

Optional filter to list only mount targets associated with the specified S3 File System
ID or Amazon Resource Name (ARN). If provided, only mount targets for this file system will be returned in the
response.

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[maxResults](#API_S3Files_ListMountTargets_RequestSyntax)**

The maximum number of mount targets to return in a single response.

Valid Range: Minimum value of 1. Maximum value of 100.

**[nextToken](#API_S3Files_ListMountTargets_RequestSyntax)**

A pagination token returned from a previous call to continue listing mount
targets.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "mountTargets": [
      {
         "availabilityZoneId": "string",
         "fileSystemId": "string",
         "ipv4Address": "string",
         "ipv6Address": "string",
         "mountTargetId": "string",
         "networkInterfaceId": "string",
         "ownerId": "string",
         "status": "string",
         "statusMessage": "string",
         "subnetId": "string",
         "vpcId": "string"
      }
   ],
   "nextToken": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[mountTargets](#API_S3Files_ListMountTargets_ResponseSyntax)**

An array of mount target descriptions.

Type: Array of [ListMountTargetsDescription](api-s3files-listmounttargetsdescription.md) objects

**[nextToken](#API_S3Files_ListMountTargets_ResponseSyntax)**

A pagination token to use in a subsequent request if more results are
available.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/listmounttargets.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/listmounttargets.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListFileSystems

ListTagsForResource

All content copied from https://docs.aws.amazon.com/.
