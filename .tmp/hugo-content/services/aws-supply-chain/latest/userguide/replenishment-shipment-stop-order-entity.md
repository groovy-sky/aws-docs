# shipment\_stop\_order

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnshipment\_stop\_ordershipment\_stop\_order\_id, shipment\_stop\_id, shipment\_id

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

shipment\_stop\_order\_id

string

Yes

Shipment stop order ID.

shipment\_stop\_id1

string

Yes

Shipment stop ID.

shipment\_id1

string

Yes

Shipment ID.

company\_id1

string

No

Company ID.

site\_id1

string

No

Site ID.

inbound\_order\_id1

string

No

Inbound order ID.

inbound\_order\_line\_id1

string

No

Inbound order line ID.

inbound\_order\_line\_schedule\_id1

string

No

Inbound order line schedule ID.

action

string

No

Pickup or drop off shipment.

quantity

double

No

Quantity associated with action and order.

quantity\_uom

string

No

Quantity UOM of the shipment.

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

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidsite\_idNetworksiteidshipment\_idInboundshipmentidshipment\_stop\_idInboundshipment\_stopshipment\_stop\_idinbound\_order\_idInboundinbound\_order\_lineorder\_idinbound\_order\_line\_idInboundinbound\_order\_lineidinbound\_order\_line\_schedule\_idInboundinbound\_order\_line\_scheduleid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

shipment\_stop

shipment\_lot

All content copied from https://docs.aws.amazon.com/.
