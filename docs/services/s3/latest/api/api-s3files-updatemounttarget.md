# UpdateMountTarget

Updates the mount target resource, specifically security group configurations.

## Request Syntax

```nohighlight

PUT /mount-targets/mountTargetId HTTP/1.1
Content-type: application/json

{
   "securityGroups": [ "string" ]
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[mountTargetId](#API_S3Files_UpdateMountTarget_RequestSyntax)**

The ID of the mount target to update.

Length Constraints: Minimum length of 22. Maximum length of 45.

Pattern: `fsmt-[0-9a-f]{17,40}`

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[securityGroups](#API_S3Files_UpdateMountTarget_RequestSyntax)**

An array of VPC security group IDs to associate with the mount target's network
interface. This replaces the existing security groups. All security groups must belong
to the same VPC as the mount target's subnet.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Length Constraints: Minimum length of 11. Maximum length of 43.

Pattern: `(sg-[0-9a-f]{8,40})`

Required: Yes

## Response Syntax

```nohighlight

HTTP/1.1 200
Content-type: application/json

{
   "availabilityZoneId": "string",
   "fileSystemId": "string",
   "ipv4Address": "string",
   "ipv6Address": "string",
   "mountTargetId": "string",
   "networkInterfaceId": "string",
   "ownerId": "string",
   "securityGroups": [ "string" ],
   "status": "string",
   "statusMessage": "string",
   "subnetId": "string",
   "vpcId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[availabilityZoneId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The Availability Zone ID where the mount target is located.

Type: String

**[fileSystemId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The ID of the S3 File System.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[ipv4Address](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The IPv4 address of the mount target.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 15.

Pattern: `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`

**[ipv6Address](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The IPv6 address of the mount target.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 39.

**[mountTargetId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The ID of the mount target.

Type: String

Length Constraints: Minimum length of 22. Maximum length of 45.

Pattern: `fsmt-[0-9a-f]{17,40}`

**[networkInterfaceId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The ID of the network interface associated with the mount target.

Type: String

**[ownerId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The AWS account ID of the mount target owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

**[securityGroups](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The security groups associated with the mount target.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Length Constraints: Minimum length of 11. Maximum length of 43.

Pattern: `(sg-[0-9a-f]{8,40})`

**[status](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The current status of the mount target.

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

**[statusMessage](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

Additional information about the mount target status.

Type: String

**[subnetId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The ID of the subnet where the mount target is located.

Type: String

Length Constraints: Minimum length of 15. Maximum length of 47.

Pattern: `subnet-[0-9a-f]{8,40}`

**[vpcId](#API_S3Files_UpdateMountTarget_ResponseSyntax)**

The ID of the VPC where the mount target is located.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/updatemounttarget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/updatemounttarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UntagResource

Amazon S3 on Outposts

All content copied from https://docs.aws.amazon.com/.
