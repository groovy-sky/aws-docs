# Considerations for DB cluster snapshot exports

## Limitations

Exporting DB snapshot data to Amazon S3 has the following limitations:

- You can't run multiple export tasks for the same DB cluster snapshot simultaneously. This applies to both full and
partial exports.

- You can have up to five concurrent DB snapshot export tasks in progress per AWS account.

- You can't export snapshot data from Aurora Serverless v1 DB clusters to S3.

- Exports to S3 don't support S3 prefixes containing a colon (:).

- The following characters in the S3 file path are converted to underscores (\_) during export:

```nohighlight

\ ` " (space)
```

- If a database, schema, or table has characters in its name other than the following, partial export isn't
supported. However, you can export the entire DB snapshot.

- Latin letters (A–Z)

- Digits (0–9)

- Dollar symbol ($)

- Underscore (\_)

- Spaces ( ) and certain characters aren't supported in database table column names. Tables with the following
characters in column names are skipped during export:

```nohighlight

, ; { } ( ) \n \t = (space)
```

- Tables with slashes (/) in their names are skipped during export.

- Aurora PostgreSQL temporary and unlogged tables are skipped during export.

- If the data contains a large object, such as a BLOB or CLOB, that is close to
or greater than 500 MB, then the export fails.

- If a table contains a large row that is close to or greater than 2 GB, then
the table is skipped during export.

- For partial exports, the `ExportOnly` list has a maximum size of
200 KB.

- We strongly recommend that you use a unique name for each export task. If you don't use a unique task name, you might
receive the following error message:

**`ExportTaskAlreadyExistsFault: An error occurred (ExportTaskAlreadyExists) when calling the StartExportTask**
**operation: The export task with the ID xxxxx already exists.`**

- You can delete a snapshot while you're exporting its data to S3, but you're still charged for the storage
costs for that snapshot until the export task has completed.

- You can't restore exported snapshot data from S3 to a new DB cluster.

## File naming convention

Exported data for specific tables is stored in the format `base_prefix/files`, where the base prefix is
the following:

```nohighlight

export_identifier/database_name/schema_name.table_name/
```

For example:

```nohighlight

export-1234567890123-459/rdststdb/rdststdb.DataInsert_7ADB5D19965123A2/
```

There are two conventions for how files are named.

- Current convention:

```nohighlight

batch_index/part-partition_index-random_uuid.format-based_extension
```

The batch index is a sequence number that represents a batch of data read from the table. If we can't partition your
table into small chunks to be exported in parallel, there will be multiple batch indexes. The same thing happens if your
table is partitioned into multiple tables. There will be multiple batch indexes, one for each of the table partitions of
your main table.

If we can partition your table into small chunks to be read in parallel, there will be only the batch index
`1` folder.

Inside the batch index folder, there are one or more Parquet files that contain your table's data. The prefix of the
Parquet filename is `part-partition_index`. If your table is partitioned,
there will be multiple files starting with the partition index `00000`.

There can be gaps in the partition index sequence. This happens because each partition is obtained from a ranged query
in your table. If there is no data in the range of that partition, then that sequence number is skipped.

For example, suppose that the `id` column is the table's primary key, and its minimum and maximum values
are `100` and `1000`. When we try to export this table with nine partitions, we read it with
parallel queries such as the following:

```nohighlight

SELECT * FROM table WHERE id <= 100 AND id < 200
  	SELECT * FROM table WHERE id <= 200 AND id < 300
```

This should generate nine files, from
`part-00000-random_uuid.gz.parquet` to
`part-00008-random_uuid.gz.parquet`. However, if there are no rows
with IDs between `200` and `350`, one of the completed partitions is empty, and no file is created
for it. In the previous example, `part-00001-random_uuid.gz.parquet` isn't
created.

- Older convention:

```nohighlight

part-partition_index-random_uuid.format-based_extension
```

This is the same as the current convention, but without the `batch_index`
prefix, for example:

```nohighlight

part-00000-c5a881bb-58ff-4ee6-1111-b41ecff340a3-c000.gz.parquet
  	part-00001-d7a881cc-88cc-5ab7-2222-c41ecab340a4-c000.gz.parquet
  	part-00002-f5a991ab-59aa-7fa6-3333-d41eccd340a7-c000.gz.parquet

