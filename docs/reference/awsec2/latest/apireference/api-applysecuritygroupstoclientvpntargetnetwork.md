# ApplySecurityGroupsToClientVpnTargetNetwork

Applies a security group to the association between the target network and the Client VPN endpoint. This action replaces the existing
security groups with the specified security groups.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SecurityGroupId.N**

The IDs of the security groups to apply to the associated target network. Up to 5 security groups can
be applied to an associated target network.

Type: Array of strings

Required: Yes

**VpcId**

The ID of the VPC in which the associated target network is located.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**securityGroupIds**

The IDs of the applied security groups.

Type: Array of strings

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example applies a security group to a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ApplySecurityGroupsToClientVpnTargetNetwork
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&SecurityGroupId=sg-0618575f05EXAMPLE
&VpcId=vpc-3db97056EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<ApplySecurityGroupsToClientVpnTargetNetworkResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>5ef84b7f-505e-4e39-80cd-a11dbEXAMPLE</requestId>
    <securityGroupIds>
        <item>sg-0618575f05EXAMPLE</item>
    </securityGroupIds>
</ApplySecurityGroupsToClientVpnTargetNetworkResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/applysecuritygroupstoclientvpntargetnetwork.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AllocateIpamPoolCidr

AssignIpv6Addresses
