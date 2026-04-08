# AssociateClientVpnTargetNetwork

Associates a target network with a Client VPN endpoint. A target network is a subnet in a VPC. You can associate multiple subnets from the same VPC with a Client VPN endpoint. You can associate only one subnet in each Availability Zone. We recommend that you associate at least two subnets to provide Availability Zone redundancy.

If you specified a VPC when you created the Client VPN endpoint or if you have previous subnet associations, the specified subnet must be in the same VPC. To specify a subnet that's in a different VPC, you must first modify the Client VPN endpoint ( [ModifyClientVpnEndpoint](api-modifyclientvpnendpoint.md)) and change the VPC that's associated with it.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SubnetId**

The ID of the subnet to associate with the Client VPN endpoint.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**associationId**

The unique ID of the target network association.

Type: String

**requestId**

The ID of the request.

Type: String

**status**

The current state of the target network association.

Type: [AssociationStatus](api-associationstatus.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example associates a subnet with a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateClientVpnTargetNetwork
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&SubnetId=subnet-057fa0918fEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<AssociateClientVpnTargetNetworkResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>7d1f819b-7f2a-4f81-aabf-81caeEXAMPLE</requestId>
    <status>
        <code>associating</code>
    </status>
    <associationId>cvpn-assoc-0822b0983cEXAMPLE</associationId>
</AssociateClientVpnTargetNetworkRespons>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/associateclientvpntargetnetwork.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associateclientvpntargetnetwork.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateCapacityReservationBillingOwner

AssociateDhcpOptions
