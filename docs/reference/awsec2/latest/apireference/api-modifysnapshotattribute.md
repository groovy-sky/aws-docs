# ModifySnapshotAttribute

Adds or removes permission settings for the specified snapshot. You may add or remove
specified AWS account IDs from a snapshot's list of create volume permissions, but you cannot
do both in a single operation. If you need to both add and remove account IDs for a snapshot,
you must use multiple operations. You can make up to 500 modifications to a snapshot in a single operation.

Encrypted snapshots and snapshots with AWS Marketplace product codes cannot be made
public. Snapshots encrypted with your default KMS key cannot be shared with other accounts.

For more information about modifying snapshot permissions, see [Share a snapshot](../../../../services/ebs/latest/userguide/ebs-modifying-snapshot-permissions.md) in the
_Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The snapshot attribute to modify. Only volume creation permissions can be modified.

Type: String

Valid Values: `productCodes | createVolumePermission`

Required: No

**CreateVolumePermission**

A JSON representation of the snapshot attribute modification.

Type: [CreateVolumePermissionModifications](api-createvolumepermissionmodifications.md) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**OperationType**

The type of operation to perform to the attribute.

Type: String

Valid Values: `add | remove`

Required: No

**SnapshotId**

The ID of the snapshot.

Type: String

Required: Yes

**UserGroup.N**

The group to modify for the snapshot.

Type: Array of strings

Required: No

**UserId.N**

The account ID to modify for the snapshot.

Type: Array of strings

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

This example makes the `snap-1234567890abcdef0` snapshot public, and gives
the account with ID `111122223333` permission to create volumes from
the snapshot.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifySnapshotAttribute
&SnapshotId=snap-1234567890abcdef0
&CreateVolumePermission.Add.1.UserId=111122223333
&CreateVolumePermission.Add.1.Group=all
&AUTHPARAMS
```

#### Sample Response

```

<ModifySnapshotAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</ModifySnapshotAttributeResponse>
```

### Example

This example makes the `snap-1234567890abcdef0` snapshot public, and
removes the account with ID `111122223333` from the list of users
with permission to create volumes from the snapshot.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifySnapshotAttribute
&SnapshotId=snap-1234567890abcdef0
&CreateVolumePermission.Remove.1.UserId=111122223333
&CreateVolumePermission.Add.1.Group=all
&AUTHPARAMS
```

#### Sample Response

```

<ModifySnapshotAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</ModifySnapshotAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifysnapshotattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifysnapshotattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifySecurityGroupRules

ModifySnapshotTier
