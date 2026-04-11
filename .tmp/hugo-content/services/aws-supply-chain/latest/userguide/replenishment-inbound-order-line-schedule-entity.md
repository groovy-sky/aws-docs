# inbound\_order\_line\_schedule

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumninbound\_order\_line\_scheduleid, order\_id, order\_line\_id, product\_id

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes 1

Order line ID. The value must be unique.

order\_id2

string

Yes 1

ID of parent order.

order\_line\_id2

string

Yes

ID of parent order line.

company\_id2

string

No

Company ID.

status

string

No

Status of line, for example, submitted, or confirmed. The
following are the reserved values for AWS Supply Chain.

- Can celled - Populated in SAP mapping. Also used for
deleted.

- Open - Not populated in SAP mapping.

- Closed - Not populated in SAP mapping.

- InTransit - Not populated in SAP mapping.

- Confirmed - Not populated in SAP mapping.

###### Note

Null is also an accepted value, or you can enter your own
value.

schedule\_creation\_date

timestamp

No

Schedule creation date.

product\_id2

string

Yes 1

Product ID.

external\_line\_number

string

No

External line number.

expected\_delivery\_date

timestamp

No

Expected delivery date of the products.

confirmation\_date

timestamp

No

Date and time when the vendor confirmed the order line
schedule, or order.

goods\_issue\_date

timestamp

No

Date and time when the material was available at origin to
ship.

material\_availability\_date

timestamp

No

Date and time when the material was available at origin to
ship.

ship\_date

timestamp

No

Date and time when vendor will ship products in this
order-line-schedule.

delivery\_date

timestamp

No

Date and time when the vendor can deliver products in this
schedule.

quantity\_submitted

double

No

Quantity submitted to vendor (POs) or for transfer.

quantity\_confirmed

double

No

Quantity confirmed by the vendor.

quantity\_received

double

No

Quantity received into inventory at the destination.

sap\_lips\_\_vbeln

string

No

Delivery Number. Predicate key for SAP mapping. Upsert key for VTTP

sap\_vttp\_\_tknum

string

No

Shipment Number. Predicate key for SAP mapping. Upsert key for VTTK

source\_event\_id

string

No

ID of the event created in the source system.

source\_update\_dttm

timestamp

No

Date time stamp of the update made in the source
system.

1You must enter a value. When you ingest data from SAP
or EDI, the default value for _string_ is SCN\_RESERVED\_NO\_VALUE\_PROVIDED.

2Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidproduct\_idProductproductidorder\_idInboundinbound\_orderidorder\_line\_idInboundinbound\_order\_lineid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

inbound\_order\_line

shipment

All content copied from https://docs.aws.amazon.com/.
