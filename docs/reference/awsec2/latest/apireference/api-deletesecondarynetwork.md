# DeleteSecondaryNetwork

Deletes a secondary network. You must delete all secondary subnets in the secondary network before you can delete the secondary network.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensure Idempotency](https://docs.aws.amazon.com/ec2/latest/devguide/ec2-api-idempotency.html).

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

Type: [SecondaryNetwork](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SecondaryNetwork.html) object

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteSecondaryNetwork)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DeleteSecondaryNetwork)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteRouteTable

DeleteSecondarySubnet
