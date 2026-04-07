# DeregisterTransitGatewayMulticastGroupMembers

Deregisters the specified members (network interfaces) from the transit gateway multicast group.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupIpAddress**

The IP address assigned to the transit gateway multicast group.

Type: String

Required: No

**NetworkInterfaceIds.N**

The IDs of the group members' network interfaces.

Type: Array of strings

Required: No

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**deregisteredMulticastGroupMembers**

Information about the deregistered members.

Type: [TransitGatewayMulticastDeregisteredGroupMembers](api-transitgatewaymulticastderegisteredgroupmembers.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example deregisters the network interface as a group member
`eni-0e246d3269EXAMPLE` from the multicast domain
`tgw-mcast-domain-0c4905cef7EXAMPLE`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeregisterTransitGatewayMulticastGroupMembers
&TransitGatewayMulticastDomainId=tgw-mcast-domain-0c4905cef7EXAMPLE
&NetworkInterfaceIds=eni-0e246d3269EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DeregisterTransitGatewayMulticastGroupMembersResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>6f4167cd-0870-4858-8872-f1c34EXAMPLE</requestId>
    <registeredMulticastGroupMembers>
        <groupIpAddress>224.0.1.0</groupIpAddress>
        <registeredNetworkInterfaceIds>
            <item>eni-0e246d3269EXAMPLE</item>
        </registeredNetworkInterfaceIds>
        <transitGatewayMulticastDomainId>tgw-mcast-domain-0c4905cef7EXAMPLE</transitGatewayMulticastDomainId>
    </registeredMulticastGroupMembers>
</DeregisterTransitGatewayMulticastGroupMembersResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeregisterTransitGatewayMulticastGroupMembers)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeregisterTransitGatewayMulticastGroupMembers)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeregisterTransitGatewayMulticastGroupMembers)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupmembers.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeregisterInstanceEventNotificationAttributes

DeregisterTransitGatewayMulticastGroupSources
