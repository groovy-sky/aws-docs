# product\_hierarchy

**Primary key (PK)**

The table below lists the column names that are uniquely identified in the data entity.

NameColumnproduct\_hierarchyid

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes

Product group ID.

description

string

No

Description of the product group.

company\_id1

string

No

Company ID.

parent\_product\_group\_id1

string

No

Parent of this product group. If null, it indicates this
record is a top level product group.

creation\_date

timestamp

No

Date when the product group was created.

update\_date

timestamp

No

Date when the product group was updated.

sourcestringNoSource of data.

1Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidparent\_product\_group\_idProductproduct\_hierarchyid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

product

product\_uom

All content copied from https://docs.aws.amazon.com/.
