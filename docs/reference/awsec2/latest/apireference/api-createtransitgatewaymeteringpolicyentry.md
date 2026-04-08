# CreateTransitGatewayMeteringPolicyEntry

Creates an entry in a transit gateway metering policy to define traffic measurement rules.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DestinationCidrBlock**

The destination CIDR block for traffic matching.

Type: String

Required: No

**DestinationPortRange**

The destination port range for traffic matching.

Type: String

Required: No

**DestinationTransitGatewayAttachmentId**

The ID of the destination transit gateway attachment for traffic matching.

Type: String

Required: No

**DestinationTransitGatewayAttachmentType**

The type of the destination transit gateway attachment for traffic matching. Note that the `tgw-peering` resource type has been deprecated. To configure metering policies for Connect, use the transport attachment type.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**MeteredAccount**

The AWS account ID to which the metered traffic should be attributed.

Type: String

Valid Values: `source-attachment-owner | destination-attachment-owner | transit-gateway-owner`

Required: Yes

**PolicyRuleNumber**

The rule number for the metering policy entry. Rules are processed in order from lowest to highest number.

Type: Integer

Required: Yes

**Protocol**

The protocol for traffic matching (1, 6, 17, etc.).

Type: String

Required: No

**SourceCidrBlock**

The source CIDR block for traffic matching.

Type: String

Required: No

**SourcePortRange**

The source port range for traffic matching.

Type: String

Required: No

**SourceTransitGatewayAttachmentId**

The ID of the source transit gateway attachment for traffic matching.

Type: String

Required: No

**SourceTransitGatewayAttachmentType**

The type of the source transit gateway attachment for traffic matching. Note that the `tgw-peering` resource type has been deprecated. To configure metering policies for Connect, use the transport attachment type.

Type: String

Valid Values: `vpc | vpn | vpn-concentrator | direct-connect-gateway | connect | peering | tgw-peering`

Required: No

**TransitGatewayMeteringPolicyId**

The ID of the transit gateway metering policy to add the entry to.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayMeteringPolicyEntry**

Information about the created transit gateway metering policy entry.

Type: [TransitGatewayMeteringPolicyEntry](api-transitgatewaymeteringpolicyentry.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgatewaymeteringpolicyentry.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGatewayMeteringPolicy

CreateTransitGatewayMulticastDomain
