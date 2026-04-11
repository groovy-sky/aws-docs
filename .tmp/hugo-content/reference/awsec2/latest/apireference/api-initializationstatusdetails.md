# InitializationStatusDetails

Information about the volume initialization. For more information, see [Initialize Amazon EBS volumes](../../../../services/ebs/latest/userguide/initalize-volume.md).

## Contents

**estimatedTimeToCompleteInSeconds**

The estimated remaining time, in seconds, for volume initialization to complete. Returns
`0` when volume initialization has completed.

Only available for volumes created with Amazon EBS Provisioned Rate for Volume Initialization.

Type: Long

Required: No

**initializationType**

The method used for volume initialization. Possible values include:

- `default` \- Volume initialized using the default volume initialization
rate or fast snapshot restore.

- `provisioned-rate` \- Volume initialized using an Amazon EBS Provisioned
Rate for Volume Initialization.

- `volume-copy` \- Volume copy initialized at the rate for volume copies.

Type: String

Valid Values: `default | provisioned-rate | volume-copy`

Required: No

**progress**

The current volume initialization progress as a percentage (0-100). Returns `100`
when volume initialization has completed.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/initializationstatusdetails.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/initializationstatusdetails.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/initializationstatusdetails.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

InferenceDeviceMemoryInfo

Instance
