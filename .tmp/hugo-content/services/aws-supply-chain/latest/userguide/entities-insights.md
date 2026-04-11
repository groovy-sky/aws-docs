# Insights

The table below list the data entities and columns used by Insights for the Inventory Visibility, Network Map, Inventory Insights, and Rebalance Recommendations features. See the table below
on how each feature in Insights uses the data entities.

###### Note

**How to read the table:**

- **Required** – The column name is mandatory in your dataset and you must populate the column name with values.

- **Optional** – The column name is optional. For enhanced feature output, it is recommended to add the column name with values.

- **Not required** – Data entity not required.

Data entityColumnIs the column used for Inventory visibility?Is the column used for Network map?Is the column used for Inventory Insights?Is the column used for Rebalance recommendations?Is the column used for Lead time Insights?

[site](network-site-entity.md)

id

RequiredRequiredRequiredRequiredRequired

description

RequiredRequiredRequiredRequiredOptional

geo\_id

Required – This field is required for filters to group sites by geographical groups such as region/country/state and so on.

Required – This field is required for filters to group sites by geographical groups such as region/country/state and so on.

Required – This field is required for filters to group sites by geographical groups such as region/country/state and so on.

Required

Required – This field is required for filters to group sites by geographical groups such as region/country/state and so on.

site\_type

Optional – Populating this column will display the site type on the inventory visibility page such as RDC, CDC, manufacturing site and so on.

Optional

Optional

Optional

Optional

company\_id

Optional

Optional

Optional

Optional

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

latitude

Optional

Required – This field is used to view the _site_ on the Network Map page.

Optional

Optional

Column name _latitude_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

longitude

Optional

Required – This field is used to view the _site_ on the Network Map page.

Optional

Optional

Column name _longitude_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

is\_active

Required – Identifies if the site needs to be considered for Insights computation. **Note:** If you want a site to be excluded from the Insights computation, make sure you set the column value to _False_.
If the column is blank or null, the site is considered active.

Required – Identifies if the site needs to be considered for Insights computation. **Note:** If you want a site to be excluded from the Insights computation, make sure you set the column value to _False_.
If the column is blank or null, the site is considered active.

Required – Identifies if the site needs to be considered for Insights computation. **Note:** If you want a site to be excluded from the Insights computation, make sure you set the column value to _False_.
If the column is blank or null, the site is considered active.

Required – Identifies if the site needs to be considered for Insights computation. **Note:** If you want a site to be excluded from the Insights computation, make sure you set the column value to _False_.
If the column is blank or null, the site is considered active.

Required – Identifies if the site needs to be considered for Insights computation. **Note:** If you want a site to be excluded from the Insights computation, make sure you set the column value to _False_.
If the column is blank or null, the site is considered active.

open\_date

Optional

Optional

Optional

Optional

Column name _open\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

end\_date

Optional

Optional

Optional

Optional

Column name _end\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

[transportation\_lane](network-transporation-lane-entity.md)

id

Not required

Not required

Not required

RequiredRequired

from\_site\_id

Not required

Not required

Not required

RequiredRequired. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

to\_site\_id

Not required

Not required

Not required

RequiredRequired. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

product\_group\_id

Not required

Not required

Not required

Required

Column name _product\_group\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

transit\_time

Not required

Not required

Not required

Required

Column name _transit\_time_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

time\_uom

Not required

Not required

Not required

Required – Supports day or days as units.

Column name _time\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

distance

Not required

Not required

Not required

Required

Column name _distance_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

distance\_uom

Not required

Not required

Not required

Required – Supports mile(s), km(s), or Kilometer(s) as units.

Column name _distance\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

eff\_start\_date

Not required

Not required

Not required

Optional

Column name _eff\_start\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

eff\_end\_date

Not required

Not required

Not required

Optional

Column name _eff\_end\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

product\_id

Not required

Not required

Not required

Optional – Either _product\_id_ or _product-group-id_ is required. When the lane is linked with a product, this field is mandatory.

