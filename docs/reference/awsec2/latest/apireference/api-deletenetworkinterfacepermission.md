# DeleteNetworkInterfacePermission

Deletes a permission for a network interface. By default, you cannot delete the
permission if the account for which you're removing the permission has attached the
network interface to an instance. However, you can force delete the permission,
regardless of any attachment.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Force**

Specify `true` to remove the permission even if the network interface is
attached to an instance.

Type: Boolean

Required: No

**NetworkInterfacePermissionId**

The ID of the network interface permission.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes the specified network interface permission.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteNetworkInterfacePermission
&NetworkInterfacePermissionId=eni-perm-06fd19020ede149ea
&AUTHPARAMS
```

#### Sample Response

```

<DeleteNetworkInterfacePermissionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>7a296942-8fa0-45a3-8406-09e9example</requestId>
    <return>true</return>
</DeleteNetworkInterfacePermissionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletenetworkinterfacepermission.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletenetworkinterfacepermission.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteNetworkInterface

DeletePlacementGroup
