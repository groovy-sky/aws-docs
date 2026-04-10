# reference\_field

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnreference\_fieldobject\_name, object\_field, object\_field\_value, object\_field\_desc

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

company\_id2

string

No

Company ID.

object\_name

string

Yes1

For example, sites, or transportation lanes.

object\_field

string

Yes1

For example, site\_type, or trans\_mode.

object\_field\_value

string

Yes1

For example, site\_type:01, or trans\_mode:01.

object\_field\_desc

string

Yes1

For example, site\_type:01:DC, or trans\_mode:01:Surface.

1You must enter a value. When you ingest data from SAP
or EDI, the default value for _string_ is SCN\_RESERVED\_NO\_VALUE\_PROVIDED.

2Foreign key

**Foreign key (FK)**

The table below lists the column names with the associated data entity and category:

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Reference

calendar

All content copied from https://docs.aws.amazon.com/.
