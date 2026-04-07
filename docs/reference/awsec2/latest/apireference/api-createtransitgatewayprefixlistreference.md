# CreateTransitGatewayPrefixListReference

Creates a reference (route) to a prefix list in a specified transit gateway route table.

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

The ID of the prefix list that is used for destination matches.

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

This example creates a reference to a prefix list in the specified transit
gateway route table.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateTransitGatewayPrefixListReference
&TransitGatewayRouteTableId=tgw-rtb-0f98a0a5d09abcabc
&PrefixListId=pl-001122334455aabbc
&TransitGatewayAttachmentId=tgw-attach-01234567abcabcabc
&AUTHPARAMS
```

#### Sample Response

```

<CreateTransitGatewayPrefixListReferenceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>326fdc31-cd8d-491a-824a-example</requestId>
    <transitGatewayPrefixListReference>
        <blackhole>false</blackhole>
        <prefixListId>pl-001122334455aabbc</prefixListId>
        <prefixListOwnerId>123456789012</prefixListOwnerId>
        <state>pending</state>
        <transitGatewayAttachment>
            <resourceId>vpn-12312312312312312</resourceId>
            <resourceType>vpn</resourceType>
            <transitGatewayAttachmentId>tgw-attach-01234567abcabcabc</transitGatewayAttachmentId>
        </transitGatewayAttachment>
        <transitGatewayRouteTableId>tgw-rtb-0f98a0a5d09abcabc</transitGatewayRouteTableId>
    </transitGatewayPrefixListReference>
</CreateTransitGatewayPrefixListReferenceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateTransitGatewayPrefixListReference)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateTransitGatewayPrefixListReference)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateTransitGatewayPrefixListReference)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgatewayprefixlistreference.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGatewayPolicyTable

CreateTransitGatewayRoute
