# AssociateDhcpOptions

Associates a set of DHCP options (that you've previously created) with the specified VPC, or associates no DHCP options with the VPC.

After you associate the options with the VPC, any existing instances and all new instances that you launch in that VPC use the options. You don't need to restart or relaunch the instances. They automatically pick up the changes within a few hours, depending on how frequently the instance renews its DHCP lease. You can explicitly renew the lease using the operating system on the instance.

For more information, see [DHCP option sets](../../../../services/vpc/latest/userguide/vpc-dhcp-options.md)
in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DhcpOptionsId**

The ID of the DHCP options set, or `default` to associate
no DHCP options with the VPC.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

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

### Example 1

This example associates the DHCP options with the ID `dopt-7a8b9c2d`
with the VPC with the ID `vpc-1a2b3c4d`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateDhcpOptions
&DhcpOptionsId=dopt-7a8b9c2d
&VpcId=vpc-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<AssociateDhcpOptionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</AssociateDhcpOptionsResponse>
```

### Example 2

This example changes the VPC with the ID `vpc-1a2b3c4d` to have no
associated DHCP options set.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AssociateDhcpOptions
&DhcpOptionsId=default
&VpcId=vpc-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<AssociateDhcpOptionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</AssociateDhcpOptionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AssociateDhcpOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AssociateDhcpOptions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/associatedhcpoptions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AssociateDhcpOptions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/associatedhcpoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AssociateClientVpnTargetNetwork

AssociateEnclaveCertificateIamRole
