# AttachInternetGateway

Attaches an internet gateway or a virtual private gateway to a VPC, enabling connectivity
between the internet and the VPC. For more information, see [Internet gateways](../../../../services/vpc/latest/userguide/vpc-internet-gateway.md) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InternetGatewayId**

The ID of the internet gateway.

Type: String

Required: Yes

**VpcId**

The ID of the VPC.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example attaches the internet gateway with the ID `igw-eaad4883`
to the VPC with the ID `vpc-11ad4878`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AttachInternetGateway
&InternetGatewayId=igw-eaad4883
&VpcId=vpc-11ad4878
&AUTHPARAMS
```

#### Sample Response

```

<AttachInternetGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</AttachInternetGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AttachInternetGateway)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AttachInternetGateway)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/attachinternetgateway.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AttachInternetGateway)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/attachinternetgateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttachClassicLinkVpc

AttachNetworkInterface
