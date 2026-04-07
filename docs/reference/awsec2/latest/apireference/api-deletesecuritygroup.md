# DeleteSecurityGroup

Deletes a security group.

If you attempt to delete a security group that is associated with an instance or network interface, is
referenced by another security group in the same VPC, or has a VPC association, the operation fails with
`DependencyViolation`.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupId**

The ID of the security group.

Type: String

Required: No

**GroupName**

\[Default VPC\] The name of the security group. You can specify either the
security group name or the security group ID. For security groups in a nondefault VPC,
you must specify the security group ID.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**groupId**

The ID of the deleted security group.

Type: String

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes the specified security group.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteSecurityGroup
&GroupId=sg-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<DeleteSecurityGroupResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DeleteSecurityGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteSecurityGroup)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteSecurityGroup)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletesecuritygroup.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteSecurityGroup)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletesecuritygroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteSecondarySubnet

DeleteSnapshot
