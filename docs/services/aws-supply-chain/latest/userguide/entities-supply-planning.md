# Supply Planning

The table below list the data entities and columns used by Supply Planning.

###### Note

**How to read the table:**

- **Required** – The column name is mandatory in your dataset and you must populate the column name with values.

- **Optional** – The column name is optional. For enhanced feature output, it is recommended to add the column name with values.

- **Not required** – Data entity not required.

Data entityColumnIs the column used for Auto Replenishment?Is the column used for Manufacturing Plan?

[site](network-site-entity.md)

id

RequiredRequired

description

RequiredRequired

geo\_id

Required - Without this field, filters cannot group sites by category
such as region, country, state, zip code and so on.Required - Without this field, filters cannot group sites by category
such as region, country, state, zip code and so on.

site\_type

NANA

company\_id

OptionalOptional

latitude

NANA

longitude

NANA

is\_active

Required - Identifies if a site needs to be considered for planning.
Note, set the value to _False_ if a site should not
to be considered. If the field is kept blank or null, the site will be
considered.Required - Identifies if a site needs to be considered for planning.
Note, set the value to _False_ if a site should not
to be considered. If the field is kept blank or null, the site will be
considered.

open\_date

NANA

end\_date

NANA

[transportation\_lane](network-transporation-lane-entity.md)

id

RequiredRequired

from\_site\_id

RequiredRequired

to\_site\_id

RequiredRequired

product\_group\_id

RequiredRequired

transit\_time

RequiredRequired

time\_uom

Required - Supported values include Day.Required - Supported values include Day.

distance

Not requiredNot required

distance\_uom

Not requiredNot required

eff\_start\_date

OptionalOptional

eff\_end\_date

OptionalOptional

product\_id

OptionalOptional

emissions\_per\_unit

Not requiredNot required

emissions\_per\_weight

Not requiredNot required

company\_id

OptionalOptional

from\_geo\_id

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

to\_geo\_id

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

carrier\_tpartner\_id

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

service\_type

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

trans\_mode

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

cost\_per\_unit

OptionalOptional

cost\_currency

OptionalOptional

[product](product-product-entity.md)

id

RequiredRequired

description

RequiredRequired

product\_group\_id

Required - Without this field, filters cannot group by product
category such as dairy, clothes, and so on.Required - Without this field, filters cannot group by product
category such as dairy, clothes, and so on.

is\_deleted

Required - Identifies if a product needs to be considered for
planning. Set the field to _False_ to consider this
product and _True_ to not consider the product. If
this field is left blank or null, then the value will be defaulted to
_True_.Required - Identifies if a product needs to be considered for
planning. Set the field to _False_ to consider this
product and _True_ to not consider the product. If
this field is left blank or null, then the value will be defaulted to
_True_.

product\_type

Not requiredNot required

parent\_product\_id

OptionalOptional

base\_uom

OptionalOptional

unit\_cost

OptionalOptional

unit\_price

OptionalOptional

[product\_hierarchy](product-hierarchy-entity.md)

id

RequiredRequired

description

Required – This field is used by filters to group by a product
category such as dairy, clothes, and so on.Required – This field is used by filters to group by a product
category such as dairy, clothes, and so on.

parent\_product\_group\_id

Optional – This field is used by filters to support multiple
product category hierarchy such as dairy, full fat milk and so
on.Optional – This field is used by filters to support multiple
product category hierarchy such as dairy, full fat milk and so
on.

[geography](organization-geography-entity.md)

id

RequiredRequired

description

RequiredRequired

parent\_geo\_id

Optional – This field is used by filters to support multiple
location hierarchy such as USA → USA-EAST.Optional – This field is used by filters to support multiple
location hierarchy such as USA → USA-EAST.

[trading\_partner](organization-trading-partner-entity.md)

id

RequiredRequired

description

OptionalOptional

country

OptionalOptional

eff\_start\_date

