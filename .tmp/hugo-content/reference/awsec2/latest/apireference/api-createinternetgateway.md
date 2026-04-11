# CreateInternetGateway

Creates an internet gateway for use with a VPC. After creating the internet gateway,
you attach it to a VPC using [AttachInternetGateway](api-attachinternetgateway.md).

For more information, see [Internet gateways](../../../../services/vpc/latest/userguide/vpc-internet-gateway.md) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to assign to the internet gateway.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**internetGateway**

Information about the internet gateway.

Type: [InternetGateway](api-internetgateway.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates an internet gateway.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateInternetGateway
&AUTHPARAMS
```

#### Sample Response

```

<CreateInternetGatewayResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <internetGateway>
      <internetGatewayId>igw-eaad4883</internetGatewayId>
      <attachmentSet/>
      <tagSet/>
   </internetGateway>
</CreateInternetGatewayResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createinternetgateway.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createinternetgateway.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateInstanceExportTask

CreateInterruptibleCapacityReservationAllocation
