# shipment\_stop

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnshipment\_stopshipment\_stop\_id, shipment\_id

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

shipment\_stop\_id

string

Yes

Shipment stop ID.

shipment\_id1

string

Yes

Shipment ID.

sequence

int

No

Sequence of the shipment.

company\_id1

string

No

Company ID.

site\_id1

string

No

Site ID.

planned\_arrival\_start\_dttm

timestamp

No

Planned start date and time for the shipment arrival.

planned\_arrival\_end\_dttm

timestamp

No

Planned end date and time for the shipments arrival.

planned\_departure\_start\_dttm

timestamp

No

Planned start date and time for the shipment departure.

planned\_departure\_end\_dttm

timestamp

No

Planned end date and time for the shipment departure.

actual\_arrival\_start\_dttm

timestamp

No

Actual start date and time for the shipment arrival.

actual\_arrival\_end\_dttm

timestamp

No

Actual end date and time for the shipments arrival.

actual\_departure\_start\_dttm

timestamp

No

Actual start date and time for the shipment departure.

actual\_departure\_end\_dttm

timestamp

No

Actual end date and time for the shipment departure.

appointment\_number

###### Note

AWS Supply Chain web application will display this column as _appointment\_number_.

string

No

Appointment number.

delivery\_number

string

No

Delivery number of the shipment.

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

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidsite\_idNetworksiteidshipment\_idInboundshipmentid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

shipment

shipment\_stop\_order

All content copied from https://docs.aws.amazon.com/.
