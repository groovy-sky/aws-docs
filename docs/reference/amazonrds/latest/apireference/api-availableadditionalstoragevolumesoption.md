# AvailableAdditionalStorageVolumesOption

Contains the available options for additional storage volumes for a DB instance class.

## Contents

###### Note

In the following list, the required parameters are described first.

**MaxIops**

The maximum number of I/O operations per second (IOPS) that the additional storage volume supports.

Type: Integer

Required: No

**MaxIopsPerGib**

The maximum ratio of I/O operations per second (IOPS) to gibibytes (GiB) of storage for the additional storage volume.

Type: Double

Required: No

**MaxStorageSize**

The maximum amount of storage that you can allocate for the additional storage volume, in gibibytes (GiB).

Type: Integer

Required: No

**MaxStorageThroughput**

The maximum storage throughput that the additional storage volume supports, in mebibytes per second (MiBps).

Type: Integer

Required: No

**MinIops**

The minimum number of I/O operations per second (IOPS) that the additional storage volume supports.

Type: Integer

Required: No

**MinIopsPerGib**

The minimum ratio of I/O operations per second (IOPS) to gibibytes (GiB) of storage for the additional storage volume.

Type: Double

Required: No

**MinStorageSize**

The minimum amount of storage that you can allocate for the additional storage volume, in gibibytes (GiB).

Type: Integer

Required: No

**MinStorageThroughput**

The minimum storage throughput that the additional storage volume supports, in mebibytes per second (MiBps).

Type: Integer

Required: No

**StorageType**

The storage type for the additional storage volume.

Valid Values: `GP3 | IO2`

Type: String

Required: No

**SupportsIops**

Indicates whether the additional storage volume supports provisioned IOPS.

Type: Boolean

Required: No

**SupportsStorageAutoscaling**

Indicates whether the additional storage volume supports storage autoscaling.

Type: Boolean

Required: No

**SupportsStorageThroughput**

Indicates whether the additional storage volume supports configurable storage throughput.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/availableadditionalstoragevolumesoption.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/availableadditionalstoragevolumesoption.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/availableadditionalstoragevolumesoption.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AvailabilityZone

AvailableProcessorFeature

All content copied from https://docs.aws.amazon.com/.
