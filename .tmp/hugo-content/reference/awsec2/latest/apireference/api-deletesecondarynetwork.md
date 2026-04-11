# DeleteSecondaryNetwork

Deletes a secondary network. You must delete all secondary subnets in the secondary network before you can delete the secondary network.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensure Idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SecondaryNetworkId**

The ID of the secondary network.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier to ensure the idempotency of the request. Only returned if a client token was provided in the request.

Type: String

**requestId**

The ID of the request.

Type: String

**secondaryNetwork**

Information about the secondary network.

Type: [SecondaryNetwork](api-secondarynetwork.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes a secondary network.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteSecondaryNetwork
   &SecondaryNetworkId=sn-0123456789abcdef0
   &ClientToken=550e8400-e29b-41d4-a716-446655440000
   &AUTHPARAMS
```

#### Sample Response

```

<DeleteSecondaryNetworkResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <secondaryNetwork>
      <secondaryNetworkId>sn-0123456789abcdef0</secondaryNetworkId>
      <secondaryNetworkArn>arn:aws:ec2:us-east-2:123456789012:secondary-network/sn-0123456789abcdef0</secondaryNetworkArn>
      <ownerId>123456789012</ownerId>
      <type>rdma</type>
      <state>delete-in-progress</state>
      <ipv4CidrBlockAssociations>
         <item>
            <associationId>sn-cidr-assoc-56789901234abcdef0</associationId>
            <cidrBlock>10.0.0.0/16</cidrBlock>
            <state>disassociating</state>
         </item>
      </ipv4CidrBlockAssociations>
      <tagSet>
         <item>
            <key>Name</key>
            <value>Prod Secondary Network</value>
         </item>
      </tagSet>
   </secondaryNetwork>
   <clientToken>550e8400-e29b-41d4-a716-446655440000</clientToken>
</DeleteSecondaryNetworkResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletesecondarynetwork.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletesecondarynetwork.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteRouteTable

DeleteSecondarySubnet
