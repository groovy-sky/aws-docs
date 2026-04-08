# ValidStorageOptions

Information about valid modifications that you can make to your DB instance.
Contains the result of a successful call to the
`DescribeValidDBInstanceModifications` action.

## Contents

###### Note

In the following list, the required parameters are described first.

**IopsToStorageRatio.DoubleRange.N**

The valid range of Provisioned IOPS to gibibytes of storage multiplier.
For example, 3-10,
which means that provisioned IOPS can be between 3 and 10 times storage.

Type: Array of [DoubleRange](api-doublerange.md) objects

Required: No

**ProvisionedIops.Range.N**

The valid range of provisioned IOPS.
For example, 1000-256,000.

Type: Array of [Range](api-range.md) objects

Required: No

**ProvisionedStorageThroughput.Range.N**

The valid range of provisioned storage throughput. For example,
500-4,000 mebibytes per second (MiBps).

Type: Array of [Range](api-range.md) objects

Required: No

**StorageSize.Range.N**

The valid range of storage in gibibytes (GiB).
For example, 100 to 16,384.

Type: Array of [Range](api-range.md) objects

Required: No

**StorageThroughputToIopsRatio.DoubleRange.N**

The valid range of storage throughput to provisioned IOPS ratios. For example,
0-0.25.

Type: Array of [DoubleRange](api-doublerange.md) objects

Required: No

**StorageType**

The valid storage types for your DB instance.
For example: gp2, gp3, io1, io2.

Type: String

Required: No

**SupportsStorageAutoscaling**

Indicates whether or not Amazon RDS can automatically scale storage for DB instances that use the new instance class.

Type: Boolean

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/validstorageoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/validstorageoptions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/validstorageoptions.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValidDBInstanceModificationsMessage

ValidVolumeOptions

All content copied from https://docs.aws.amazon.com/.
