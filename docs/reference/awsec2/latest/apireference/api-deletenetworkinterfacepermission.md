# DeleteNetworkInterfacePermission

Deletes a permission for a network interface. By default, you cannot delete the
permission if the account for which you're removing the permission has attached the
network interface to an instance. However, you can force delete the permission,
regardless of any attachment.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/CommonParameters.html).

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

For information about the errors that are common to all actions, see [Common client error codes](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html#CommonErrors).

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteNetworkInterfacePermission)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteNetworkInterfacePermission)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteNetworkInterface

DeletePlacementGroup
