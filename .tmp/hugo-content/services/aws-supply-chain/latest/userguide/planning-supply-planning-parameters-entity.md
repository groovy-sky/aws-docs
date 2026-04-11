# supply\_planning\_parameters

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnsupply\_planning\_parametersproduct\_id, product\_group\_id, site\_id, eff\_start\_date, eff\_end\_date, connection\_id

The table below lists the column names supported by the _supply\_planning\_parameters_ data entity:

ColumnData typeRequiredDescription

product\_id1

string

Yes

ID of product

product\_group\_id1

string

Yes

For future Use. Please populate SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.

site\_id1

string

Yes

For future Use. Please populate SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.

planner\_name

string

No

name of the supply planner who manages a product or a product group

demand\_time\_fence\_days

int

No

For future Use.

forecast\_consumption\_backward\_days

int

No

For future Use

forecast\_consumption\_forward\_days

int

No

For future Use.

eff\_start\_date

timestamp

Yes

effective start date time

eff\_end\_date

timestamp

Yes

effective end date time

connection\_id

string

Yes

Unique identifier for the data source (i.e. connection). Auto populated by ASC.

1Foreign key

**Foreign key (FK)**

The table below lists the column names with the associated data entity and category:

ColumnCategoryFK/Data entityFK/Columnproduct\_idProductproductidproduct\_group\_idProductproduct\_hierarchyidsite\_idNetworksiteid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

reservation

Operation

All content copied from https://docs.aws.amazon.com/.
