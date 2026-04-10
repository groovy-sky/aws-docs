# outbound\_shipment

**Primary key (PK)**

The table below lists the colum names that are uniquely identified in the data entity.

NameColumnoutbound\_shipmentid, cust\_order\_id, cust\_order\_line\_id, product\_id

The table below lists the column names supported by the data entity:

ColumnData typeRequiredDescription

id

string

Yes 1

Outbound shipment ID.

company\_id2

string

No

Company ID.

cust\_order\_id2

string

Yes 1

Customer order ID.

cust\_order\_line\_id2

string

Yes 1

Customer
order line ID.

product\_id2

string

Yes 1

Product ID.

shipped\_qty

double

No

Shipment quantity.

cust\_shipment\_status

string

No

Status of the shipment, for example, canceled, open, closed,
or delivered.

expected\_ship\_date

timestamp

No

Date product was expected to ship from the company
location.

actual\_ship\_date

timestamp

No

Date product was actually shipped from the company
location.

from\_site\_id2

string

No

Site ID where the product is shipped from.

to\_site\_id2

string

No

Destination site ID for outbound shipments.

expected\_delivery\_date

timestamp

No

Expected delivery date of the products to the customer.

actual\_delivery\_date

timestamp

No

Displays when the product was actually delivered to the
customer.

shipping\_cost

double

No

Final shipping cost.

tracking\_number

string

No

Tracking number associated with the shipment.

bill\_weight

double

No

Shipped weight of product used for billing.

sap\_2lis\_08trtlp\_\_vbeln

string

No

Delivery number. Predicate key for SAP mapping. Upsert key for 2LIS\_12\_VCITM.

sap\_2lis\_08trtlp\_\_posnr

string

No

Delivery item number. Predicate key for SAP mapping. Upsert key for 2LIS\_12\_VCITM.

sap\_2lis\_08trtlp\_\_tknum

string

No

Shipment item number. Predicate key for SAP mapping. Upsert key for 2LIS\_08TRTK.

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

tpartner\_id

string

No

Unique identifier for a trading partner.

service\_level

string

No

Focuses on the quality and speed of the shipment. For example, Standard, next day, two-day, expedited, and so on.

1You must enter a value. When you ingest data from SAP
or EDI, the default value for _string_ is SCN\_RESERVED\_NO\_
VALUE\_PROVIDED.

2Foreign key

**Foreign key (FK)**

The table below lists the column names with the associated data entity and category:

ColumnCategoryFK/Data entityFK/Columncompany\_idOrganizationcompanyidproduct\_idProductproductidcust\_order\_line\_idOutboundFulfillmentoutbound\_order\_lineidcust\_order\_idOutboundFulfillmentoutbound\_order\_linecust\_order\_idfrom\_site\_id, to\_site\_idNetworksiteidtpartner\_idOrganizationtrading\_partnerid

[Document Conventions](../../../../general/latest/gr/docconventions.md)

outbound\_order\_line

Cost management

All content copied from https://docs.aws.amazon.com/.
