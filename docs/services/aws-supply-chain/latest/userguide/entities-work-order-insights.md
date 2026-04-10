# Order Planning and Tracking

###### Note

To generate an order insight, in addition to ingesting the required data entities and columns, you must configure your milestone and process definitions. For more information on configuring orders, see [Configuring Order Planning and Tracking for the first time](setting-up-work-orders.md).

The table below lists the required data entities and columns to generate a order planning and tracking process.

Data entityColumnIs the column used by Order Planning and Tracking?

[site](network-site-entity.md)

###### Note

The _site_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

[product](product-product-entity.md)

###### Note

The _product_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

id

[vendor\_product](vendor-management-product-entity.md)

###### Note

The _vendor\_product_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

vendor\_tpartner\_id

product\_id

eff\_start\_date

eff\_end\_date

[geography](organization-geography-entity.md)

id

Required – This column is used by conditional filters to display regions or country.

[inbound\_order](replenishment-inbound-order-entity.md)

###### Note

The _inbound\_order_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

id

Required

tpartner\_id

Required

[inbound\_order\_line](replenishment-inbound-order-line-entity.md)

###### Note

The _inbound\_order\_line_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

order\_id

tpartner\_id

product\_id

[shipment](replenishment-shipment-entity.md)

###### Note

The _shipment_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

id

supplier\_tpartner\_id

product\_id

order\_id

order\_line\_id

package\_id

[reservation](planning-reservation-entity.md)

###### Note

The _reservation_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

reservation\_id

Required – This column is a required key for the _reservation\_id_ column in the _process\_product_ data entity.

reservation\_type

Required – This column is used when defining a default order plan.

reservation\_detail\_id

Required – This column is a required key for the _reservation\_detail\_id_ column in the _process\_product_ data entity.

[process\_header](operation-process-header-entity.md)

###### Note

The _process\_header_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

process\_id

Required

site\_id

Required – This column is used by the _site\_id_ column in the _process\_header_ data entity. For example, this column can be referenced in the milestone rules for specific processes.

status

Required

required\_on\_site

Required – This date is required to calculate the forecast completion date and to determine the Order line status.

[process\_product](operation-process-product-entity.md)

###### Note

The _process\_product_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

process\_product\_id

Required – This column is part of the primary key in the _process\_product_ data entity and is used as a reference in other entities.

process\_id

Required – This column is part of the primary key in the _process\_product_ data entity and is used to associate the header with the line.

product\_id

Required

reservation\_id

Required

reservation\_detail\_id

Required

requested\_availability\_date

Required – The field is displayed as _Required on site date_ in the AWS Supply Chain web application. This date is required to calculate the forecast completion date and to determine the Order line status.
When you ingest data, you must enter a value for _requested\_availability\_date_. When information is not available for the _requested\_availability\_date_ column, order planning and tracking will use the column values from _process\_header_ \> _planned\_start\_date_
to calculate the forecast completion date.

[work\_order\_plan](work-order-plan-entity.md)

process\_id

Required

product\_id

Required

business\_process\_id

Required

business\_process\_sequence

Required

preferred\_source

Required

duration

Required – This column provides the process lead time to determine the target date of the process completion.

The following table describes the data entities that are _not_ required to generate order planning and tracking. If these data entities are included in your dataset, the required columns are listed in the table below.

Data entityColumnIs the column used by Order Planning and Tracking?

[trading\_partner](organization-trading-partner-entity.md)

id

Required – This column is used to link the trading partner.

tpartner\_type

geo\_id

eff\_start\_date

eff\_end\_date

[process\_operation](operation-process-operation-entity.md)

###### Note

The _process\_operation_ data entity columns not listed in this table are _optional_ for order planning and tracking. AWS Supply Chain highly recommends ingesting data
for the _optional_ columns to enhance the feature output. When data is ingested for the _optional_ columns, you can use them to configure rules to evaluate the process milestones.

process\_operation\_id

Required

process\_id

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Insights

Demand Planning

All content copied from https://docs.aws.amazon.com/.
