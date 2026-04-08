# Considerations for DB cluster exports

Use the following sections to learn about the limitations, file naming conventions, and data conversion and storage when exporting DB cluster data to Amazon S3.

###### Topics

- [Limitations](#export-cluster-data.Limits)

- [File naming convention](#export-cluster-data.FileNames)

- [Data conversion and storage format](#export-cluster-data.data-types)

## Limitations

Exporting DB cluster data to Amazon S3 has the following limitations:

- You can't run multiple export tasks for the same DB cluster simultaneously. This applies to both full and partial
exports.

- You can have up to five concurrent DB snapshot export tasks in progress per AWS account.

- Aurora Serverless v1 DB clusters don't support exports to S3.

- Aurora MySQL and Aurora PostgreSQL support exports to S3 only for the provisioned engine mode.

- Exports to S3 don't support S3 prefixes containing a colon (:).

- The following characters in the S3 file path are converted to underscores (\_) during export:

```nohighlight

\ ` " (space)
```

- If a database, schema, or table has characters in its name other than the following, partial export isn't
supported. However, you can export the entire DB cluster.

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

- If the data contains a large object, such as a BLOB or CLOB, that is close to or greater than 500 MB, then the export
fails.

- If a table contains a large row that is close to or greater than 2 GB, then the table is skipped during export.

- For partial exports, the `ExportOnly` list has a maximum size of
200 KB.

- We strongly recommend that you use a unique name for each export task. If you don't use a unique task name, you might
receive the following error message:

**`ExportTaskAlreadyExistsFault: An error occurred (ExportTaskAlreadyExists) when calling the StartExportTask**
**operation: The export task with the ID xxxxx already exists.`**

- Because some tables might be skipped, we recommend that you verify row and table counts in the data after
export.

## File naming convention

Exported data for specific tables is stored in the format
`base_prefix/files`, where the base prefix is
the following:

```nohighlight

export_identifier/database_name/schema_name.table_name/
```

For example:

```nohighlight

export-1234567890123-459/rdststcluster/mycluster.DataInsert_7ADB5D19965123A2/
```

Output files use the following naming convention, where `partition_index` is alphanumeric:

```nohighlight

partition_index/part-00000-random_uuid.format-based_extension
```

For example:

```nohighlight

1/part-00000-c5a881bb-58ff-4ee6-1111-b41ecff340a3-c000.gz.parquet
    a/part-00000-d7a881cc-88cc-5ab7-2222-c41ecab340a4-c000.gz.parquet

```

The file naming convention is subject to change. Therefore, when reading target tables, we recommend that you read everything
inside the base prefix for the table.

## Data conversion and storage format

When you export a DB cluster to an Amazon S3 bucket, Amazon Aurora converts, exports, and stores data in the Parquet
format. For more information, see [Data conversion when exporting to an Amazon S3 bucket](aurora-export-snapshot-considerations.md#aurora-export-snapshot.data-types).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Exporting DB cluster data to Amazon S3

Setting up access to an S3 bucket

All content copied from https://docs.aws.amazon.com/.
