# CapacityManagerDataExportResponse

Contains information about a Capacity Manager data export configuration, including export settings, delivery status, and recent export activity.

## Contents

**capacityManagerDataExportId**

The unique identifier for the data export configuration.

Type: String

Required: No

**createTime**

The timestamp when the data export configuration was created.

Type: Timestamp

Required: No

**latestDeliveryS3LocationUri**

The S3 URI of the most recently delivered export file.

Type: String

Required: No

**latestDeliveryStatus**

The status of the most recent export delivery.

Type: String

Valid Values: `pending | in-progress | delivered | failed`

Required: No

**latestDeliveryStatusMessage**

A message describing the status of the most recent export delivery, including any error details if the delivery failed.

Type: String

Required: No

**latestDeliveryTime**

The timestamp when the most recent export was delivered to S3.

Type: Timestamp

Required: No

**outputFormat**

The file format of the exported data.

Type: String

Valid Values: `csv | parquet`

Required: No

**s3BucketName**

The name of the S3 bucket where export files are delivered.

Type: String

Required: No

**s3BucketPrefix**

The S3 key prefix used for organizing export files within the bucket.

Type: String

Required: No

**schedule**

The frequency at which data exports are generated.

Type: String

Valid Values: `hourly`

Required: No

**TagSet.N**

The tags associated with the data export configuration.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/capacitymanagerdataexportresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/capacitymanagerdataexportresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/capacitymanagerdataexportresponse.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CapacityManagerCondition

CapacityManagerDimension
