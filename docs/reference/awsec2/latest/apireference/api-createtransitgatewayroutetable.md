# CreateTransitGatewayRouteTable

Creates a route table for the specified transit gateway.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecifications.N**

The tags to apply to the transit gateway route table.

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

**transitGatewayRouteTable**

Information about the transit gateway route table.

Type: [TransitGatewayRouteTable](api-transitgatewayroutetable.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a transit gateway route table for the specified transit
gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateTransitGatewayRouteTable
&TransitGatewayId=tgw-02f776b1a7EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<CreateTransitGatewayRouteTableResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>9c6751fa-a1ee-4006-92a8-c6cc1816a0f5</requestId>
    <transitGatewayRouteTable>
        <creationTime>2019-07-17T20:27:26.000Z</creationTime>
        <defaultAssociationRouteTable>false</defaultAssociationRouteTable>
        <defaultPropagationRouteTable>false</defaultPropagationRouteTable>
        <state>pending</state>
        <transitGatewayId>tgw-02f776b1a7EXAMPLE</transitGatewayId>
        <transitGatewayRouteTableId>tgw-rtb-0b6f6aaa01EXAMPLE</transitGatewayRouteTableId>
    </transitGatewayRouteTable>
</CreateTransitGatewayRouteTableResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createtransitgatewayroutetable.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createtransitgatewayroutetable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTransitGatewayRoute

CreateTransitGatewayRouteTableAnnouncement