Column name _product\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

emissions\_per\_unit

Not required

Not required

Not required

Optional

Column name _emissions\_per\_unit_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

emissions\_per\_weight

Not required

Not required

Not required

Optional

Column name _emissions\_per\_unit_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

company\_id

Not required

Not required

Not required

Optional

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

from\_geo\_id

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

to\_geo\_id

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

carrier\_tpartner\_id

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

service\_type

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

trans\_mode

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

cost\_per\_unit

Not required

Not required

Not required

Optional – You can view the shipping cost unit by lane during rebalance recommendations.

Column name _cost\_per\_unit_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

cost\_currency

Not required

Not required

Not required

Optional – You can view the shipping cost unit by lane during rebalance recommendations.

Column name _cost\_currency_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

[product](product-product-entity.md)

id

RequiredRequiredRequiredRequiredRequired

description

RequiredRequiredRequiredRequiredRequired

product\_group\_id

Required – Using this field, you can group products by product category such dairy, clothes, and so on.

Required – Using this field, you can group products by product category such dairy, clothes, and so on.

Required – Using this field, you can group products by product category such dairy, clothes, and so on.

Required

Required – Using this field, you can group products by product category such dairy, clothes, and so on.

is\_deleted

Required – Identifies if the product needs to be considered for Insights computation. **Note:** If you want the product to be excluded from the
Insights computation, make sure you set the column value to _True_ and set to _False_ to include this product for Insights computation. If the column
is left blank or null, the system considers the default value of _True_.

Required – Identifies if the product needs to be considered for Insights computation. **Note:** If you want the product to be excluded from the
Insights computation, make sure you set the column value to _True_ and set to _False_ to include this product for Insights computation. If the column
is left blank or null, the system considers the default value of _True_.

Required – Identifies if the product needs to be considered for Insights computation. **Note:** If you want the product to be excluded from the
Insights computation, make sure you set the column value to _True_ and set to _False_ to include this product for Insights computation. If the column
is left blank or null, the system considers the default value of _True_.

Required – Identifies if the product needs to be considered for Insights computation. **Note:** If you want the product to be excluded from the
Insights computation, make sure you set the column value to _True_ and set to _False_ to include this product for Insights computation. If the column
is left blank or null, the system considers the default value of _True_.

Required – Identifies if the product needs to be considered for Insights computation. **Note:** If you want the product to be excluded from the
Insights computation, make sure you set the column value to _True_ and set to _False_ to include this product for Insights computation. If the column
is left blank or null, the system considers the default value of _True_.

product\_type

Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.

Column name _product\_type_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

parent\_product\_id

Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.Optional – This field is required to support multiple product levels such as planning and fulfillment product.

Column name _parent\_product\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

base\_uom

Optional – This field is required for Insights to calculate the default base uom for a given product.Optional – This field is required for Insights to calculate the default base uom for a given product.Optional – This field is required for Insights to calculate the default base uom for a given product.Optional – This field is required for Insights to calculate the default base uom for a given product.

Column name _base\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

[product\_hierarchy](product-hierarchy-entity.md)

id

RequiredRequiredRequiredRequiredRequired

description

Required – Using this field, you can filter groups by product category such dairy, clothes, and so on.

Required – Using this field, you can filter groups by product category such dairy, clothes, and so on.

Required – Using this field, you can filter groups by product category such dairy, clothes, and so on.

Required – Using this field, you can filter groups by product category such dairy, clothes, and so on.

Required – Using this field, you can filter groups by product category such dairy, clothes, and so on.

parent\_product\_group\_id

Optional – This field is used by filters to support multiple product hierarchy category such as dairy, frozen diary products, fresh diary and so on.

Optional – This field is used by filters to support multiple product hierarchy category such as dairy, frozen diary products, fresh diary and so on.

Optional – This field is used by filters to support multiple product hierarchy category such as dairy, frozen diary products, fresh diary and so on.

