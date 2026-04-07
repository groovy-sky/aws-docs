# DescribeSnapshotAttribute

Describes the specified attribute of the specified snapshot. You can specify only one
attribute at a time.

For more information about EBS snapshots, see [Amazon EBS snapshots](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-snapshots.html) in the _Amazon EBS User Guide_.

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

Type: Array of [CreateVolumePermission](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateVolumePermission.html) objects

**productCodes**

The product codes.

Type: Array of [ProductCode](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ProductCode.html) objects

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeSnapshotAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeSnapshotAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeServiceLinkVirtualInterfaces

DescribeSnapshots
