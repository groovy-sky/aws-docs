# Data types in Amazon Athena

When you run `CREATE TABLE`, you specify column names and the data type that
each column can contain. The tables that you create are stored in the AWS Glue Data Catalog.

To facilitate interoperability with other query engines, Athena uses [Apache\
Hive](https://cwiki.apache.org/confluence/display/Hive/LanguageManual+Types) data type names for DDL statements like `CREATE TABLE`. For DML
queries like `SELECT`, `CTAS`, and `INSERT INTO`, Athena
uses [Trino](https://trino.io/docs/current/language/types.html) data type
names. The following table shows the data types supported in Athena. Where DDL and DML types
differ in terms of name, availability, or syntax, they are shown in separate columns.

DDLDMLDescriptionBOOLEANValues are `true` and `false`.TINYINTAn 8-bit signed integer in two's complement format, with a minimum value
of -27 and a maximum value of
27-1.SMALLINTA 16-bit signed integer in two's complement format, with a minimum value
of -215 and a maximum value of
215-1.INT, INTEGERA 32-bit signed value in two's complement format, with a minimum value of
 -231 and a maximum value of
231-1.BIGINT A 64-bit signed integer in two's complement format, with a minimum value
of -263 and a maximum value of
263-1.FLOATREALA 32-bit signed single-precision floating point number. The range is
1.40129846432481707e-45 to 3.40282346638528860e+38, positive or negative.
Follows the IEEE Standard for Floating-Point Arithmetic (IEEE 754).DOUBLEA 64-bit signed double-precision floating point number. The range is
4.94065645841246544e-324d to 1.79769313486231570e+308d, positive or
negative. Follows the IEEE Standard for Floating-Point Arithmetic (IEEE
754).DECIMAL( `precision`,
`scale`)`precision` is the total number of
digits. `scale` (optional) is the
number of digits in fractional part with a default of 0. For example, use
these type definitions: `decimal(11,5)`,
`decimal(15)`. The maximum value for
`precision` is 38, and the maximum value for
`scale` is 38.CHAR,
CHAR( `length`)

Fixed length character data, with a specified length between 1 and
255, such as char(10). If `length` is
specified, strings are truncated at the specified length when read. If
the underlying data string is longer, the underlying data string remains
unchanged.

For more information, see [CHAR Hive data type](https://cwiki.apache.org/confluence/display/Hive/LanguageManual+Types).

STRINGVARCHARVariable length character data.VARCHAR( `length`)Variable length character data with a maximum read length. Strings are
truncated at the specified length when read. If the underlying data string
is longer, the underlying data string remains unchanged.BINARYVARBINARYVariable length binary data.Not availableTIMEA time of day with millisecond precision.Not availableTIME( `precision`)A time of day with a specific precision. `TIME(3)` is
equivalent to `TIME`.Not availableTIME WITH TIME ZONEA time of day in a time zone. Time zones should be specified as offsets
from UTC.DATEA calendar date with year, month, and day.TIMESTAMPTIMESTAMP, TIMESTAMP WITHOUT TIME ZONEA calendar date and time of day with millisecond precision.Not availableTIMESTAMP( `precision`),
TIMESTAMP( `precision`) WITHOUT TIME
ZONEA calendar date and time of day with a specific precision.
`TIMESTAMP(3)` is equivalent to
`TIMESTAMP`.Not availableTIMESTAMP WITH TIME ZONEA calendar date and time of day in a time zone. Time zones can be
specified as offsets from UTC, as IANA time zone names, or using UTC, UT, Z,
or GMT.Not availableTIMESTAMP( `precision`) WITH TIME ZONEA calendar date and time of day with a specific precision, in a time
zone.Not availableINTERVAL YEAR TO MONTHAn interval of one or more whole monthsNot availableINTERVAL DAY TO SECONDAn interval of one or more seconds, minutes, hours, or daysARRAY< `element_type` >ARRAY\[ `element_type`\]An array of values. All values must be of the same data type.MAP< `key_type`,
`value_type` >MAP( `key_type`,
`value_type`)A map where values can be looked up by key. All keys must be of the same
data type, and all values must be of the same data type. For example,
`map<string, integer>`.STRUCT< `field_name_1`: `field_type_1`,
`field_name_2`: `field_type_2`,
…>ROW( `field_name_1` `field_type_1`,
`field_name_2` `field_type_2`, …)A data structure with named fields and their values.Not availableJSONJSON value type, which can be a JSON object, a JSON array, a JSON number,
a JSON string, `true`, `false` or
`null`.Not availableUUIDA UUID (Universally Unique IDentifier).Not availableIPADDRESSAn IPv4 or IPv6 address.Not available[HyperLogLog](https://trino.io/docs/current/language/types.html)These data types support approximate function internals. For
more information about each type, visit the link to the corresponding entry
in the Trino documentation.[P4HyperLogLog](https://trino.io/docs/current/language/types.html)[SetDigest](https://trino.io/docs/current/language/types.html)[QDigest](https://trino.io/docs/current/language/types.html)[TDigest](https://trino.io/docs/current/language/types.html)

###### Topics

- [Data type examples](data-types-examples.md)

- [Considerations for data types](data-types-considerations.md)

- [Work with timestamp data](data-types-timestamps.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SQL reference for Athena

Data type examples

All content copied from https://docs.aws.amazon.com/.