Column name _parent\_product\_group\_id_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional – This field is used by filters to support multiple product hierarchy category such as dairy, frozen diary products, fresh diary and so on.

[product\_uom](product-uom-entity.md)

###### Note

This data entity is _optional_. For product uom conversions, data is _required_ in either _product-uom_, _uom\_conversion_, or _Insights_.

product\_uom\_id

Required – This field is required to perform the product uom conversion.Required – This field is required to perform the product uom conversion.Required – This field is required to perform the product uom conversion.Required – This field is required to perform the product uom conversion.Not required

product\_id

Required

Required

Required

Required

Not required

uom

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Not required

description

Optional

Optional

Optional

Optional

Not required

quantity

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Not required

quantity\_uom

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Not required

eff\_start\_date

Optional

Optional

Optional

Optional

Not required

eff\_end\_date

Optional

Optional

Optional

Optional

Not required

company\_id

Optional

Optional

Optional

Optional

Not required

[uom\_conversion](reference-uom-conversion-entity.md)

###### Note

This data entity is _optional_. For product uom conversions, data is _required_ in either _product-uom_, _uom\_conversion_, or _Insights_.

uom

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Required – This field is required for conversion from units.

Not required

company\_id

Optional

Optional

Optional

Optional

Not required

conversion\_uom\_id

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Required – This field is required for conversion to units.

Not required

conversion\_factor

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Required – This field contains the conversion factor.

Not required

[geography](organization-geography-entity.md)

id

RequiredRequiredRequiredRequiredRequired

description

RequiredRequiredRequiredRequiredRequired

parent\_geo\_id

Optional – This field is used to support multiple location hierarchy such as USA, USA-East, and so on.

Required – This field is used to support multiple location hierarchy such as USA, USA-East, and so on.

OptionalOptional

Required – This field is used to support multiple location hierarchy such as USA, USA-East, and so on.

[trading\_partner](organization-trading-partner-entity.md)

id

RequiredRequiredRequiredRequiredRequired

description

OptionalOptionalOptionalOptionalRequired

country

OptionalOptionalOptionalOptionalOptional

eff\_start\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Column name _eff\_start\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

eff\_end\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

Column name _eff\_end\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

time\_zone

OptionalOptionalOptionalOptional

Column name _time\_zone_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

is\_active

OptionalOptionalOptionalOptional

Column name _is\_active_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

tpartner\_type

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

Column name _tpartner\_type_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

geo\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

Column name _geo\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

[inbound\_order](replenishment-inbound-order-entity.md)

###### Note

This data entity is _required_ for _Lead Time Insights_ but _optional_ for _Inventory Visibility_,
_Network Map_, _Inventory Insights_, and _Rebalance Recommendations_.

id

Not required

Not required

Not required

Not required

Required

order\_type

Not required

Not required

Not required

Not required

Optional – Data can be used by inbound order line.

order\_status

Not required

Not required

Not required

Not required

Optional

to\_site\_id

Not required

Not required

Not required

Not required

Column name _site\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

submitted\_date

Not required

Not required

Not required

Not required

Requiredtpartner\_idRequired. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

[inbound\_order\_line](replenishment-inbound-order-line-entity.md)

id

RequiredRequiredRequiredRequiredRequired

order\_id

RequiredRequiredRequiredRequiredRequired

order\_type

RequiredRequiredRequiredRequiredOptional

status

RequiredRequiredRequiredRequiredOptional

product\_id

RequiredRequiredRequiredRequiredRequired

to\_site\_id

RequiredRequiredRequiredRequiredRequired

from\_site\_id

RequiredRequiredRequiredRequiredRequired

quantity\_submitted

Required – One quantity field should be set.

Required – One quantity field should be set.

Required – One quantity field should be set.

Required – One quantity field should be set.

Required – One quantity field should be set.

quantity\_confirmed

Optional – One quantity field should be set.

Optional – One quantity field should be set.

Optional – One quantity field should be set.

Optional – One quantity field should be set.

