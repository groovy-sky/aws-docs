# reservation

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnreservationreservation\_id, reservation\_detail\_id

The table below lists the column names supported by the _reservation_ data entity:

ColumnData typeRequiredDescription

reservation\_id

string

Yes

Reservation ID.

reservation\_detail\_id

string

Yes

Reservation detail ID.

reservation\_type

string

No

Type of reservation. For example, procurement or build-to-stock.

company\_id1

string

No

Company ID.

status

string

No

Status of the reservation.

product\_id1

string

No

Product ID.

site\_id1

string

No

Site ID.

quantity

double

No

Reservation quantity.

quantity\_uom

string

No

Quantity UOM associated with reservation.

reservation\_date

timestamp

No

Date when the reservation is generated.

is\_deleted

string

No

Yes or No indicator to indicate whether the reservation is deleted or not.

requisition\_id1

string

No

Source object identifier reference to inbound order type.

requisition\_line\_id1

string

No

Source object identifier reference to inbound order line.

rfq\_id1

string

No

Source object identifier reference to inbound order type RFQ.

rfq\_line\_id1

string

No

Source object identifier reference to inbound order line of type RFQ.

order\_id1

string

No

Source object identifier reference to inbound order.

order\_line\_id1

string

No

Source object identifier reference to inbound order line.

order\_line\_schedule\_id1

string

No

Source object identifier reference to inbound order line schedule.

stock\_transfer\_1\_order\_id

string

No

Stock transfer order ID.

stock\_transfer\_1\_order\_line\_id

string

No

Stock transfer order line ID.

stock\_transfer\_2\_order\_id

string

No

Stock transfer order ID.

stock\_transfer\_2\_order\_line\_id

string

No

Stock transfer order line ID.

source\_update\_dttm

timestamp

No

Date time stamp of the update made in the source system.

source\_event\_id

string

No

ID of the event created in the source system.

source

string

No

Source of data.

flex\_1

string

No

Reservation flexible field 1

flex\_2

string

No

Reservation flexible field 2

flex\_3

string

No

Reservation flexible field 3

flex\_4

string

No

Reservation flexible field 4

flex\_5

string

No

Reservation flexible field 5

1Foreign key

**Foreign key (FK)**

The table below lists the column names with the associated data entity and category:

ColumnCategoryFK/Data entityFK/Columnsite\_idNetworksiteidcompany\_idOrganizationcompanyidproduct\_idProductproductidrequisition\_id, rfq\_idInboundinbound\_order\_lineorder\_idrequisition\_line\_id, rfq\_line\_idInboundinbound\_order\_lineidorder\_line\_schedule\_idInboundinbound\_order\_line\_scheduleid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

sourcing\_schedule\_details

supply\_planning\_parameters

All content copied from https://docs.aws.amazon.com/.
