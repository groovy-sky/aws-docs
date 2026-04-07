# CreateNetworkInterfacePermission

Grants an AWS-authorized account permission to attach the specified
network interface to an instance in their account.

You can grant permission to a single AWS account only, and only one
account at a time.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AwsAccountId**

The AWS account ID.

Type: String

Required: No

**AwsService**

The AWS service. Currently not supported.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: Yes

**Permission**

The type of permission to grant.

Type: String

Valid Values: `INSTANCE-ATTACH | EIP-ASSOCIATE`

Required: Yes

## Response Elements

The following elements are returned by the service.

**interfacePermission**

Information about the permission for the network interface.

Type: [NetworkInterfacePermission](api-networkinterfacepermission.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example grants permission to account `123456789012` to attach
network interface `eni-1a2b3c4d` to an instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateNetworkInterfacePermission
&NetworkInterfaceId=eni-1a2b3c4d
&AwsAccountId=123456789012
&Permission=INSTANCE-ATTACH
&AUTHPARAMS
```

#### Sample Response

```

<CreateNetworkInterfacePermissionResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>e9633d41-093e-4944-981b-ca7example</requestId>
    <interfacePermission>
        <awsAccountId>123456789012</awsAccountId>
        <networkInterfaceId>eni-1a2b3c4d</networkInterfaceId>
        <networkInterfacePermissionId>eni-perm-06fd19020ede149ea</networkInterfacePermissionId>
        <permission>INSTANCE-ATTACH</permission>
        <permissionState>
            <state>GRANTED</state>
        </permissionState>
    </interfacePermission>
</CreateNetworkInterfacePermissionResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateNetworkInterfacePermission)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateNetworkInterfacePermission)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createnetworkinterfacepermission.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateNetworkInterfacePermission)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createnetworkinterfacepermission.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateNetworkInterface

CreatePlacementGroup
