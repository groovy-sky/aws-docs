# AssociateAddress

Associates an Elastic IP address, or carrier IP address (for instances that are in
subnets in Wavelength Zones) with an instance or a network interface. Before you can use an
Elastic IP address, you must allocate it to your account.

If the Elastic IP address is already
associated with a different instance, it is disassociated from that instance and associated
with the specified instance. If you associate an Elastic IP address with an instance that has
an existing Elastic IP address, the existing address is disassociated from the instance, but
remains allocated to your account.

\[Subnets in Wavelength Zones\] You can associate an IP address from the telecommunication
carrier to the instance or network interface.

You cannot associate an Elastic IP address with an interface in a different network border group.

###### Important

This is an idempotent operation. If you perform the operation more than once, Amazon EC2
doesn't return an error, and you may be charged for each time the Elastic IP address is
remapped to the same instance. For more information, see the _Elastic IP_
_Addresses_ section of [Amazon EC2\
Pricing](http://aws.amazon.com/ec2/pricing).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationId**

The allocation ID. This is required.

Type: String

Required: No

**AllowReassociation**

Reassociation is automatic, but you can specify false to ensure the operation fails if the Elastic IP address is already associated with another resource.

Type: Boolean

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance. The instance must have exactly one attached network interface.
You can specify either the instance ID or the network interface ID, but not both.

Type: String

Required: No

**NetworkInterfaceId**

The ID of the network interface. If the instance has more than one network interface, you must specify a network interface ID.

You can specify either the instance ID or the network interface ID, but not both.

Type: String

Required: No

**PrivateIpAddress**

The primary or secondary private IP address to associate with the Elastic IP address. If no private IP address is specified, the Elastic IP address is associated with the primary private IP address.

Type: String

Required: No

**PublicIp**

Deprecated.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**associationId**

The ID that represents the association of the Elastic IP address with an instance.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request associates a Elastic IP address with an instance.
The `AllowReassignment` parameter allows the Elastic IP address to be associated with
the specified instance even if it's already associated with a different instance or a network interface.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateAddress
&InstanceId=i-0598c7d356eba48d7
&AllocationId=eipalloc-5723d13e
&AllowReassignment=true
&AUTHPARAMS
```

#### Sample Response

```

<AssociateAddressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
   <associationId>eipassoc-fc5ca095</associationId>
</AssociateAddressResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/associateaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associateaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssignPrivateNatGatewayAddress

AssociateCapacityReservationBillingOwner
