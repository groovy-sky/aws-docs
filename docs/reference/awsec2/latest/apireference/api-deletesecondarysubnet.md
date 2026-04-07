# DeleteSecondarySubnet

Deletes a secondary subnet. A secondary subnet must not contain any secondary interfaces prior to deletion.

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

**SecondarySubnetId**

The ID of the secondary subnet to delete.

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

**secondarySubnet**

Information about the secondary subnet being deleted.

Type: [SecondarySubnet](api-secondarysubnet.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes a secondary subnet.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteSecondarySubnet
   &SecondarySubnetId=ss-0123456789abcdef0
   &ClientToken=550e8400-e29b-41d4-a716-446655440000
   &AUTHPARAMS
```

#### Sample Response

```

<DeleteSecondarySubnetResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <secondarySubnet>
      <secondarySubnetId>ss-0123456789abcdef0</secondarySubnetId>
      <secondarySubnetArn>arn:aws:ec2:us-east-2:123456789012:secondary-subnet/ss-0123456789abcdef0</secondarySubnetArn>
      <secondaryNetworkId>sn-0123456789abcdef0</secondaryNetworkId>
      <secondaryNetworkType>rdma</secondaryNetworkType>
      <ownerId>123456789012</ownerId>
      <availabilityZoneId>use2-az1</availabilityZoneId>
      <availabilityZone>us-east-2a</availabilityZone>
      <ipv4CidrBlockAssociations>
         <item>
            <associationId>ss-cidr-assoc-56789901234abcdef0</associationId>
            <cidrBlock>10.0.0.0/24</cidrBlock>
            <state>disassociating</state>
         </item>
      </ipv4CidrBlockAssociations>
      <state>delete-in-progress</state>
      <tagSet>
         <item>
            <key>Name</key>
            <value>Prod Secondary Subnet</value>
         </item>
      </tagSet>
   </secondarySubnet>
   <clientToken>550e8400-e29b-41d4-a716-446655440000</clientToken>
</DeleteSecondarySubnetResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteSecondarySubnet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteSecondarySubnet)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletesecondarysubnet.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteSecondarySubnet)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletesecondarysubnet.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteSecondaryNetwork

DeleteSecurityGroup
