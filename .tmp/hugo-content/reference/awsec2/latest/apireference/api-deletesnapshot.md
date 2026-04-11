# DeleteSnapshot

Deletes the specified snapshot.

When you make periodic snapshots of a volume, the snapshots are incremental, and only the
blocks on the device that have changed since your last snapshot are saved in the new snapshot.
When you delete a snapshot, only the data not needed for any other snapshot is removed. So
regardless of which prior snapshots have been deleted, all active snapshots will have access
to all the information needed to restore the volume.

You cannot delete a snapshot of the root device of an EBS volume used by a registered AMI.
You must first deregister the AMI before you can delete the snapshot.

For more information, see [Delete an Amazon EBS snapshot](../../../../services/ebs/latest/userguide/ebs-deleting-snapshot.md) in the
_Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**SnapshotId**

The ID of the EBS snapshot.

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

This example request deletes the snapshot with the ID
`snap-1234567890abcdef0`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteSnapshot

&SnapshotId.1=snap-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<DeleteSnapshotResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DeleteSnapshotResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletesnapshot.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletesnapshot.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteSecurityGroup

DeleteSpotDatafeedSubscription