Required – You must enter a value for eff\_start\_date and
eff\_end\_date. If you don't have a value, enter `1900-01-01
                                    00:00:00` for eff\_start\_date, and
`9999-12-31 23:59:59` for
eff\_end\_date.

Required – You must enter a value for eff\_start\_date and
eff\_end\_date. If you don't have a value, enter `1900-01-01
                                    00:00:00` for eff\_start\_date, and
`9999-12-31 23:59:59` for
eff\_end\_date.

eff\_end\_date

Required – You must enter a value for eff\_start\_date and
eff\_end\_date. If you don't have a value, enter `1900-01-01
                                    00:00:00` for eff\_start\_date, and
`9999-12-31 23:59:59` for
eff\_end\_date.

Required – You must enter a value for eff\_start\_date and
eff\_end\_date. If you don't have a value, enter `1900-01-01
                                    00:00:00` for eff\_start\_date, and
`9999-12-31 23:59:59` for
eff\_end\_date.

time\_zone

OptionalOptional

is\_active

OptionalOptional

tpartner\_type

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

geo\_id

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

[inbound\_order](replenishment-inbound-order-entity.md)

id

RequiredRequired

order\_type

RequiredRequired

order\_status

Not requiredNot required

to\_site\_id

RequiredRequired

submitted\_date

OptionalOptional

tpartner\_id

Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required. When you ingest data from SAP or EDI, the default value for
string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the
Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

[inbound\_order\_line](replenishment-inbound-order-line-entity.md)

id

RequiredRequired

order\_id

RequiredRequired

order\_type

RequiredRequired

status

Not requiredNot required

product\_id

RequiredRequired

to\_site\_id

RequiredRequired

from\_site\_id

Not requiredNot required

quantity\_submitted

Required – You must set one quantity field.Required – You must set one quantity field.

quantity\_confirmed

Optional – You must set one quantity field.Optional – You must set one quantity field.

quantity\_received

Optional – You must set one quantity field.Optional – You must set one quantity field.

expected\_delivery\_date

RequiredRequired

submitted\_date

Not requiredNot required

incoterm

Not requiredNot required

company\_id

OptionalOptional

tpartner\_id

Required – This field is required for successful
ingestion.Required – This field is required for successful
ingestion.

quantity\_uom

Not requiredNot required

reservation\_id

Not requiredNot required

reference\_object\_type

Optional – This field is used for associating purchase order
requests to purchase orders to track plan to PO conversion in the
ERP.Optional – This field is used for associating purchase order
requests to purchase orders to track plan to PO conversion in the
ERP.

reference\_object\_id

Optional – This field is used for associating purchase order
requests to purchase orders to track plan to PO conversion in the
ERP.Optional – This field is used for associating purchase order
requests to purchase orders to track plan to PO conversion in the
ERP.

[inv\_policy](planning-inv-policy-entity.md)

site\_id

RequiredRequired

id

RequiredRequired

dest\_geo\_id

RequiredRequired

product\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

product\_group\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

company\_id

OptionalOptional

ss\_policy

Required – The accepted values for this field are abs\_level,
doc\_dem, doc\_fcst, and sl.Required – The accepted values for this field are abs\_level,
doc\_dem, doc\_fcst, and sl.

target\_inventory\_qty

Required – This field is required when ss\_policy is set to
abs\_level.Required – This field is required when ss\_policy is set to
abs\_level.

target\_doc\_limit

Required – This field is required when ss\_policy is set to
doc\_dem or doc\_fcst.Required – This field is required when ss\_policy is set to
doc\_dem or doc\_fcst.

target\_sl

Required – This field is required when ss\_policy is set to
sl.Required – This field is required when ss\_policy is set to
sl.

[sourcing\_rules](planning-sourcing-rules-entity.md)

sourcing\_rule\_id

RequiredRequired

company\_id

OptionalOptional

product\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

product\_group\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

from\_site\_id

Optional – This field is required for sourcing\_rule types
transfer.Optional – This field is required for sourcing\_rule types
transfer.

to\_site\_id

RequiredRequired

sourcing\_rule\_type

Required – The allowed values for this field are transfer,
buy, and manufacture.Required – The allowed values for this field are transfer,
buy, and manufacture. Only lower case is allowed.

tpartner\_id

Optional – This field is required for sourcing\_rule types
buy.Optional – This field is required for sourcing\_rule types
buy.

transportation\_lane\_id

Optional – This field is required for sourcing\_rule types
transfer.Optional – This field is required for sourcing\_rule types
transfer.

production\_process\_id

Optional – This field is required for sourcing\_rule types
manufacture.Optional – This field is required for sourcing\_rule types
manufacture.

sourcing\_priority

OptionalOptional

min\_qty

OptionalOptional

max\_qty

OptionalOptional

qty\_multiple

OptionalOptional

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

[sourcing\_schedule](planning-sourcing-schedule-entity.md)

###### Note

This data entity is optional.

sourcing\_schedule\_id

RequiredRequired

company\_id

OptionalOptional

tpartner\_id

Optional – This field is required for schedule\_type
InboundOrdering.Optional – This field is required for schedule\_type
InboundOrdering.

status

RequiredRequired

from\_site\_id

Optional – This field is required for schedule\_type
OutboundShipping.Optional – This field is required for schedule\_type
OutboundShipping.

to\_site\_id

RequiredRequired

schedule\_type

Required – The allowed values for this field are
InboundOrdering and OutboundShipping.Required – The allowed values for this field are
InboundOrdering and OutboundShipping.

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

[sourcing\_schedule\_details](planning-sourcing-schedule-details-entity.md)

###### Note

This data entity is optional.

sourcing\_schedule\_detail\_id

RequiredRequired

sourcing\_schedule\_id

RequiredRequired

company\_id

OptionalOptional

product\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

product\_group\_id

Optional – Either product\_id or product\_group\_id is
required.Optional – Either product\_id or product\_group\_id is
required.

day\_of\_week

OptionalOptional

week\_of\_month

OptionalOptional

time\_of\_day

OptionalOptional

date

OptionalOptional

[product\_bom](planning-product-bom-entity.md)

id

Not requiredRequired

product\_id

Not requiredRequired

company\_id

OptionalOptional

site\_id

Not requiredRequired

production\_process\_id

Not requiredRequired

component\_product\_id

Not requiredRequired

component\_quantity\_per

Not requiredRequired

assembly\_cost

Not requiredOptional

assembly\_cost\_uom

Not requiredOptional

priority

Not requiredOptional

eff\_start\_date

Not requiredRequired

eff\_end\_date

Not requiredRequired

[production\_process](operation-production-process-entity.md)

production\_process\_id

Not requiredRequired

production\_process\_name

Not requiredOptional

product\_id

Not requiredRequired

site\_id

Not requiredRequired

company\_id

OptionalOptional

setup\_time

Not requiredOptional

setup\_time\_uom

Not requiredOptional

operation\_time

Not requiredOptional

operation\_time\_uom

Not requiredOptional

[inv\_level](inventory-mgmnt-inv-level-entity.md)

snapshot\_date

RequiredRequired

site\_id

RequiredRequired

product\_id

RequiredRequired

company\_id

OptionalOptional

on\_hand\_inventory

RequiredRequired

allocated\_inventory

Not requiredNot required

bound\_inventory

Not requiredNot required

lot\_number

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

expiry\_date

Not requiredNot required

[forecast](forecast-forecast-entity.md)

site\_id

RequiredRequired

product\_id

RequiredRequired

mean

OptionalOptional

p10

OptionalOptional

p50

OptionalOptional

p90

OptionalOptional

forecast\_start\_dttm

RequiredRequired

forecast\_end\_dttm

RequiredRequired

snapshot\_date

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

region\_id

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

product\_group\_id

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

[vendor\_product](vendor-management-product-entity.md)

company\_id

OptionalOptional

vendor\_tpartner\_id

RequiredRequired

product\_id

RequiredRequired

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

[vendor\_lead\_time](vendor-management-lead-time-entity.md)

company\_id

OptionalOptional

vendor\_tpartner\_id

RequiredRequired

product\_id

OptionalOptional

site\_id

RequiredRequired

planned\_lead\_time

RequiredRequired

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

product\_group\_id

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

region\_id

Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.Required – When you ingest data from SAP or EDI, the default
value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data
using the Amazon S3 connector, you must enter a value or use
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

[outbound\_order\_line](outbound-fulfillment-order-line-entity.md)

id

Required – This field determines the outbound shipment
id.Required – This field determines the outbound shipment
id.

product\_id

Required – This field determines the id of the product
shipped.Required – This field determines the id of the product
shipped.

cust\_order\_id

Required – This field determines the id of the outbound
order.Required – This field determines the id of the outbound
order.

ship\_from\_site\_id

Required – This field determines the site from where the
product units are requested.Required – This field determines the site from where the
product units are requested.

ship\_to\_site\_id

Not requiredNot required

init\_quantity\_requested

Optional – This field determines the final quantity after any
cancellations and changes.Optional – This field determines the final quantity after any
cancellations and changes.

quantity\_promised

Optional – This field displays the promised quantity.Optional – This field displays the promised quantity.

quantity\_delivered

Optional – This field displays the actual quantity
delivered.Optional – This field displays the actual quantity
delivered. final\_quantity\_requested Optional – Final quantity after any cancellations or changes Optional – Final quantity after any cancellations or changes

status

Optional – This field determines the status of the order line,
that is, canceled, open, closed, and so on.Optional – This field determines the status of the order line,
that is, canceled, open, closed, and so on.

requested\_delivery\_date

RequiredRequired

promised\_delivery\_date

OptionalOptional

actual\_delivery\_date

OptionalOptional

[segmentation](planning-segmentation-entity.md)

segment\_id

RequiredRequired

creation\_date

RequiredRequired

company\_id

OptionalOptional

site\_id

RequiredRequired

product\_id

RequiredRequired

segment\_description

OptionalOptional

segment\_type

OptionalOptional

segment\_value

OptionalOptional

source

OptionalOptional

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

[company](organization-company-entity.md)

###### Note

This data entity is optional.

id

RequiredRequired

description

OptionalOptional

address\_1

OptionalOptional

address\_2

OptionalOptional

address\_3

OptionalOptional

city

OptionalOptional

state\_prov

OptionalOptional

postal\_code

OptionalOptional

country

OptionalOptional

phone\_number

OptionalOptional

time\_zone

OptionalOptional

calendar\_id

OptionalOptional

[supply\_planning\_paramters](planning-supply-planning-parameters-entity.md)

###### Note

This data entity is optional.

product\_id

RequiredRequired

product\_group\_id

Required. For future Use. Please populate
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.Required. For future Use. Please populate
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.

site\_id

Required. For future Use. Please populate
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.Required. For future Use. Please populate
SCN\_RESERVED\_NO\_VALUE\_PROVIDED for now.

planner\_name

OptionalOptional

demand\_time\_fence\_days

Optional.For future useOptional.For future use

forecast\_consumption\_backward\_days

Optional.For future useOptional.For future use

forecast\_consumption\_forward\_days

Optional.For future useOptional.For future use

eff\_start\_date

RequiredRequired

eff\_end\_date

RequiredRequired

[shipment](replenishment-shipment-entity.md)

id

RequiredNA

ship\_to\_site\_id

RequiredNA

product\_id

RequiredNA

ship\_from\_site\_id

Required – Supply Planning can use the value from
_ship\_from\_site\_id_ or
_supplier\_tpartner\_id_.

NA

supplier\_tpartner\_id

Required – Supply Planning can use the value from
_ship\_from\_site\_id_ or
_supplier\_tpartner\_id_.

NA

order\_type

RequiredNA

units\_shipped

RequiredNA

planned\_delivery\_date

Required – Supply Planning can use the value from
_planned\_delivery\_date_,
_actual\_delivery\_date_, or
_carrier\_eta\_date_.

NA

actual\_delivery\_date

carrier\_eta\_date

planned\_ship\_date

Required – Supply Planning can use the value from
_planned\_ship\_date_, or
_actual\_ship\_date_.

NA

actual\_ship\_date

creation\_date

OptionalNA

shipment\_status

OptionalNA

order\_id

Required. When you ingest data from SAP or EDI, the
default value for string is
_SCN\_RESERVED\_NO\_VALUE\_PROVIDED_. When you upload
data using the Amazon S3 connector, you must enter a value or use
_SCN\_RESERVED\_NO\_VALUE\_PROVIDED_ for successful
ingestion.NA

order\_line\_id

package\_id

[shipment\_lot](replenishment-shipment-lot-entity.md)

id

RequiredNA

lot\_qty

RequiredNA

expiry\_date

OptionalNA

shipment\_id

RequiredNA

product\_id

Required. When you ingest data from SAP or EDI, the
default value for string is
_SCN\_RESERVED\_NO\_VALUE\_PROVIDED_. When you upload
data using the Amazon S3 connector, you must enter a value or use
_SCN\_RESERVED\_NO\_VALUE\_PROVIDED_ for successful
ingestion.NA

tpartner\_id

order\_id

order\_line\_id

package\_id

[Document Conventions](../../../../general/latest/gr/docconventions.md)

N-Tier Visibility

Insights

All content copied from https://docs.aws.amazon.com/.
