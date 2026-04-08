# CreateTransitGateway

Creates a transit gateway.

You can use a transit gateway to interconnect your virtual private clouds (VPC) and on-premises networks.
After the transit gateway enters the `available` state, you can attach your VPCs and VPN
connections to the transit gateway.

To attach your VPCs, use [CreateTransitGatewayVpcAttachment](api-createtransitgatewayvpcattachment.md).

To attach a VPN connection, use [CreateCustomerGateway](api-createcustomergateway.md) to create a customer
gateway and specify the ID of the customer gateway and the ID of the transit gateway in a call to
[CreateVpnConnection](api-createvpnconnection.md).

When you create a transit gateway, we create a default transit gateway route table and use it as the default association route table
and the default propagation route table. You can use [CreateTransitGatewayRouteTable](api-createtransitgatewayroutetable.md) to create
additional transit gateway route tables. If you disable automatic route propagation, we do not create a default transit gateway route table.
You can use [EnableTransitGatewayRouteTablePropagation](api-enabletransitgatewayroutetablepropagation.md) to propagate routes from a resource
attachment to a transit gateway route table. If you disable automatic associations, you can use [AssociateTransitGatewayRouteTable](api-associatetransitgatewayroutetable.md) to associate a resource attachment with a transit gateway route table.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Description**

A description of the transit gateway.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Options**

The transit gateway options.

Type: [TransitGatewayRequestOptions](api-transitgatewayrequestoptions.md) object

Required: No

**TagSpecification.N**

The tags to apply to the transit gateway.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGateway**

Information about the transit gateway.

Type: [TransitGateway](api-transitgateway.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createtransitgateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTrafficMirrorTarget

CreateTransitGatewayConnect
