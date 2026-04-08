# DeleteTransitGatewayPrefixListReference

Deletes a reference (route) to a prefix list in a specified transit gateway route table.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

**TransitGatewayRouteTableId**

The ID of the route table.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**transitGatewayPrefixListReference**

Information about the deleted prefix list reference.

Type: [TransitGatewayPrefixListReference](api-transitgatewayprefixlistreference.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes the specified prefix list reference in the specified transit gateway route table.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteTransitGatewayPrefixListReference
&TransitGatewayRouteTableId=tgw-rtb-0f98a0a5d09abcabc
&PrefixListId=pl-001122334455aabbc
&AUTHPARAMS
```

#### Sample Response

```

<DeleteTransitGatewayPrefixListReferenceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>482823e8-8165-4312-86ee-example</requestId>
    <transitGatewayPrefixListReference>
        <blackhole>false</blackhole>
        <prefixListId>pl-001122334455aabbc</prefixListId>
        <prefixListOwnerId>123456789012</prefixListOwnerId>
        <state>deleting</state>
        <transitGatewayRouteTableId>tgw-rtb-0f98a0a5d09abcabc</transitGatewayRouteTableId>
    </transitGatewayPrefixListReference>
</DeleteTransitGatewayPrefixListReferenceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletetransitgatewayprefixlistreference.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteTransitGatewayPolicyTable

DeleteTransitGatewayRoute
