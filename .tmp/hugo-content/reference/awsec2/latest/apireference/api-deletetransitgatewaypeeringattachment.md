# DeleteTransitGatewayPeeringAttachment

Deletes a transit gateway peering attachment.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TransitGatewayAttachmentId**

The ID of the transit gateway peering attachment.

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

This example deletes the specified transit gateway peering attachment by
specifying its attachment ID.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTransitGatewayPeeringAttachment
&TransitGatewayAttachmentId=tgw-attach-12345678901abcd12
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletetransitgatewaypeeringattachment.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteTransitGatewayMulticastDomain

DeleteTransitGatewayPolicyTable
