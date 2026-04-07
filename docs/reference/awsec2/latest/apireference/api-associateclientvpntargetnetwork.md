# AssociateClientVpnTargetNetwork

Associates a target network with a Client VPN endpoint. A target network is a subnet in a VPC. You can associate multiple subnets from the same VPC with a Client VPN endpoint. You can associate only one subnet in each Availability Zone. We recommend that you associate at least two subnets to provide Availability Zone redundancy.

If you specified a VPC when you created the Client VPN endpoint or if you have previous subnet associations, the specified subnet must be in the same VPC. To specify a subnet that's in a different VPC, you must first modify the Client VPN endpoint ( [ModifyClientVpnEndpoint](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ModifyClientVpnEndpoint.html)) and change the VPC that's associated with it.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
For more information, see [Ensuring idempotency](https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html).

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

Type: [AssociationStatus](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AssociationStatus.html) object

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/AssociateClientVpnTargetNetwork)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssociateCapacityReservationBillingOwner

AssociateDhcpOptions
