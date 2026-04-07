# CreateTransitGatewayConnectPeer

Creates a Connect peer for a specified transit gateway Connect attachment between a
transit gateway and an appliance.

The peer address and transit gateway address must be the same IP address family (IPv4 or IPv6).

For more information, see [Connect peers](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html#tgw-connect-peer)
in the _AWS Transit Gateways Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**BgpOptions**

The BGP options for the Connect peer.

Type: [TransitGatewayConnectRequestBgpOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayConnectRequestBgpOptions.html) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InsideCidrBlocks.N**

The range of inside IP addresses that are used for BGP peering. You must specify a
size /29 IPv4 CIDR block from the `169.254.0.0/16` range. The first address
from the range must be configured on the appliance as the BGP IP address. You can also
optionally specify a size /125 IPv6 CIDR block from the `fd00::/8`
range.

Type: Array of strings

Required: Yes

**PeerAddress**

The peer IP address (GRE outer IP address) on the appliance side of the Connect peer.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the Connect peer.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TransitGatewayAddress**

The peer IP address (GRE outer IP address) on the transit gateway side of the Connect peer, which must be
specified from a transit gateway CIDR block. If not specified, Amazon automatically assigns
the first available IP address from the transit gateway CIDR block.

Type: String

Required: No

**TransitGatewayAttachmentId**

The ID of the Connect attachment.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayConnectPeer**

Information about the Connect peer.

Type: [TransitGatewayConnectPeer](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayConnectPeer.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateTransitGatewayConnectPeer)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateTransitGatewayConnect

CreateTransitGatewayMeteringPolicy
