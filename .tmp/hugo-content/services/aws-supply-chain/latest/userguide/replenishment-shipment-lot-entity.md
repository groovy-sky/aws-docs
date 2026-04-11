# shipment\_lot

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnshipment\_lotid, product\_id, tpartner\_id, order\_id, shipment\_id, order\_line\_id, package\_id

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes

Shipment ID. Unique shipment identifier.

product\_id2

string

Yes

Product ID. Unique product identifier.

serial\_number

string

No

Unique serial number assigned to the lot. Serial numbers are often used for tracking and traceability purposes, particularly in industries where lot-level tracking is crucial.

lot\_qty

double

Yes

Quantity or number of units within the specific lot. It allows you to track the size or volume of each lot.

mfg\_date

timestamp

No

Manufacturing date.

expiry\_date

timestamp

No

Expiry date.

tpartner\_id2

string

No1

Partner that is sending the shipment. For example, shipments generated under POs, this will be vendors.

order\_id

string

No1

Order ID.

shipment\_id2

string

Yes1

Shipment ID. Unique shipment identifier.

order\_line\_id2

string

No1

Order line ID.

package\_id2

string

No1

Package ID. One shipment can have multiple packages in EDI.

source\_event\_id

string

No

ID of the event created in the source system.

source\_update\_dttm

timestamp

No

Date timestamp of the update made in the source system.

1You must enter a value. When you ingest data from SAP
or EDI, the default value for _string_ is SCN\_RESERVED\_NO\_VALUE\_PROVIDED.

2Foreign key

1Foreign key

**Foreign key (FK)**

The table below lists the columns with the associated foreign key.

ColumnCategoryFK/Data entityFK/Columnproduct\_idInboundshipmentproduct\_idtpartner\_idInboundshipmentsupplier\_tpartner\_idorder\_idInboundshipmentorder\_idshipment\_idInboundshipmentidorder\_line\_idInboundshipmentorder\_line\_idpackage\_idInboundshipmentpackage\_id

[Document Conventions](../../../../general/latest/gr/docconventions.md)

shipment\_stop\_order

Outbound fulfillment

All content copied from https://docs.aws.amazon.com/.