Optional – One quantity field should be set.

quantity\_received

Optional – This field should be blank for open orders.

Optional – This field should be blank for open orders.

Optional – This field should be blank for open orders.

Optional – This field should be blank for open orders.

Optional – This field should be blank for open orders.

quantity\_uom

Required – This field is required to determine the unit for quantity fields.

Required – This field is required to determine the unit for quantity fields.

Required – This field is required to determine the unit for quantity fields.

Required – This field is required to determine the unit for quantity fields.

Column name _quantity\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

expected\_delivery\_date

RequiredRequiredRequiredRequiredRequired

submitted\_date

Column name _submitted\_date_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _submitted\_date_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _submitted\_date_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _submitted\_date_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Required

incoterm

Column name _incoterm_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _incoterm_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _incoterm_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _incoterm_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

product\_group\_id

Column name _product\_group\_id_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _product\_group\_id_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _product\_group\_id_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _product\_group\_id_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

company\_id

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

tpartner\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

reservation\_id

Optional – This field is used to determine the connection between order line and order line schedule. For example, 1001 - A, where 1001 is the order\_id and A is the order\_line\_id in the inbound\_order\_line\_schedule table.Optional – This field is used to determine the connection between order line and order line schedule. For example, 1001 - A, where 1001 is the order\_id and A is the order\_line\_id in the inbound\_order\_line\_schedule table.Optional – This field is used to determine the connection between order line and order line schedule. For example, 1001 - A, where 1001 is the order\_id and A is the order\_line\_id in the inbound\_order\_line\_schedule table.Optional – This field is used to determine the connection between order line and order line schedule. For example, 1001 - A, where 1001 is the order\_id and A is the order\_line\_id in the inbound\_order\_line\_schedule table.Column name

_reservation\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

order\_receive\_date

Column name _order\_receive\_date_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _order\_receive\_date_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _order\_receive\_date_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _order\_receive\_date_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

[inbound\_order\_line\_schedule](replenishment-inbound-order-line-schedule-entity.md)

###### Note

This data entity is _required_ for _Lead Time Insights_ but _optional_ for _Inventory Visibility_,
_Network Map_, _Inventory Insights_, and _Rebalance Recommendations_.

When data is not ingested for this data entity, Insights will use the supply data from _inbound\_order\_line_ data entity. For custom configurations, contact [Get support for AWS Supply Chain](admin-support-ug.md).

id

RequiredRequiredRequiredRequiredRequired

order\_id

Required – This field is required to link back to an order line along with the order\_line\_id.Required – This field is required to link back to an order line along with the order\_line\_id.Required – This field is required to link back to an order line along with the order\_line\_id.Required – This field is required to link back to an order line along with the order\_line\_id.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

order\_line\_id

Required – This field is required to link back to an order line along with order\_id.Required – This field is required to link back to an order line along with order\_id.Required – This field is required to link back to an order line along with order\_id.Required – This field is required to link back to an order line along with order\_id.

Column name _order\_line\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

company\_id

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

Column name _company\_id_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

product\_id

RequiredRequiredRequiredRequiredRequired

expected\_delivery\_date

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional

delivery\_date

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Optional – _delivery\_date_ or _expected\_delivery date_ must be provided.

Column name _delivery\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

ship\_date

Optional – Date when the order was shipped.

Optional – Date when the order was shipped.

Optional – Date when the order was shipped.

Optional – Date when the order was shipped.

Column name _ship\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

quantity\_submitted

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Column name _quantity\_submitted_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

quantity\_confirmed

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Column name _quantity\_confirmed_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

quantity\_received

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Required – One quantity field should be set. This field uses the uom set at the line level.

Column name _quantity\_received_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

[shipment](replenishment-shipment-entity.md)

id

RequiredRequiredRequiredRequiredOptional

order\_id

Required – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.RequiredRequired – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.Required – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.Required

order\_line\_id

