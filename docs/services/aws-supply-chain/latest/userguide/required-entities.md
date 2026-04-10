# Demand Planning

The following table lists the data entities and columns used by Demand Planning.

###### How to read the table:

- **Required** – The columns in this data entity are mandatory to execute a demand forecast without any failures.

- **Conditionally required** – The columns in this data entity are required depending on the configurations set under demand plan settings. For more information, see [Manage Demand Plan settings](settings.md).

- **Recommended for forecast quality** – The columns in this data entity are required for the quality for the forecast.

- **Optional** – The column name is optional. For enhanced feature output, it is recommended to add the column name with values.

Data entityIs this data entity required?How is this data entity used?ColumnIs the column required?How is this column used in Forecasting?

outbound\_order\_line

RequiredDemand Planning uses this data as the primary
source of historical demand for forecast. Additionally, fields selected
as granularity are sent for training and are available as filters to
review the demand plan.

id

Required_id_, _cust\_order\_id_, and _product\_id_ are used to uniquely identify a record in the data entity and this combination should always be unique. Make sure the column values do not have invalid characters such as asterisk and double-quotes.

cust\_order\_id

Required

product\_id

Required

order\_date

RequiredRequired for forecast creation. Identifies the period for time-series forecasting.

final\_quantity\_requested

RequiredRequired for forecast creation. Identifies the quantity used for time-series forecasting. This column must not contain null values and must be _numerical_. Make sure there are no commas in the values. For example, 500000.00 is an accepted value in Demand Planning.

ship\_from\_site\_id

Conditionally requiredThis column is conditionally required for forecast creation _if_ the column is selected for forecast dimension (Site Hierarchy). This column must have a value and is used for filtering and analysis of data. For information on how to map data for different fulfillment scenarios,
see [Data mapping example for fulfillment](fulfillment-scenario.md).

ship\_to\_site\_id

Conditionally required

channel\_id

Conditionally requiredThis column is conditionally required for forecast creation _if_ the column is selected for forecast dimension (Channel Hierarchy). This column must have a value and is used for
filtering and analysis of data. For information on how to map data for different fulfillment scenarios, see [Data mapping example for fulfillment](fulfillment-scenario.md).

customer\_tpartner\_id

Conditionally requiredThis column is conditionally required for forecast creation _if_ the column is selected for forecast dimension (Customer Hierarchy). This column must have a value and is used for filtering and analysis of data.
For information on how to map data for different fulfillment scenarios, see [Data mapping example for fulfillment](fulfillment-scenario.md).

ship\_to\_site\_address\_city

Conditionally requiredThis column is conditionally required for forecast creation _if_ the column is selected for forecast dimension (Site Hierarchy). This column must have a value and is used for filtering and analysis of data. For information on how to map data for different fulfillment
scenarios, see [Data mapping example for fulfillment](fulfillment-scenario.md).

ship\_to\_site\_address\_state

Conditionally required

ship\_to\_site\_address\_country

Conditionally required

status

Recommended for forecast qualityThis column is recommended for forecast quality. Orders with _canceled_ status are not considered as forecast input.

product

RequiredDemand Planning uses the product attributes to
establish hierarchy filters for demand plan review and for model
training.

id

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). Make sure the column values do not have duplicate IDs and special characters such as asterix and double-quotes.

description

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). This column can contain special characters such as asterix, hyphen, quotes, and double-quotes.

parent\_product\_id

Conditionally requiredThis column is conditionally required for forecast creation _if_ the column is selected for forecast dimensions (Product Hierarchy). Make sure the column has values and is used for filtering and analysis of data and model training.

product\_group\_id

Conditionally required

product\_type

Conditionally required

brand\_name

Conditionally required

color

Conditionally required

display\_desc

Conditionally required

product\_available\_day

Recommended for forecast qualityRecommended. The value in this column improves forecast quality by allowing the forecasting model to consider the timing of new product introductions.

discontinue\_day

Recommended for forecast qualityRecommended. The value in this column improves forecast quality by allowing the forecasting model to consider the timing for product retirements.

