# CreateTransitGatewayPeeringAttachment

Requests a transit gateway peering attachment between the specified transit gateway
(requester) and a peer transit gateway (accepter). The peer transit gateway can be in
your account or a different AWS account.

After you create the peering attachment, the owner of the accepter transit gateway
must accept the attachment request.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Options**

Requests a transit gateway peering attachment.

Type: [CreateTransitGatewayPeeringAttachmentRequestOptions](api-createtransitgatewaypeeringattachmentrequestoptions.md) object

Required: No

**PeerAccountId**

The ID of the AWS account that owns the peer transit gateway.

Type: String

Required: Yes

**PeerRegion**

The Region where the peer transit gateway is located.

Type: String

Required: Yes

**PeerTransitGatewayId**

The ID of the peer transit gateway with which to create the peering attachment.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the transit gateway peering attachment.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TransitGatewayId**

The ID of the transit gateway.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayPeeringAttachment**

The transit gateway peering attachment.

Type: [TransitGatewayPeeringAttachment](api-transitgatewaypeeringattachment.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a transit gateway peering attachment for the specified
transit gateways. The accepter (peer) transit gateway is in the
`us-west-2` Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateTransitGatewayPeeringAttachment
&TransitGatewayId=tgw-11223344aabbcc112
&PeerTransitGatewayId=tgw-1234567890abc1234
&PeerAccountId=123456789012
&PeerRegion=us-west-2
&AUTHPARAMS
```

#### Sample Response

```

<CreateTransitGatewayPeeringAttachmentResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>701859fa-6a57-4e55-858c-e63example</requestId>
    <transitGatewayPeeringAttachment>
        <accepterTgwInfo>
            <ownerId>123456789012</ownerId>
            <region>us-west-2</region>
            <transitGatewayId>tgw-1234567890abc1234</transitGatewayId>
        </accepterTgwInfo>
        <creationTime>2019-11-11T11:36:30.000Z</creationTime>
        <requesterTgwInfo>
            <ownerId>123456789012</ownerId>
            <region>us-east-1</region>
            <transitGatewayId>tgw-11223344aabbcc112</transitGatewayId>
        </requesterTgwInfo>
        <state>initiatingRequest</state>
        <transitGatewayAttachmentId>tgw-attach-0a73702c5c7123123</transitGatewayAttachmentId>
    </transitGatewayPeeringAttachment>
</CreateTransitGatewayPeeringAttachmentResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgatewaypeeringattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGatewayMulticastDomain

CreateTransitGatewayPolicyTable
