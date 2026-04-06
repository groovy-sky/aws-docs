# Choose a SerDe for your data

The following table lists the data formats supported in Athena and their corresponding
SerDe libraries.

Supported data formats and SerDesData formatDescriptionSerDe types supported in AthenaAmazon IonAmazon Ion is a richly-typed, self-describing data format that is a
superset of JSON, developed and open-sourced by Amazon.Use the [Amazon Ion Hive SerDe](https://docs.aws.amazon.com/athena/latest/ug/ion-serde.html).

Apache Avro

A format for storing data in Hadoop that uses JSON-based schemas for
record values.

Use the [Avro SerDe](https://docs.aws.amazon.com/athena/latest/ug/avro-serde.html).

Apache Parquet

A format for columnar storage of data in Hadoop.

Use the [Parquet SerDe](https://docs.aws.amazon.com/athena/latest/ug/parquet-serde.html)
and SNAPPY compression.

Apache WebServer logs

A format for storing logs in Apache WebServer.

Use the [Grok SerDe](https://docs.aws.amazon.com/athena/latest/ug/grok-serde.html) or
[Regex SerDe](regex-serde.md).

CloudTrail logs

A format for storing logs in CloudTrail.

- Use the [Hive JSON SerDe](https://docs.aws.amazon.com/athena/latest/ug/hive-json-serde.html). For more information, see
[Query AWS CloudTrail logs](cloudtrail-logs.md).

CSV (Comma-Separated Values)

For data in CSV, each line represents a data record, and each record
consists of one or more fields, separated by commas.

- Use the [Lazy Simple SerDe for CSV, TSV, and custom-delimited files](lazy-simple-serde.md) if your data does not
include values enclosed in quotes or if it uses the
`java.sql.Timestamp` format.

- Use the [Open CSV SerDe for processing CSV](https://docs.aws.amazon.com/athena/latest/ug/csv-serde.html)
when your data includes quotes in values or uses the UNIX
numeric format for `TIMESTAMP` (for example,
`1564610311`).

Custom-Delimited

For data in this format, each line represents a data record, and
records are separated by a custom single-character delimiter.

Use the [Lazy Simple SerDe for CSV, TSV, and custom-delimited files](lazy-simple-serde.md) and specify a custom
single-character delimiter.

JSON (JavaScript Object Notation)

For JSON data, each line represents a data record, and each record
consists of attribute-value pairs and arrays, separated by
commas.

- Use the [Hive JSON SerDe](https://docs.aws.amazon.com/athena/latest/ug/hive-json-serde.html).

- Use the [OpenX JSON SerDe](https://docs.aws.amazon.com/athena/latest/ug/openx-json-serde.html).

Logstash logs

A format for storing logs in Logstash.

Use the [Grok SerDe](https://docs.aws.amazon.com/athena/latest/ug/grok-serde.html).

ORC (Optimized Row Columnar)

A format for optimized columnar storage of Hive data.

Use the [ORC SerDe](https://docs.aws.amazon.com/athena/latest/ug/orc-serde.html) and ZLIB
compression.

TSV (Tab-Separated Values)

For data in TSV, each line represents a data record, and each record
consists of one or more fields, separated by tabs.

Use the [Lazy Simple SerDe for CSV, TSV, and custom-delimited files](lazy-simple-serde.md) and specify the separator
character as `FIELDS TERMINATED BY '\t'`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Use SerDes

Use a SerDe to create a table
