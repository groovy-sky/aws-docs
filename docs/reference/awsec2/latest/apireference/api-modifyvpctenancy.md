# ModifyVpcTenancy

Modifies the instance tenancy attribute of the specified VPC. You can change the
instance tenancy attribute of a VPC to `default` only. You cannot change the
instance tenancy attribute to `dedicated`.

After you modify the tenancy of the VPC, any new instances that you launch into the
VPC have a tenancy of `default`, unless you specify otherwise during launch.
The tenancy of any existing instances in the VPC is not affected.

For more information, see [Dedicated Instances](../../../../services/ec2/latest/userguide/dedicated-instance.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceTenancy**

The instance tenancy attribute for the VPC.

Type: String

Valid Values: `default`

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

Returns `true` if the request succeeds; otherwise, returns an
error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example modifies the tenancy of `vpc-1a2b3c4d` to
`default`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVpcTenancy
&VpcId=vpc-1a2b3c4d
&InstanceTenancy=default
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVpcTenancyResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <return>true</return>
    <requestId>125acea6-ba5c-4c6e-8e17-example</requestId>
</ModifyVpcTenancyResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvpctenancy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpctenancy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpcPeeringConnectionOptions

ModifyVpnConnection
