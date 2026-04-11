# InstanceSpecification

The instance details to specify which volumes should be snapshotted.

## Contents

**InstanceId**

The instance to specify which volumes should be snapshotted.

Type: String

Required: Yes

**ExcludeBootVolume**

Excludes the root volume from being snapshotted.

Type: Boolean

Required: No

**ExcludeDataVolumeId.N**

The IDs of the data (non-root) volumes to exclude from the multi-volume snapshot set.
If you specify the ID of the root volume, the request fails. To exclude the root volume,
use **ExcludeBootVolume**.

You can specify up to 40 volume IDs per request.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/instancespecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/instancespecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/instancespecification.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InstanceSecondaryInterfaceSpecificationRequest

InstanceState
