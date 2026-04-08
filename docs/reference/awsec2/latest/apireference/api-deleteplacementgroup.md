# DeletePlacementGroup

Deletes the specified placement group. You must terminate all instances in the
placement group before you can delete the placement group. For more information, see
[Placement groups](../../../../services/ec2/latest/userguide/placement-groups.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupName**

The name of the placement group.

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

This example deletes the placement group named
`XYZ-cluster`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeletePlacementGroup
&GroupName=XYZ-cluster
&AUTHPARAMS
```

#### Sample Response

```

<DeletePlacementGroupResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>d4904fd9-82c2-4ea5-adfe-a9cc3EXAMPLE</requestId>
   <return>true</return>
</DeletePlacementGroupResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deleteplacementgroup.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deleteplacementgroup.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteNetworkInterfacePermission

DeletePublicIpv4Pool
