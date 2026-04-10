# N-Tier Visibility

The table below list the data entities and columns used by N-Tier Visibility.

###### Note

**How to read the table:**

- **Required** – The column name is mandatory in your dataset and you must populate the column name with values.

- **Optional** – The column name is optional. For enhanced feature output, it is recommended to add the column name with values.

- **Not required** – Data entity not required.

Data entityColumnIs the column used by N-Tier Visibility?

[trading\_partner](organization-trading-partner-entity.md)

id

Required

description

Required

company\_id

Optional

tpartner\_type

Required – When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

geo\_id

Required – When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for successful ingestion.

eff\_end\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

eff\_start\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

[trading\_partner\_poc](organization-trading-partner-poc-entity.md)

tpartner\_id

Required

email

Required

[product](product-product-entity.md)

id

Required – Data entity is optional but _id_ is used to generate Partner Network View.

[product\_hierarchy](product-hierarchy-entity.md)

id

[site](network-site-entity.md)

id

[sourcing\_rules](planning-sourcing-rules-entity.md)

sourcing\_rule\_id

Required – Data entity is optional but _sourcing\_rule\_id_ is used to generate Partner Network View.

[supply\_plan](supply-plan-entity.md)

supply\_plan\_id

Required

snapshot\_date

Optional

creation\_date

Optional

tpartner\_id

Required

product\_id

Required

to\_site\_id

Required

from\_site\_id

Optional

plan\_quantity

Required

plan\_type

Required

plan\_need\_by\_date

Required

quantity\_uom

Optional

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Sustainability

Supply Planning

All content copied from https://docs.aws.amazon.com/.
