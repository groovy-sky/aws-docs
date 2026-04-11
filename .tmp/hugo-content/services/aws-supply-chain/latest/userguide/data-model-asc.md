# Data entities supported in AWS Supply Chain

The following is an overview of the data entities supported in AWS Supply Chain.

###### Note

The data entities listed in this chapter are required for Data Lake ingestion. For data entities required for each AWS Supply Chain module, see [Data entities and columns used in AWS Supply Chain](data-model.md).

For information on application datasets displayed in AWS Supply Chain Analytics, see [Application datasets used in AWS Supply Chain Analytics](application-datasets.md).

CategoryCategory typeData entity and description

Organization

Non-transactional data

[company](organization-company-entity.md) \- Entity to store the
name and location of your company.

Non-transactional data

[geography](organization-geography-entity.md) \- Entity stores
geographical hierarchy of your company.

Non-transactional data

[trading\_partner](organization-trading-partner-entity.md) \- Contains the
partners that have trading relationship with your company, such as
vendors, 3PLs, channel partners, or distributors.

Non-transactional data

[trading\_partner\_poc](organization-trading-partner-poc-entity.md) \- Contains information that can be identified about the point of contacts at the partners such as vendors, 3PLs, channel partners,
or distributors, that have trading relationship with your company.

Product

Non-transactional data

[product](product-product-entity.md) \- Contains the key product
attributes, including name, description, brand, codes, category,
business group, and price.

Non-transactional data

[product\_hierarchy](product-hierarchy-entity.md) \- Contains the product
categories and sub-categories.

Non-transactional data

[product\_uom](product-uom-entity.md) \- Contains the product
packaging options and conversations between packages.

Non-transactional data

[product\_alternate](product-alternate-entity.md) \- Contains information about alternative products, including type of alternative.

Non-transactional data

[un\_details](product-un-details-entity.md) \- Contains information about hazardous products.

Network

Non-transactional data

[site](network-site-entity.md) \- Stores information for sites holding inventory such as Stores, Distribution Centers ,including ID, name, address, geographical region, and site type.

Non-transactional data

[transportation\_lane](network-transporation-lane-entity.md) \- Contains
information about transportation lanes, including from and to sites,
transportation mode, and transit time.

Vendor management

Non-transactional data

[vendor\_product](vendor-management-product-entity.md) \- Contains the
product information per vendor, including price, lead-time, and inbound
sites.

Non-transactional data

[vendor\_lead\_time](vendor-management-lead-time-entity.md) \- Contains the
planned and actual lead times from the vendor.

Non-transactional data

[vendor\_holiday](vendor-management-holiday-entity.md) \- Displays
information on vendor outages due to holidays and shutdowns.

Planning

Non-transactional data

[inv\_policy](planning-inv-policy-entity.md) \- Contains inventory
policies such as minimum and maximum safety stock policy, target inventory quantity, minimum or ma maximum order quantity and so on, for product, product-site, and other possible
combinations.

Non-transactional data

[segmentation](planning-segmentation-entity.md) \- Used to store segments. Segments are used in conjunction with product, site, and effective dates for uniqueness.
For example, HV1 for High Value, HLW for Halloween Products, seasonal, volatile and so on.

Non-transactional data

[sourcing\_rules](planning-sourcing-rules-entity.md) \- Defines rules at product-site level to specify the sourcing related attributes (for example, rule type, to and from site, transportation lane, minimum and
maximum quantity, priority, ratio, and so on).

Non-transactional data

[sourcing\_schedule](planning-sourcing-schedule-entity.md) \- Sourcing schedule determines when to source. For example, source from vendors or transfer between sites.

Non-transactional data

[sourcing\_schedule\_details](planning-sourcing-schedule-details-entity.md) \- Provides sourcing schedule details. For example, the days in a week, a product be sourced from a vendor.

Transactional data

[reservation](planning-reservation-entity.md) \- Provides details about inventory reservation. For example, reservation ID, type, date, quantity, product ID.

Transactional data

[product\_bom](planning-product-bom-entity.md) \- Displays bill of material for product with type, level, ratios, quantities, and cost attributes.

Operation

Transactional data

[process\_header](operation-process-header-entity.md) \- Track execution activities within a plant or site. For example, manufacturing, maintenance or repairs.

