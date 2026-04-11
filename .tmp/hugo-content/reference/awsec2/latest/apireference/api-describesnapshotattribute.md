# DescribeSnapshotAttribute

Describes the specified attribute of the specified snapshot. You can specify only one
attribute at a time.

For more information about EBS snapshots, see [Amazon EBS snapshots](../../../../services/ebs/latest/userguide/ebs-snapshots.md) in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The snapshot attribute you would like to view.

Type: String

Valid Values: `productCodes | createVolumePermission`

Required: Yes

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

**createVolumePermission**

The users and groups that have the permissions for creating volumes from the
snapshot.

Type: Array of [CreateVolumePermission](api-createvolumepermission.md) objects

**productCodes**

The product codes.

Type: Array of [ProductCode](api-productcode.md) objects

**requestId**

The ID of the request.

Type: String

**snapshotId**

The ID of the EBS snapshot.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the create volume permissions for the specified
snapshot.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeSnapshotAttribute
&SnapshotId=snap-1234567890abcdef0
&Attribute=createVolumePermission
&AUTHPARAMS
```

#### Sample Response

```

<DescribeSnapshotAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <snapshotId>snap-1234567890abcdef0</snapshotId>
   <createVolumePermission>
      <item>
         <group>all</group>
      </item>
   </createVolumePermission>
</DescribeSnapshotAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describesnapshotattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describesnapshotattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeServiceLinkVirtualInterfaces

DescribeSnapshots
