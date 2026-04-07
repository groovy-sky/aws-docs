# CreateRestoreImageTask

Starts a task that restores an AMI from an Amazon S3 object that was previously created by
using [CreateStoreImageTask](api-createstoreimagetask.md).

To use this API, you must have the required permissions. For more information, see [Permissions for storing and restoring AMIs using S3](../../../../services/ec2/latest/userguide/work-with-ami-store-restore.md#ami-s3-permissions) in the
_Amazon EC2 User Guide_.

For more information, see [Store and restore an AMI using\
S3](../../../../services/ec2/latest/userguide/ami-store-restore.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Bucket**

The name of the Amazon S3 bucket that contains the stored AMI object.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Name**

The name for the restored AMI. The name must be unique for AMIs in the Region for this
account. If you do not provide a name, the new AMI gets the same name as the original
AMI.

Type: String

Required: No

**ObjectKey**

The name of the stored AMI object in the bucket.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the AMI and snapshots on restoration. You can tag the AMI, the
snapshots, or both.

- To tag the AMI, the value for `ResourceType` must be
`image`.

- To tag the snapshots, the value for `ResourceType` must be
`snapshot`. The same tag is applied to all of the snapshots that are
created.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**imageId**

The AMI ID.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateRestoreImageTask)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateRestoreImageTask)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createrestoreimagetask.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateRestoreImageTask)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createrestoreimagetask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateReservedInstancesListing

CreateRoute
