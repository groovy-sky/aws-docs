# CreateLocalGatewayVirtualInterface

Create a virtual interface for a local gateway.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**LocalAddress**

The IP address assigned to the local gateway virtual interface on the Outpost side. Only IPv4 is supported.

Type: String

Required: Yes

**LocalGatewayVirtualInterfaceGroupId**

The ID of the local gateway virtual interface group.

Type: String

Required: Yes

**OutpostLagId**

References the Link Aggregation Group (LAG) that connects the Outpost to on-premises network devices.

Type: String

Required: Yes

**PeerAddress**

The peer IP address for the local gateway virtual interface. Only IPv4 is
supported.

Type: String

Required: Yes

**PeerBgpAsn**

The Autonomous System Number (ASN) of the Border Gateway Protocol (BGP) peer.

Type: Integer

Required: No

**PeerBgpAsnExtended**

The extended 32-bit ASN of the BGP peer for use with larger ASN values.

Type: Long

Required: No

**TagSpecification.N**

The tags to apply to a resource when the local gateway virtual interface is being created.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**Vlan**

The virtual local area network (VLAN) used for the local gateway virtual interface.

Type: Integer

Required: Yes

## Response Elements

The following elements are returned by the service.

**localGatewayVirtualInterface**

Information about the local gateway virtual interface.

Type: [LocalGatewayVirtualInterface](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LocalGatewayVirtualInterface.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateLocalGatewayVirtualInterface)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateLocalGatewayRouteTableVpcAssociation

CreateLocalGatewayVirtualInterfaceGroup
