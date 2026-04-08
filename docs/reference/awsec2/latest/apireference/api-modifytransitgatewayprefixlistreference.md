# ModifyTransitGatewayPrefixListReference

Modifies a reference (route) to a prefix list in a specified transit gateway route table.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Blackhole**

Indicates whether to drop traffic that matches this route.

Type: Boolean

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

**TransitGatewayAttachmentId**

The ID of the attachment to which traffic is routed.

Type: String

Required: No

**TransitGatewayRouteTableId**

The ID of the transit gateway route table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayPrefixListReference**

Information about the prefix list reference.

Type: [TransitGatewayPrefixListReference](api-transitgatewayprefixlistreference.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example modifies the prefix list reference in the specified route table
by changing the attachment to which traffic is routed.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyTransitGatewayPrefixListReference
&TransitGatewayRouteTableId=tgw-rtb-0f98a0a5d09abcabc
&PrefixListId=pl-001122334455aabbc
&TransitGatewayAttachmentId=tgw-attach-11223344aabbcc112
&AUTHPARAMS
```

#### Sample Response

```

<ModifyTransitGatewayPrefixListReferenceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>bbd3e523-3e5b-4d3b-b010-example</requestId>
    <transitGatewayPrefixListReference>
        <blackhole>false</blackhole>
        <prefixListId>pl-001122334455aabbc</prefixListId>
        <prefixListOwnerId>123456789012</prefixListOwnerId>
        <state>modifying</state>
        <transitGatewayAttachment>
            <resourceId>tgw-012233aabbcc11223</resourceId>
            <resourceType>peering</resourceType>
            <transitGatewayAttachmentId>tgw-attach-11223344aabbcc112</transitGatewayAttachmentId>
        </transitGatewayAttachment>
        <transitGatewayRouteTableId>tgw-rtb-0f98a0a5d09abcabc</transitGatewayRouteTableId>
    </transitGatewayPrefixListReference>
</ModifyTransitGatewayPrefixListReferenceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifytransitgatewayprefixlistreference.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyTransitGatewayMeteringPolicy

ModifyTransitGatewayVpcAttachment
