# Use Hive table compression

The compression options for Hive tables in Athena vary by engine version and file
format.

## Hive compression support in Athena engine version 3

The following table summarizes the compression format support in Athena engine version 3 for storage
file formats in Apache Hive. Text file format includes TSV, CSV, JSON, and custom SerDes
for text. "Yes" or "No" in a cell apply equally to read and write operations except
where noted. For the purposes of this table, CREATE TABLE, CTAS, and INSERT INTO are
considered write operations. For more information about using ZSTD compression levels in
Athena, see [Use ZSTD compression levels](compression-support-zstd-levels.md).

AvroIonORCParquetText fileBZIP2YesYesNoNoYesDEFLATEYesNoNoNoNoGZIPNoYesNoYesYesLZ4NoYesYes

Write - No

Read - Yes

YesLZONo

Write - No

Read - Yes

No

Write - No

Read - Yes

Write - No

Read - Yes

SNAPPYYesYesYesYesYesZLIBNoNoYesNoNoZSTDYesYesYesYesYesNONEYesYesYesYesYes

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use compression

Iceberg table
compression

All content copied from https://docs.aws.amazon.com/.
