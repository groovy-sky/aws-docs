# SHOW TBLPROPERTIES

Lists table properties for the named table.

## Synopsis

```sql

SHOW TBLPROPERTIES table_name [('property_name')]
```

## Parameters

**\[('property\_name')\]**

If included, only the value of the property named `property_name` is listed.

## Examples

```sql

SHOW TBLPROPERTIES orders;
```

```sql

SHOW TBLPROPERTIES orders('comment');
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SHOW TABLES

SHOW VIEWS
