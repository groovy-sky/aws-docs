---
title: "ALTER TABLE UNSET TBLPROPERTIES"
---

# ALTER TABLE UNSET TBLPROPERTIES
<a name="querying-iceberg-alter-table-unset-properties"></a>

Drops existing properties from an Iceberg table.

## Synopsis
<a name="querying-iceberg-alter-table-unset-properties-synopsis"></a>

```
ALTER TABLE [{{db_name}}.]{{table_name}} UNSET TBLPROPERTIES ('{{property_name}}' [ , ... ])
```

## Example
<a name="querying-iceberg-alter-table-unset-properties-example"></a>

```
ALTER TABLE iceberg_table UNSET TBLPROPERTIES ('write_compression')
```

The following example removes the `write.data.path` property from an Iceberg table.

```
ALTER TABLE iceberg_table UNSET TBLPROPERTIES ('write.data.path')
```

All content copied from https://docs.aws.amazon.com/.
