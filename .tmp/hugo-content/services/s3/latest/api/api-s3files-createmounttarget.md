# CreateMountTarget

Creates a mount target resource as an endpoint for mounting the S3 File System from
compute resources in a specific Availability Zone and VPC. Mount targets provide network
access to the file system.

## Request Syntax

```nohighlight

PUT /mount-targets HTTP/1.1
Content-type: application/json

{
   "fileSystemId": "string",
   "ipAddressType": "string",
   "ipv4Address": "string",
   "ipv6Address": "string",
   "securityGroups": [ "string" ],
   "subnetId": "string"
}
```

## URI Request Parameters

The request does not use any URI parameters.

## Request Body

The request accepts the following data in JSON format.

**[fileSystemId](#API_S3Files_CreateMountTarget_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the S3 File System to create the mount target for.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

Required: Yes

**[ipAddressType](#API_S3Files_CreateMountTarget_RequestSyntax)**

The IP address type for the mount target. If not specified, `IPV4_ONLY` is
used. The IP address type must match the IP configuration of the specified
subnet.

Type: String

Valid Values: `IPV4_ONLY | IPV6_ONLY | DUAL_STACK`

Required: No

**[ipv4Address](#API_S3Files_CreateMountTarget_RequestSyntax)**

A specific IPv4 address to assign to the mount target. If not specified and the IP
address type supports IPv4, an address is automatically assigned from the subnet's
available IPv4 address range. The address must be within the subnet's CIDR block and not
already in use.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 15.

Pattern: `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`

Required: No

**[ipv6Address](#API_S3Files_CreateMountTarget_RequestSyntax)**

A specific IPv6 address to assign to the mount target. If not specified and the IP
address type supports IPv6, an address is automatically assigned from the subnet's
available IPv6 address range. The address must be within the subnet's IPv6 CIDR block
and not already in use.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 39.

Required: No

**[securityGroups](#API_S3Files_CreateMountTarget_RequestSyntax)**

An array of VPC security group IDs to associate with the mount target's network
interface. These security groups control network access to the mount target. If not
specified, the default security group for the subnet's VPC is used. All security groups
must belong to the same VPC as the subnet.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Length Constraints: Minimum length of 11. Maximum length of 43.

Pattern: `(sg-[0-9a-f]{8,40})`

Required: No

**[subnetId](#API_S3Files_CreateMountTarget_RequestSyntax)**

The ID of the subnet where the mount target will be created. The subnet must be in the
same AWS Region as the file system. For file systems with regional availability, you
can create mount targets in any subnet within the Region. The subnet determines the
Availability Zone where the mount target will be located.

Type: String

Length Constraints: Minimum length of 15. Maximum length of 47.

Pattern: `subnet-[0-9a-f]{8,40}`

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

**[availabilityZoneId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The unique and consistent identifier of the Availability Zone where the mount target
is located. For example, `use1-az1` is an Availability Zone ID for the
`us-east-1`
AWS Region, and it has the same location in every
AWS account.

Type: String

**[fileSystemId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The ID of the S3 File System associated with the mount target.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 128.

Pattern: `(arn:aws[-a-z]*:s3files:[0-9a-z-:]+:file-system/fs-[0-9a-f]{17,40}|fs-[0-9a-f]{17,40})`

**[ipv4Address](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The IPv4 address assigned to the mount target.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 15.

Pattern: `[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}`

**[ipv6Address](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The IPv6 address assigned to the mount target.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 39.

**[mountTargetId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The ID of the mount target, assigned by S3 Files. This ID is used to reference the
mount target in subsequent API calls.

Type: String

Length Constraints: Minimum length of 22. Maximum length of 45.

Pattern: `fsmt-[0-9a-f]{17,40}`

**[networkInterfaceId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The ID of the network interface that S3 Files created when it created the mount
target. This network interface is managed by the service.

Type: String

**[ownerId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The AWS account ID of the mount target owner.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 12.

Pattern: `(\d{12})|(\d{4}-{4}-\d{4})`

**[securityGroups](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The security groups associated with the mount target's network interface.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Length Constraints: Minimum length of 11. Maximum length of 43.

Pattern: `(sg-[0-9a-f]{8,40})`

**[status](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The lifecycle state of the mount target. Valid values are: `AVAILABLE` (the
mount target is available for use), `CREATING` (the mount target is being
created), `DELETING` (the mount target is being deleted),
`DELETED` (the mount target has been deleted), or `ERROR` (the
mount target is in an error state), or `UPDATING` (the mount target is being
updated).

Type: String

Valid Values: `available | creating | deleting | deleted | error | updating`

**[statusMessage](#API_S3Files_CreateMountTarget_ResponseSyntax)**

Additional information about the mount target status. This field provides more details
when the status is `ERROR`, or during state transitions.

Type: String

**[subnetId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The ID of the subnet where the mount target is located.

Type: String

Length Constraints: Minimum length of 15. Maximum length of 47.

Pattern: `subnet-[0-9a-f]{8,40}`

**[vpcId](#API_S3Files_CreateMountTarget_ResponseSyntax)**

The ID of the VPC where the mount target is located.

Type: String

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for Python](../../../goto/boto3/s3files-2025-05-05/createmounttarget.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3files-2025-05-05/createmounttarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateFileSystem

DeleteAccessPoint

All content copied from https://docs.aws.amazon.com/.
