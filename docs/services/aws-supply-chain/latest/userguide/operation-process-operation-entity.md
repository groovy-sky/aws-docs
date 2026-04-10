# process\_operation

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnprocess\_operationprocess\_operation\_id, process\_id

The table below lists the column names supported by the _process\_operation_ data entity:

ColumnData typeRequiredDescription

process\_operation\_id

string

Yes

Type of process operation.

process\_id1

string

Yes

Process ID. For example, process, work order, or maintenance order.

company\_id1

string

No

Company ID.

type

string

No

Type of operation within the process. For example, open machine.

site\_location

string

No

Name of the location or section in site or plant.

status

string

No

Status of the process.

operation\_name

string

No

Name of the operation.

operation\_sequence

string

No

Sequence of the operation within the process.

planned\_start\_dttm

timestamp

No

Planned start date-time of operation.

planned\_end\_dttm

timestamp

No

Planned end date-time of operation.

1Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columnprocess\_idOperationprocess\_headerprocess\_idcompany\_idOrganizationcompanyid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

process\_header

process\_product

All content copied from https://docs.aws.amazon.com/.
