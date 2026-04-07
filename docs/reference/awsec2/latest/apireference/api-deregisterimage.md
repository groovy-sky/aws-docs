# DeregisterImage

Deregisters the specified AMI. A deregistered AMI can't be used to launch new
instances.

If a deregistered EBS-backed AMI matches a Recycle Bin retention rule, it moves to the
Recycle Bin for the specified retention period. It can be restored before its retention period
expires, after which it is permanently deleted. If the deregistered AMI doesn't match a
retention rule, it is permanently deleted immediately. For more information, see [Recover deleted Amazon EBS\
snapshots and EBS-backed AMIs with Recycle Bin](../../../../services/ec2/latest/userguide/recycle-bin.md) in the _Amazon EBS User_
_Guide_.

When deregistering an EBS-backed AMI, you can optionally delete its associated snapshots
at the same time. However, if a snapshot is associated with multiple AMIs, it won't be deleted
even if specified for deletion, although the AMI will still be deregistered.

Deregistering an AMI does not delete the following:

- Instances already launched from the AMI. You'll continue to incur usage costs for the
instances until you terminate them.

- For EBS-backed AMIs: Snapshots that are associated with multiple AMIs. You'll continue
to incur snapshot storage costs.

- For instance store-backed AMIs: The files uploaded to Amazon S3 during AMI creation. You'll
continue to incur S3 storage costs.

For more information, see [Deregister an Amazon EC2 AMI](../../../../services/ec2/latest/userguide/deregister-ami.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DeleteAssociatedSnapshots**

Specifies whether to delete the snapshots associated with the AMI during
deregistration.

###### Note

If a snapshot is associated with multiple AMIs, it is not deleted, regardless of this setting.

Default: The snapshots are not deleted.

Type: Boolean

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageId**

The ID of the AMI.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**deleteSnapshotResultSet**

The deletion result for each snapshot associated with the AMI, including the snapshot ID
and its success or error code.

Type: Array of [DeleteSnapshotReturnCode](api-deletesnapshotreturncode.md) objects

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request deregisters the specified AMI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeregisterImage
&ImageId=ami-4fa54026
&AUTHPARAMS
```

#### Sample Response

```

<DeregisterImageResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DeregisterImageResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeregisterImage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeregisterImage)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deregisterimage.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeregisterImage)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deregisterimage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeprovisionPublicIpv4PoolCidr

DeregisterInstanceEventNotificationAttributes
