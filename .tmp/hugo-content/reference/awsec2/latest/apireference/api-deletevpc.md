# DeleteVpc

Deletes the specified VPC. You must detach or delete all gateways and resources that are associated
with the VPC before you can delete it. For example, you must terminate all instances running in the VPC,
delete all security groups associated with the VPC (except the default one), delete all route tables
associated with the VPC (except the default one), and so on. When you delete the VPC, it deletes the
default security group, network ACL, and route table for the VPC.

If you created a flow log for the VPC that you are deleting, note that flow logs for deleted
VPCs are eventually automatically removed.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

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

### Example

This example deletes the specified VPC.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteVpc
&VpcId=vpc-1a2b3c4d
&AUTHPARAMS
```

#### Sample Response

```

<DeleteVpcResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>7a62c49f-347e-4fc4-9331-6e8eEXAMPLE</requestId>
   <return>true</return>
</DeleteVpcResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletevpc.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletevpc.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteVolume

DeleteVpcBlockPublicAccessExclusion
