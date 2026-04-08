# Data conversion when exporting to an Amazon S3 bucket for Amazon RDS

When you export a DB snapshot to an Amazon S3 bucket, Amazon RDS converts data to,
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
After an export task completes, Amazon RDS notifies you if any string conversion occurred.
The underlying data exported is always the same as the data from the source. However,
due to the encoding difference in UTF-8, some characters might appear different from the
source when read in tools such as Athena.

For more information, see [Parquet\
logical type definitions](https://github.com/apache/parquet-format/blob/master/LogicalTypes.md) in the Parquet documentation.

###### Topics

- [MySQL and MariaDB data type mapping to Parquet](#USER_ExportSnapshot.data-types.MySQL)

- [PostgreSQL data type mapping to Parquet](#USER_ExportSnapshot.data-types.PostgreSQL)

## MySQL and MariaDB data type mapping to Parquet

The following table shows the mapping from MySQL and MariaDB data types to Parquet data types when data is converted and
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

## PostgreSQL data type mapping to Parquet

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

File naming conventions

Using AWS Backup

All content copied from https://docs.aws.amazon.com/.
