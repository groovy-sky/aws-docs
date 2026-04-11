# geography

**Primary key (PK)**

The table below lists the column names that are uniquely identified in the data entity.

NameColumngeographyid

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes

Geographical ID. Referred to by other entities as
geo\_id or region\_id.

description

string

No

Geographical location.

company\_id 1

string

No

Company ID.

parent\_geo\_id 1

string

No

Stores parent geographical ID for this record. If blank, this
is a top level region in the company.

address\_1

string

No

City corresponding to this geo-region.

address\_2

string

No

City corresponding to this geo-region.

address\_3

string

No

City corresponding to this geo-region.

city

string

No

Displays the city corresponding to this geo-region.

state\_prov

string

No

State corresponding to this geo-region.

postal\_code

string

No

Postal code corresponding to this
geo-region.

country

string

No

Country corresponding to this geo-region.

phone\_number

string

No

Company's contact number.

time\_zone

string

No

Company local time zone.

source

string

No

Source of data.

source\_event\_id

string

No

ID of the event created in the source system.

source\_update\_dttm

timestamp

No

Date time stamp of the update made in the source
system.

1 Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidparent\_geo\_idOrganizationgeographyid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

company

trading\_partner

All content copied from https://docs.aws.amazon.com/.
