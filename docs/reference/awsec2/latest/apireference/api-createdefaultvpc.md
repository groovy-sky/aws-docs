# CreateDefaultVpc

Creates a default VPC with a size `/16` IPv4 CIDR block and a default subnet
in each Availability Zone. For more information about the components of a default VPC,
see [Default VPCs](../../../../services/vpc/latest/userguide/default-vpc.md)
in the _Amazon VPC User Guide_. You cannot specify the components of the
default VPC yourself.

If you deleted your previous default VPC, you can create a default VPC. You cannot have
more than one default VPC per Region.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

**vpc**

Information about the VPC.

Type: [Vpc](api-vpc.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates a default VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateDefaultVpc
&AUTHPARAMS
```

#### Sample Response

```

<CreateDefaultVpcResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>056298f3-5f3e-48fb-9221-7c0example</requestId>
    <vpc>
        <cidrBlock>172.31.0.0/16</cidrBlock>
        <dhcpOptionsId>dopt-61079b07</dhcpOptionsId>
        <instanceTenancy>default</instanceTenancy>
        <ipv6CidrBlockAssociationSet/>
        <isDefault>true</isDefault>
        <state>pending</state>
        <tagSet/>
        <vpcId>vpc-3f139646</vpcId>
    </vpc>
</CreateDefaultVpcResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateDefaultVpc)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateDefaultVpc)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createdefaultvpc.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateDefaultVpc)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createdefaultvpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateDefaultSubnet

CreateDelegateMacVolumeOwnershipTask
