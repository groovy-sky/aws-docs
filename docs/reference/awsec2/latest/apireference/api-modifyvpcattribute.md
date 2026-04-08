# ModifyVpcAttribute

Modifies the specified attribute of the specified VPC.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**EnableDnsHostnames**

Indicates whether the instances launched in the VPC get DNS hostnames. If enabled, instances in the VPC get DNS hostnames; otherwise, they do not.

You cannot modify the DNS resolution and DNS hostnames attributes in the same request. Use separate requests for each attribute. You can only enable DNS hostnames if you've enabled DNS support.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

Required: No

**EnableDnsSupport**

Indicates whether the DNS resolution is supported for the VPC. If enabled, queries to
the Amazon provided DNS server at the 169.254.169.253 IP address, or the reserved IP
address at the base of the VPC network range "plus two" succeed. If disabled, the Amazon
provided DNS service in the VPC that resolves public DNS hostnames to IP addresses is
not enabled.

You cannot modify the DNS resolution and DNS hostnames attributes in the same request. Use separate requests for each attribute.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

Required: No

**EnableNetworkAddressUsageMetrics**

Indicates whether Network Address Usage metrics are enabled for your VPC.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

Required: No

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

This example disables support for DNS hostnames in the specified
VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVpcAttribute
&VpcId=vpc-1a2b3c4d
&EnableDnsHostnames.Value=false
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVpcAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>d742de94-5f3e-4c3d-b6d4-440cexample</requestId>
    <return>true</return>
</ModifyVpcAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvpcattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpcattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVolumeAttribute

ModifyVpcBlockPublicAccessExclusion
