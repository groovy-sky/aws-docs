# DisassociateRouteTable

Disassociates a subnet or gateway from a route table.

After you perform this action, the subnet no longer uses the routes in the route table.
Instead, it uses the routes in the VPC's main route table. For more information
about route tables, see [Route\
tables](../../../../services/vpc/latest/userguide/vpc-route-tables.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AssociationId**

The association ID representing the current association between the route table and subnet or gateway.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

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

This example disassociates the specified route table from the subnet it's
associated to.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisassociateRouteTable
&AssociationId=rtbassoc-0531ae3257956bdfb
&AUTHPARAMS
```

#### Sample Response

```

<DisassociateRouteTableResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</DisassociateRouteTableResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisassociateRouteTable)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisassociateRouteTable)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociateroutetable.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisassociateRouteTable)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociateroutetable.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateRouteServer

DisassociateSecurityGroupVpc
