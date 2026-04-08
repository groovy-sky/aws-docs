# ReleaseAddress

Releases the specified Elastic IP address.

\[Default VPC\] Releasing an Elastic IP address automatically disassociates it
from any instance that it's associated with. Alternatively, you can disassociate an Elastic IP address without
releasing it.

\[Nondefault VPC\] You must disassociate the Elastic IP address
before you can release it. Otherwise, Amazon EC2 returns an error ( `InvalidIPAddress.InUse`).

After releasing an Elastic IP address, it is released to the IP address pool.
Be sure to update your DNS records and any servers or devices that communicate with the address.
If you attempt to release an Elastic IP address that you already released, you'll get an
`AuthFailure` error if the address is already allocated to another AWS account.

After you release an Elastic IP address, you might be able to recover it.
For more information, see [Release an Elastic IP address](../../../../services/ec2/latest/userguide/using-instance-addressing-eips-releasing.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllocationId**

The allocation ID. This parameter is required.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**NetworkBorderGroup**

The set of Availability Zones, Local Zones, or Wavelength Zones from which AWS advertises
IP addresses.

If you provide an incorrect network border group, you receive an `InvalidAddress.NotFound` error.

Type: String

Required: No

**PublicIp**

Deprecated.

Type: String

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

This example releases the specified Elastic IP address.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ReleaseAddress
&AllocationId=eipalloc-5723d13e
&AUTHPARAMS
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/releaseaddress.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/releaseaddress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RejectVpcPeeringConnection

ReleaseHosts
