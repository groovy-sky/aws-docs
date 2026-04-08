# DisassociateIpamByoasn

Remove the association between your Autonomous System Number (ASN) and your BYOIP CIDR. You may want to use this action to disassociate an ASN from a CIDR or if you want to swap ASNs.
For more information, see [Tutorial: Bring your ASN to IPAM](../../../../services/vpc/latest/ipam/tutorials-byoasn.md) in the _Amazon VPC IPAM guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Asn**

A public 2-byte or 4-byte ASN.

Type: String

Required: Yes

**Cidr**

A BYOIP CIDR.

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

**asnAssociation**

An ASN and BYOIP CIDR association.

Type: [AsnAssociation](api-asnassociation.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/disassociateipambyoasn.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disassociateipambyoasn.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisassociateInstanceEventWindow

DisassociateIpamResourceDiscovery