base\_uom

Recommended for forecast qualityUnit of measure for product. Default is Eaches.

is\_deleted

Recommended for forecast qualityRecommended. Enter _Y_ if the product ID should be excluded from forecasting.

pkg\_height

Recommended for forecast qualityRecommended. The physical characteristics of the product that the forecasting models can understand.

pkg\_length

Recommended for forecast quality

pkg\_width

Recommended for forecast quality

shipping\_dimension

Recommended for forecast quality

casepack\_size

Recommended for forecast quality

product\_alternate

Recommended for forecast qualityDemand Planning uses the data of product’s predecessor(s) or alternate(s) to create forecast for new products. When data is ingested into the _product\_alternate_ data entity, Product lineage support for forecast is enabled. For more information,
see [Product lineage](product-lineage.md). You can skip ingesting data into the _product\_alternate_
data entity and the forecast can still be generated.

alternative\_product\_id

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). Unique record identifier.

product\_id

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). ID of the new product or new version of the product. Make sure _product\_id_ is populated in the _product_ data entity.

product\_alternate\_id

RequiredRequired for data ingestion into SCDL. Identifier for a similar product or previous version of the product. To consider multiple similar products as a single _product\_id_, enter the products in separate rows. Make sure _product\_alternate\_id_ is populated in the _product_ data entity.

alternate\_type

RequiredRequired for applying product supercession or lineage. Use the static value _similar\_demand\_product_ in all the rows.

alternate\_product\_qty

RequiredRequired for applying product supercession or lineage. Enter the proportion of history of the _alternate\_product\_id_ you want to use for forecasting _product\_id_. For example, if it is 60%, enter 60. When you have multiple _alternative\_product\_id_
for a single _product\_id_, the _alternate\_product\_qty_ does not have to add up to 100.

alternate\_product\_qty\_uom

RequiredRequired for applying product supercession or lineage. Use the specific static value "percentage".

eff\_start\_date

RequiredRequired for data ingestion into SCDL. Enter the start timeframe to consider the history of a similar product. Make sure this date is on or before the _eff\_end\_date_ or you can leave this field empty and
Demand Planning will auto-fill the year with 1000.

eff\_end\_date

RequiredRequired for data ingestion into SCDL. Enter the end timeframe to consider in history of a similar product. Make sure this date is on or after the _eff\_start\_date_ or you can leave this field empty and
Demand Planning will auto-fill the year with 9999..

status

Recommended for forecast qualityRecommended. Enter _Inactive_ to ignore the product supercession or lineage mapping.

supplementary\_time\_series

Recommended for forecast qualityDemand Planning uses this data as the primary source for tagging casual factors such as promotional events, discounts, holidays, and so on.

id

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). Unique record identifier.

order\_date

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). Timestamp when the timeseries was recorded.

time\_series\_name

RequiredRequired for data ingestion into Supply Chain Data Lake (SCDL). Name of the specific type of time series. The _time\_series\_name_ column must start with a letter, be 2 to 56
characters long, and can contain letters, numbers, and underscores. No other special characters are allowed.

time\_series\_value

RequiredRequired for data ingestion into SCDL. Value corresponding to the specific time series. Demand Planning only supports numerical input and time-series with categorical value is not considered.

product\_id

OptionalRecommended. Unique identifier for a specific product. Use this column if the demand driver is available at product level.

site\_id

OptionalRecommended. Unique identifier for a specific site or location. Use this column if the demand driver is available at site level. This column can represent either _ship\_from\_site\_id_ or _ship\_to\_site\_id_ based on the lowest level site hierarchy configuration.

channel\_id

OptionalRecommended. Unique identifier for a specific channel. Use this column if the demand driver is available at channel level.

customer\_tpartner\_id

OptionalRecommended. Unique identifier for a specific customer. Use this column if the demand driver is available at customer level.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Order Planning and Tracking

Prequisites before uploading your dataset

All content copied from https://docs.aws.amazon.com/.