Required – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.RequiredRequired – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.Required – This field is required to calculate the _in-transit_ and _on-order_ values for projected inventory visibility.Required

product\_id

RequiredRequiredRequiredRequiredRequired

ship\_to\_site\_id

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Required

actual\_delivery\_date

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Required

units\_shipped

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

uom

Optional – This field is used to determine the unit for quantity fields.

Optional – This field is used to determine the unit for quantity fields.

Optional – This field is used to determine the unit for quantity fields.

Optional – This field is used to determine the unit for quantity fields.

Optional – This field is used to determine the unit for quantity fields.

planned\_ship\_date

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Column name _planned\_ship\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

actual\_ship\_date

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Optional – _planned\_ship\_date_ or _actual\_ship\_date_ must be provided.

Column name _actual\_ship\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

planned\_delivery\_date

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Optional – _planned\_delivery\_date_ or _actual\_delivery\_date_ must be provided.

Column name _planned\_delivery\_date_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

ship\_from\_site\_id

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional – Derived from inbound order line.

Optional

supplier\_tpartner\_id

Column name _supplier\_tpartner\_id_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _supplier\_tpartner\_id_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _supplier\_tpartner\_id_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _supplier\_tpartner\_id_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

transportation\_mode

Column name _transportation\_mode_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _transportation\_mode_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _transportation\_mode_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _transportation\_mode_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

ship\_from\_site\_address\_country

Column name _ship\_from\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _ship\_from\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _ship\_from\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _ship\_from\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

ship\_to\_site\_address\_country

Column name _ship\_to\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _ship\_to\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _ship\_to\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _ship\_to\_site\_address\_country_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

carrier\_id

Column name _carrier\_id_ should be available in your dataset. Value for the column name is not required for Inventory visibility.

Column name _carrier\_id_ should be available in your dataset. Value for the column name is not required for Network map.

Column name _carrier\_id_ should be available in your dataset. Value for the column name is not required for Inventory Insights.

Column name _carrier\_id_ should be available in your dataset. Value for the column name is not required for Rebalance Recommendations.

Optional

package\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

[inv\_policy](planning-inv-policy-entity.md)

id

RequiredRequiredRequiredRequiredRequired

site\_id

RequiredRequiredRequiredRequiredRequired

product\_id

RequiredRequiredRequiredRequiredRequired

min\_safety\_stock

RequiredRequiredRequiredRequiredRequired

max\_safety\_stock

RequiredRequiredRequiredRequiredRequired

qty\_uom

Optional – This field is used to determine the UOM for inventory policy.Optional – This field is used to determine the UOM for inventory policy.Optional – This field is used to determine the UOM for inventory policy.Optional – This field is used to determine the UOM for inventory policy.Optional – This field is used to determine the UOM for inventory policy.

min\_doc\_limit

Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.

Column name _min\_doc\_limit_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

max\_doc\_limit

Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.Optional – This field is required if you want to see the days of cover.

Column name _max\_doc\_limit_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

eff\_start\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

eff\_end\_date

Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.Required – You must enter a value for eff\_start\_date and eff\_end\_date. If you don't have a value, enter `1900-01-01 00:00:00 ` for eff\_start\_date, and `9999-12-31 23:59:59` for eff\_end\_date.

company\_id

OptionalOptionalOptionalOptionalOptional

ss\_policy

Required – _abs\_level_ when there is no value.

Required – _abs\_level_ when there is no value.

Required – _abs\_level_ when there is no value.

Required – _abs\_level_ when there is no value.

Required – _abs\_level_ when there is no value.

fallback\_policy\_1

OptionalOptionalOptionalOptional

Column name _fallback\_policy\_1_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

product\_group\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

dest\_geo\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

vendor\_tpartner\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

[inv\_level](inventory-mgmnt-inv-level-entity.md)

###### Note

Enter the on-hand inventory at the beginning of the day.

snapshot\_date

RequiredRequiredRequiredRequiredRequired

site\_id

RequiredRequiredRequiredRequiredRequired

