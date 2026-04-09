# Use Iceberg table compression

The compression options for Iceberg tables in Athena vary by engine version and file
format.

## Iceberg compression support in Athena engine version 3

The following table summarizes the compression format support in Athena engine version 3 for storage
file formats in Apache Iceberg. "Yes" or "No" in a cell apply equally to read and write
operations except where noted. For the purposes of this table, CREATE TABLE, CTAS, and
INSERT INTO are considered write operations. The default storage format for Iceberg in
Athena engine version 3 is Parquet. The default compression format for Iceberg in Athena engine version 3 is ZSTD. For
more information about using ZSTD compression levels in Athena, see [Use ZSTD compression levels](compression-support-zstd-levels.md).

AvroORCParquet (default)BZIP2NoNoNoGZIPYesNoYesLZ4NoYesNoSNAPPYYesYesYesZLIBNoYesNoZSTDYesYesYes (default)NONEYes (specify `None` or `Deflate`)YesYes (specify `None` or `Uncompressed`)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Hive table compression

ZSTD compression
levels

All content copied from https://docs.aws.amazon.com/.
