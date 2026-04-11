# trading\_partner\_poc

**Primary key (PK)**

The table below lists the column names that are uniquely identified in the data entity.

NameColumntrading\_partner\_poctpartner\_id, email

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

tpartner\_id 1

string

Yes

Partner ID. Referred to by other entities as
tpartner\_id unless explicitly stated otherwise.

email

string

Yes

Partner's email ID.

poc\_first\_name

string

No

Partner's first name.

poc\_last\_name

string

No

Partner's last name.

poc\_org\_unit\_name

string

No

Name of the team or internal organizational unit.

poc\_org\_unit\_description

string

No

AWS profile or description of the team's role in an organization to be shared with the customer to describe their team.

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

1Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columntpartner\_idOrganizationtrading\_partnerid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

trading\_partner

Product

All content copied from https://docs.aws.amazon.com/.