product\_id

RequiredRequiredRequiredRequiredRequired

company\_id

OptionalOptionalOptionalOptionalOptional

on\_hand\_inventory

RequiredRequiredRequiredRequiredRequired

allocated\_inventory

OptionalOptionalOptionalOptional

Column name _allocated\_inventory_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

quantity\_uom

Optional – This field is used to determine the quantity UOM for inventory records.Optional – This field is used to determine the quantity UOM for inventory records.Optional – This field is used to determine the quantity UOM for inventory records.Optional – This field is used to determine the quantity UOM for inventory records.

Column name _quantity\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

inv\_condition

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

lot\_number

Required – Insights expects one inventory level record per site and product for the given snapshot date. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required – Insights expects one inventory level record per site and product for the given snapshot date. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required – Insights expects one inventory level record per site and product for the given snapshot date. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required – Insights expects one inventory level record per site and product for the given snapshot date. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required – Insights expects one inventory level record per site and product for the given snapshot date. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

[forecast](forecast-forecast-entity.md)

site\_id

RequiredRequiredRequiredRequired

Not required

product\_id

RequiredRequiredRequiredRequired

Not required

mean

RequiredRequiredRequiredRequired

Not required

forecast\_start\_dttm

###### Note

Make sure the _forecast\_start\_dttm_ and _forecast\_end\_dttm_ are set at different dates when forecast is set at a daily interval. The _forecast\_end\_dttm_
must be set at the explicit end of the date range.

RequiredRequiredRequiredRequired

Not required

forecast\_end\_dttm

RequiredRequiredRequiredRequired

Not required

quantity\_uom

Optional – This field is used to determine the quantity UOM for forecast.Optional – This field is used to determine the quantity UOM for forecast.Optional – This field is used to determine the quantity UOM for forecast.Optional – This field is used to determine the quantity UOM for forecast.

Column name _quantity\_uom_ should be available in your dataset. Value for the column name is not required for Lead Time Insights.

snapshot\_date

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

Not required

region\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

Not required

product\_group\_id

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

Not required

[vendor\_lead\_time](vendor-management-lead-time-entity.md)

company\_id

Not required

Not required

Not required

Not required

Optional

vendor\_tpartner\_id

Not required

Not required

Not required

Not required

Required

product\_id

Not required

Not required

Not required

Not required

Required

site\_id

Not required

Not required

Not required

Not required

Required

planned\_lead\_time

Not required

Not required

Not required

Not required

Required

eff\_start\_date

Not required

Not required

Not required

Not required

Optional

eff\_end\_date

Not required

Not required

Not required

Not required

Optional

product\_group\_id

Not required

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

region\_id

Not required

Not required

Not required

Not required

Required. When you ingest data from SAP or EDI, the default value for string is SCN\_RESERVED\_NO\_VALUE\_PROVIDED. When you upload data using the Amazon S3 connector, you must enter a value or use SCN\_RESERVED\_NO\_VALUE\_PROVIDED for
successful ingestion.

source\_site\_id

Not required

Not required

Not required

Not required

Optional. Site from where the inbound shipment originated.

trans\_mode

Not required

Not required

Not required

Not required

Optional. Transportation mode used. For example, ship, truck, rail.

[outbound\_order\_line](outbound-fulfillment-order-line-entity.md)

###### Note

This data entity is _optional_. Insights will use the demand data from the _forecast_ entity. If you ingest information
for the _outbound\_order\_line_ entity, make sure the shipment information is also ingested for the _outbound\_shipment_ entity to gather the demand for the correct dates.

id

Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.

Not required

cust\_order\_id

Required. Determines the outbound order ID. Required. Determines the outbound order ID. Required. Determines the outbound order ID. Required. Determines the outbound order ID.

Not required

product\_id

Required. Determines the product ID shipped.Required. Determines the product ID shipped.Required. Determines the product ID shipped.Required. Determines the product ID shipped.

Not required

ship\_from\_site\_id

Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.

