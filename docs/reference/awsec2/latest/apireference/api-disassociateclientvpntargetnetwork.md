# DisassociateClientVpnTargetNetwork

Disassociates a target network from the specified Client VPN endpoint. When you disassociate the
last target network from a Client VPN, the following happens:

- The route that was automatically added for the VPC is deleted

- All active client connections are terminated

- New client connections are disallowed

- The Client VPN endpoint's status changes to `pending-associate`

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId**

The ID of the target network association.

Type: String

Required: Yes

**ClientVpnEndpointId**

The ID of the Client VPN endpoint from which to disassociate the target network.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**associationId**

The ID of the target network association.

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

This example disassociates a target network from a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisassociateClientVpnTargetNetwork
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&AssociationId=cvpn-assoc-0bc4bd8cecEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DisassociateClientVpnTargetNetworkResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>61312648-93ec-4a86-a1d1-098c9EXAMPLE</requestId>
    <status>
        <code>disassociating</code>
    </status>
    <associationId>cvpn-assoc-0bc4bd8cecEXAMPLE</associationId>
</DisassociateClientVpnTargetNetworkResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociateclientvpntargetnetwork.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateCapacityReservationBillingOwner

DisassociateEnclaveCertificateIamRole
