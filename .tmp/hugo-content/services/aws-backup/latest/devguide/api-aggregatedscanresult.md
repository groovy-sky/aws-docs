# AggregatedScanResult

Contains aggregated scan results across multiple scan operations, providing a summary of scan status and findings.

## Contents

**FailedScan**

A Boolean value indicating whether any of the aggregated scans failed.

Type: Boolean

Required: No

**Findings**

An array of findings discovered across all aggregated scans.

Type: Array of strings

Valid Values: `MALWARE`

Required: No

**LastComputed**

The timestamp when the aggregated scan result was last computed, in Unix format and Coordinated Universal Time (UTC).

Type: Timestamp

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/aggregatedscanresult.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/aggregatedscanresult.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/aggregatedscanresult.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AdvancedBackupSetting

BackupJob

All content copied from https://docs.aws.amazon.com/.