Not required

ship\_to\_site\_id

Optional. Site where the products should be shipped.Optional. Site where the products should be shipped.Optional. Site where the products should be shipped.Optional. Site where the products should be shipped.

Not required

final\_quantity\_requested

Optional. Final quantity after all updates and cancellations.Optional. Final quantity after all updates and cancellations.Optional. Final quantity after all updates and cancellations.Optional. Final quantity after all updates and cancellations.

Not required

quantity\_promisedRequired. Quantity agreed to be delivered.Required. Quantity agreed to be delivered.Required. Quantity agreed to be delivered.Required. Quantity agreed to be delivered.

Not required

quantity\_deliveredOptional. Actual quantity delivered.Optional. Actual quantity delivered.Optional. Actual quantity delivered.Optional. Actual quantity delivered.

Not required

statusOptional. Displays the status of the order line. For example, canceled, open, closed, and so on.Optional. Displays the status of the order line. For example, canceled, open, closed, and so on.Optional. Displays the status of the order line. For example, canceled, open, closed, and so on.Optional. Displays the status of the order line. For example, canceled, open, closed, and so on.

Not required

quantity\_uomOptional. Unit of measure for quantity. For example, eaches, cases.Optional. Unit of measure for quantity. For example, eaches, cases.Optional. Unit of measure for quantity. For example, eaches, cases.Optional. Unit of measure for quantity. For example, eaches, cases.

Not required

requested\_delivery\_dateOptionalOptionalOptionalOptional

Not required

promised\_delivery\_dateOptionalOptionalOptionalOptional

Not required

[outbound\_shipment](outbound-fulfillment-shipment-entity.md)

###### Note

This data entity is _optional_. AWS Supply Chain will use the demand data from the outbound\_order\_line or forecast data entity.

id

Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.Required. Determines the outbound shipment ID.

Not required

from\_site\_id

Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.Required. Determines the site from where the units are shipped.

Not required

product\_id

Required. Determines the product ID of the product shipped.Required. Determines the product ID of the product shipped.Required. Determines the product ID of the product shipped.Required. Determines the product ID of the product shipped.

Not required

cust\_order\_id

Required. Determines the outbound order ID.Required. Determines the outbound order ID.Required. Determines the outbound order ID.Required. Determines the outbound order ID.

Not required

cust\_order\_line\_id

Required. Determines the outbound order line ID. Required. Determines the outbound order line ID. Required. Determines the outbound order line ID. Required. Determines the outbound order line ID.

Not required

expected\_ship\_date

Required. Determines when the products exit the from\_site.Required. Determines when the products exit the from\_site.Required. Determines when the products exit the from\_site.Required. Determines when the products exit the from\_site.

Not required

actual\_ship\_date

Optional. Determines the actual date when the product exits the from\_site.Optional. Determines the actual date when the product exits the from\_site.Optional. Determines the actual date when the product exits the from\_site.Optional. Determines the actual date when the product exits the from\_site.

Not required

shipped\_qty

Required. Determines the quantity shipped from the from\_site.Required. Determines the quantity shipped from the from\_site.Required. Determines the quantity shipped from the from\_site.Required. Determines the quantity shipped from the from\_site.

Not required

cust\_shipment\_status

Optional. Status of the shipment. For example, canceled, open, closed, and so on.Optional. Status of the shipment. For example, canceled, open, closed, and so on.Optional. Status of the shipment. For example, canceled, open, closed, and so on.Optional. Status of the shipment. For example, canceled, open, closed, and so on.

Not required

to\_site\_id

Optional. Site where products should be shipped.Optional. Site where products should be shipped.Optional. Site where products should be shipped.Optional. Site where products should be shipped.

Not required

expected\_delivery\_dateOptionalOptionalOptionalOptional

Not required

actual\_delivery\_dateOptionalOptionalOptionalOptional

Not required

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Supply Planning

Order Planning and Tracking

All content copied from https://docs.aws.amazon.com/.
