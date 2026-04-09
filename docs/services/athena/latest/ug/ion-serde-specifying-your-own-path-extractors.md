# Specify your own path extractors

If your Amazon Ion fields do not map neatly to Hive columns, you can specify your own
path extractors. In the `WITH SERDEPROPERTIES` clause of your `CREATE
                TABLE` statement, use the following syntax.

```sql

WITH SERDEPROPERTIES (
   "ion.path_extractor.case_sensitive" = "<Boolean>",
   "ion.<column_name>.path_extractor" = "<path_extractor_expression>"
)
```

###### Note

By default, path extractors are case insensitive. To override this setting, set
the [ion.path\_extractor.case\_sensitive](ion-serde-using-ion-serde-properties.md#ioncase) SerDe property to
`true`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use Athena generated path extractors

Use search paths in path extractors

All content copied from https://docs.aws.amazon.com/.
