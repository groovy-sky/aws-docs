# DescribeNetworkInterfacePermissions

Describes the permissions for your network interfaces.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Filter.N**

One or more filters.

- `network-interface-permission.network-interface-permission-id` -
The ID of the permission.

- `network-interface-permission.network-interface-id` \- The ID of the
network interface.

- `network-interface-permission.aws-account-id` \- The AWS account ID.

- `network-interface-permission.aws-service` \- The AWS
service.

- `network-interface-permission.permission` \- The type of permission
( `INSTANCE-ATTACH` \| `EIP-ASSOCIATE`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of items to return for this request. To get the next page of items,
make another request with the token returned in the output. If this parameter is not
specified, up to 50 results are returned by default. For more information, see [Pagination](query-requests.md#api-pagination).

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 255.

Required: No

**NetworkInterfacePermissionId.N**

The network interface permission IDs.

Type: Array of strings

Required: No

**NextToken**

The token returned from a previous paginated request. Pagination continues from the
end of the items returned by the previous request.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**networkInterfacePermissions**

The network interface permissions.

Type: Array of [NetworkInterfacePermission](api-networkinterfacepermission.md) objects

**nextToken**

The token to include in another request to get the next page of items. This value is
`null` when there are no more items to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all of your network interface permissions.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeNetworkInterfacePermissions
&AUTHPARAMS
```

#### Sample Response

```

<DescribeNetworkInterfacePermissionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>6d4fe5e1-4bd2-4e76-8980-04cexample</requestId>
    <networkInterfacePermissions>
        <item>
            <awsAccountId>123456789012</awsAccountId>
            <networkInterfaceId>eni-b909511a</networkInterfaceId>
            <networkInterfacePermissionId>eni-perm-06fd19020ede149ea</networkInterfacePermissionId>
            <permission>INSTANCE-ATTACH</permission>
            <permissionState>
                <state>GRANTED</state>
            </permissionState>
        </item>
    </networkInterfacePermissions>
</DescribeNetworkInterfacePermissionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeNetworkInterfacePermissions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeNetworkInterfacePermissions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describenetworkinterfacepermissions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeNetworkInterfacePermissions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describenetworkinterfacepermissions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeNetworkInterfaceAttribute

DescribeNetworkInterfaces