Transactional data

[process\_operation](operation-process-operation-entity.md) \- Defines operation associated with an activity. For example, Stop machine, Oiling, and so on.

Transactional data

[process\_product](operation-process-product-entity.md) \- Define the product or material associated with an activity.

Transactional data

[production\_process](operation-production-process-entity.md) \- Defines attributes associated with the manufacturing or production process.

Inventory Management

Transactional data

[inv\_level](inventory-mgmnt-inv-level-entity.md) \- A snapshot of the product’s inventory condition in each site. For example, snapshot date, on-hand inventory, condition of the product.

Inbound

Transactional data

[inbound\_order](replenishment-inbound-order-entity.md) \- Contains information about inbound orders into your companies locations. For example, for example, purchase orders (POs), blanket POs,
production orders, or stock transfer orders).

Transactional data

[inbound\_order\_line](replenishment-inbound-order-line-entity.md) \- Stores
line level information for inbound\_order, including product\_id, and
quantity.

Transactional data

[inbound\_order\_line\_schedule](replenishment-inbound-order-line-schedule-entity.md) -
Stores schedule-line level data within an inbound\_order\_line and is
relevant only when schedules are used.

Transactional data

[shipment](replenishment-shipment-entity.md) \- Stores shipment
information like origin, carrier code, ship date, product, quantity, ship from site, expected delivery date, and actual delivery date, or
inbound orders (PO,TO and so on) including ship date, product, quantity,
ship from site, expected delivery date, and actual delivery date.

Transactional data

[shipment\_stop](replenishment-shipment-stop-entity.md) \- Contains list of shipment stops with corresponding date and time. This field is used when there are multiple stops for shipments.

Transactional data

[shipment\_stop\_order](replenishment-shipment-stop-order-entity.md) \- Contains list of orders picked and dropped per shipment stop.

Transactional data

[shipment\_lot](replenishment-shipment-lot-entity.md) \- Contains the shipment details per shipment lot.

Outbound fulfillment

Transactional data

[outbound\_order\_line](outbound-fulfillment-order-line-entity.md) \- Contains orders originating from your company and shipped to locations outside of the your network.
Outbound\_order\_line contains order date, customer location, incoterms, and so on. It also includes product, price, discount, and units.

Transactional data

[outbound\_shipment](outbound-fulfillment-shipment-entity.md) \- Stores
shipment information for outbound orders, including ship date, product,
quantity, ship from site, expected delivery date, and actual delivery
date.

Cost management

Transactional data

[customer\_cost](customer-cost-entity.md) \- Displays the information about the costs incurred by you during the supply chain operations.

Plan

Transactional data

[supply\_plan](supply-plan-entity.md) \- Displays the supply plan generated by AWS Supply Chain Supply Planning.

Forecast

Transactional data

[forecast](forecast-forecast-entity.md) \- Stores forecast over
forecast horizon for product, product-site, or other
combinations.

Transactional data

[supplementary\_time\_series](forecast-supp-timeseries-entity.md) \- Displays additional demand driver time series information such as price, promotions, and out-of-stock indicator to improve forecast quality.

Reference

Non-transactional data

[reference\_field](reference-fields-entity.md) \- Contains mapping of any
entity-field-value combination to a corresponding description, such as
mapping specific inbound\_order status code to status description.

Non-transactional data

[calendar](reference-calendar-entity.md) \- Calendars can be used
for many purposes by the application, such as planning, execution, and
reporting.

Non-transactional data

[uom\_conversion](reference-uom-conversion-entity.md) \- Contains
conversions for unit of measure (UOM).

Insights

Transactional data

[work\_order\_plan](work-order-plan-entity.md) \- Provides the supply chain process plan for a work order along with source type and duration to finish each supply chain process.

###### Note

- All fields marked as type _timestamp_ should be in ISO 8601
format.

- The dataset that you ingest into AWS Supply Chain can only include the following
special characters: ASCII 35 (number sign: #), 36 (dollar sign: $), 37 (percent
sign: %), 45 (hyphen: -), 46 (period: .), 47 (slash: /), 94 (caret), 95
(underscore: \_), 123 (left curly brace: { ), and 125 (right curly brace:
}).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data mapping example for fulfillment

Organization

All content copied from https://docs.aws.amazon.com/.