```

The file naming convention is subject to change. Therefore, when reading target
tables, we recommend that you read everything inside the base prefix for the
table.

## Data conversion when exporting to an Amazon S3 bucket

When you export a DB snapshot to an Amazon S3 bucket, Amazon Aurora converts data to,
exports data in, and stores data in the Parquet format. For more information about
Parquet, see the [Apache\
Parquet](https://parquet.apache.org/docs) website.

Parquet stores all data as one of the following primitive types:

- BOOLEAN

- INT32

- INT64

- INT96

- FLOAT

- DOUBLE

- BYTE\_ARRAY – A variable-length byte array, also known as binary

- FIXED\_LEN\_BYTE\_ARRAY – A fixed-length byte array used when the values have a
constant size

The Parquet data types are few to reduce the complexity of reading and writing the format.
Parquet provides logical types for extending primitive types. A _logical type_ is implemented as an annotation with the data in a
`LogicalType` metadata field. The logical type annotation explains how to
interpret the primitive type.

When the `STRING` logical type annotates a `BYTE_ARRAY` type, it
indicates that the byte array should be interpreted as a UTF-8 encoded character string.
After an export task completes, Amazon Aurora notifies you if any string conversion occurred.
The underlying data exported is always the same as the data from the source. However,
due to the encoding difference in UTF-8, some characters might appear different from the
source when read in tools such as Athena.

For more information, see [Parquet\
logical type definitions](https://github.com/apache/parquet-format/blob/master/LogicalTypes.md) in the Parquet documentation.

###### Topics

- [MySQL data type mapping to Parquet](#aurora-export-snapshot.data-types.MySQL)

- [PostgreSQL data type mapping to Parquet](#aurora-export-snapshot.data-types.PostgreSQL)

### MySQL data type mapping to Parquet

The following table shows the mapping from MySQL data types to Parquet data types when data is converted and
exported to Amazon S3.

Source data typeParquet primitive typeLogical type annotationConversion notes**Numeric data types**BIGINTINT64BIGINT UNSIGNEDFIXED\_LEN\_BYTE\_ARRAY(9) DECIMAL(20,0)Parquet supports only signed types, so the mapping requires an
additional byte (8 plus 1) to store the BIGINT\_UNSIGNED
type.BITBYTE\_ARRAYDECIMALINT32DECIMAL(p,s)If the source value is less than 231,
it's stored as INT32.
INT64DECIMAL(p,s)If the source value is 231 or greater,
but less than 263, it's stored as
INT64.FIXED\_LEN\_BYTE\_ARRAY(N)DECIMAL(p,s)If the source value is 263 or greater,
it's stored as FIXED\_LEN\_BYTE\_ARRAY(N).BYTE\_ARRAYSTRINGParquet doesn't support Decimal precision greater than 38.
The Decimal value is converted to a string in a BYTE\_ARRAY type and
encoded as UTF8.DOUBLEDOUBLEFLOATDOUBLEINTINT32INT UNSIGNEDINT64MEDIUMINTINT32MEDIUMINT UNSIGNEDINT64 NUMERICINT32DECIMAL(p,s)

If the source value is less than
231, it's stored as
INT32.

INT64DECIMAL(p,s)If the source value is 231 or greater,
but less than 263, it's stored as
INT64.FIXED\_LEN\_ARRAY(N)DECIMAL(p,s)If the source value is 263 or greater,
it's stored as FIXED\_LEN\_BYTE\_ARRAY(N).BYTE\_ARRAYSTRINGParquet doesn't support Numeric precision greater than 38. This
Numeric value is converted to a string in a BYTE\_ARRAY type and
encoded as UTF8.SMALLINTINT32SMALLINT UNSIGNEDINT32TINYINTINT32TINYINT UNSIGNEDINT32INT(16, true)**String data types**BINARYBYTE\_ARRAYBLOBBYTE\_ARRAYCHARBYTE\_ARRAYENUMBYTE\_ARRAYSTRINGLINESTRINGBYTE\_ARRAYLONGBLOBBYTE\_ARRAYLONGTEXTBYTE\_ARRAYSTRINGMEDIUMBLOBBYTE\_ARRAYMEDIUMTEXTBYTE\_ARRAYSTRINGMULTILINESTRINGBYTE\_ARRAYSETBYTE\_ARRAYSTRINGTEXTBYTE\_ARRAYSTRINGTINYBLOBBYTE\_ARRAYTINYTEXTBYTE\_ARRAYSTRINGVARBINARYBYTE\_ARRAYVARCHARBYTE\_ARRAYSTRING**Date and time data types**DATEBYTE\_ARRAYSTRINGA date is converted to a string in a BYTE\_ARRAY type and encoded
as UTF8.DATETIMEINT64 TIMESTAMP\_MICROSTIMEBYTE\_ARRAYSTRINGA TIME type is converted to a string in a BYTE\_ARRAY and encoded
as UTF8.TIMESTAMPINT64 TIMESTAMP\_MICROSYEARINT32**Geometric data types**GEOMETRYBYTE\_ARRAYGEOMETRYCOLLECTIONBYTE\_ARRAYMULTIPOINTBYTE\_ARRAYMULTIPOLYGONBYTE\_ARRAYPOINTBYTE\_ARRAYPOLYGONBYTE\_ARRAY**JSON data type**JSON BYTE\_ARRAYSTRING

### PostgreSQL data type mapping to Parquet

The following table shows the mapping from PostgreSQL data types to Parquet data types when
data is converted and exported to Amazon S3.

PostgreSQL data typeParquet primitive typeLogical type annotationMapping notes**Numeric data types**BIGINTINT64BIGSERIALINT64DECIMALBYTE\_ARRAYSTRINGA DECIMAL type is converted to a string in a BYTE\_ARRAY type and
encoded as UTF8.

This conversion is to avoid complications due
to data precision and data values that are not a number
(NaN).

DOUBLE PRECISIONDOUBLEINTEGERINT32MONEYBYTE\_ARRAYSTRINGREALFLOATSERIALINT32SMALLINTINT32INT(16, true)SMALLSERIALINT32INT(16, true)**String and related data types**ARRAYBYTE\_ARRAYSTRING

An array is converted to a string and encoded as BINARY
(UTF8).

This conversion is to avoid complications due to data
precision, data values that are not a number (NaN), and time
data values.

BITBYTE\_ARRAYSTRINGBIT VARYINGBYTE\_ARRAYSTRINGBYTEABINARYCHARBYTE\_ARRAYSTRINGCHAR(N)BYTE\_ARRAYSTRINGENUMBYTE\_ARRAYSTRINGNAMEBYTE\_ARRAYSTRINGTEXTBYTE\_ARRAYSTRINGTEXT SEARCHBYTE\_ARRAYSTRINGVARCHAR(N)BYTE\_ARRAYSTRINGXMLBYTE\_ARRAYSTRING**Date and time data types**DATEBYTE\_ARRAYSTRINGINTERVALBYTE\_ARRAYSTRINGTIMEBYTE\_ARRAYSTRINGTIME WITH TIME ZONEBYTE\_ARRAYSTRINGTIMESTAMPBYTE\_ARRAYSTRINGTIMESTAMP WITH TIME ZONEBYTE\_ARRAYSTRING**Geometric data types**BOXBYTE\_ARRAYSTRINGCIRCLEBYTE\_ARRAYSTRINGLINEBYTE\_ARRAYSTRINGLINESEGMENTBYTE\_ARRAYSTRINGPATHBYTE\_ARRAYSTRINGPOINTBYTE\_ARRAYSTRINGPOLYGONBYTE\_ARRAYSTRING**JSON data types**JSONBYTE\_ARRAYSTRINGJSONBBYTE\_ARRAYSTRING**Other data types**BOOLEANBOOLEANCIDRBYTE\_ARRAYSTRING Network data typeCOMPOSITEBYTE\_ARRAYSTRINGDOMAINBYTE\_ARRAYSTRINGINETBYTE\_ARRAYSTRING Network data typeMACADDRBYTE\_ARRAYSTRINGOBJECT IDENTIFIERN/APG\_LSNBYTE\_ARRAYSTRINGRANGEBYTE\_ARRAYSTRINGUUIDBYTE\_ARRAYSTRING

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Exporting DB cluster snapshot data to Amazon S3

Setting up access to an S3
bucket

All content copied from https://docs.aws.amazon.com/.
