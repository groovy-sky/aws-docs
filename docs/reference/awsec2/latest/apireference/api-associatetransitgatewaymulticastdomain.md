# AssociateTransitGatewayMulticastDomain

Associates the specified subnets and transit gateway attachments with the specified transit gateway multicast domain.

The transit gateway attachment must be in the available state before you can add a resource. Use [DescribeTransitGatewayAttachments](api-describetransitgatewayattachments.md)
to see the state of the attachment.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SubnetIds.N**

The IDs of the subnets to associate with the transit gateway multicast domain.

Type: Array of strings

Required: Yes

**TransitGatewayAttachmentId**

The ID of the transit gateway attachment to associate with the transit gateway multicast domain.

Type: String

Required: Yes

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**associations**

Information about the transit gateway multicast domain associations.

Type: [TransitGatewayMulticastDomainAssociations](api-transitgatewaymulticastdomainassociations.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example associates the transit gateway attachment
`tgw-attach-028c1dd0f8EXAMPLE` with the multicast domain
`tgw-mcast-domain-0c4905cef7EXAMPLE`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateTransitGatewayMulticastDomain
&TransitGatewayAttachmentId=tgw-attach-028c1dd0f8EXAMPLE
&TransitGatewayMulticastDomainId=tgw-mcast-domain-0c4905cef7EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<AssociateTransitGatewayMulticastDomainResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>fa968e66-0290-4479-a8ca-e5c83EXAMPLE</requestId>
    <associations>
        <resourceId>vpc-01128d2c24EXAMPLE</resourceId>
        <resourceType>vpc</resourceType>
        <subnets>
            <item>
                <state>associating</state>
                <subnetId>subnet-000de86e3bEXAMPLE</subnetId>
            </item>
        </subnets>
        <transitGatewayAttachmentId>tgw-attach-028c1dd0f8EXAMPLE</transitGatewayAttachmentId>
        <transitGatewayMulticastDomainId>tgw-mcast-domain-0c4905cef7EXAMPLE</transitGatewayMulticastDomainId>
    </associations>
</AssociateTransitGatewayMulticastDomainResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associatetransitgatewaymulticastdomain.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateSubnetCidrBlock

AssociateTransitGatewayPolicyTable
