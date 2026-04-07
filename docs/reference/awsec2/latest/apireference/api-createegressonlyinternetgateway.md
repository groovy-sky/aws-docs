# CreateEgressOnlyInternetGateway

\[IPv6 only\] Creates an egress-only internet gateway for your VPC. An egress-only
internet gateway is used to enable outbound communication over IPv6 from instances in
your VPC to the internet, and prevents hosts outside of your VPC from initiating an IPv6
connection with your instance.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to assign to the egress-only internet gateway.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VpcId**

The ID of the VPC for which to create the egress-only internet gateway.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**clientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the
request.

Type: String

**egressOnlyInternetGateway**

Information about the egress-only internet gateway.

Type: [EgressOnlyInternetGateway](api-egressonlyinternetgateway.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates an egress-only internet gateway in VPC
vpc-1a2b3c4d.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateEgressOnlyInternetGateway
&VpcId=vpc-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<CreateEgressOnlyInternetGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>c617595f-6c29-4a00-a941-example</requestId>
    <egressOnlyInternetGateway>
        <attachmentSet>
            <item>
                <state>attached</state>
                <vpcId>vpc-1a2b3c4d</vpcId>
            </item>
        </attachmentSet>
        <egressOnlyInternetGatewayId>eigw-01eadbd45ecd7943f</egressOnlyInternetGatewayId>
    </egressOnlyInternetGateway>
</CreateEgressOnlyInternetGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateEgressOnlyInternetGateway)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateEgressOnlyInternetGateway)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createegressonlyinternetgateway.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateEgressOnlyInternetGateway)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createegressonlyinternetgateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateDhcpOptions

CreateFleet
