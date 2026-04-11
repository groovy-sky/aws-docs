# Find array lengths

The `cardinality` function returns the length of an array, as in this example:

```sql

SELECT cardinality(ARRAY[1,2,3,4]) AS item_count
```

This query returns:

```

+------------+
| item_count |
+------------+
| 4          |
+------------+
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Convert array data types

Access array elements

All content copied from https://docs.aws.amazon.com/.
