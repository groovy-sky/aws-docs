# Data Validation Rules

The validations performed prior to forecast creation are below. For more information, see [Demand Planning](required-entities.md).

Rule TypeRuleDatasetsDescriptionExport error records?Data Structure ValidationMandatory columns existence validationProduct, Outbound order line, Supplementary time series

Verifies presence of critical columns in datasets in required
datasets:

Outbound order line: product\_id, order\_date, final\_quantity\_requested

Product: id, description

Verifies presence of critical columns in recommended datasets, if
provided:

Supplementary Time Series: id, order\_date, time\_series\_name,
time\_series\_value

NoData Structure ValidationGranularity columns existence validationProduct, Outbound order line

Verifies presence of columns set as forecast granularity, if set in the
demand plan settings.

Outbound order line: product\_id, ship\_from\_site\_id, ship\_to\_site\_id,
ship\_to\_site\_address\_city, ship\_to\_address\_state, ship\_to\_address\_country,
channel\_id, customer\_tpartner\_id

Product: id, product\_group\_id, product\_type, brand\_name, color,
display\_desc, parent\_product\_id

NoData Structure ValidationActive product's history validationProduct, Outbound order line,Product AlternateVerifies that there is atleast one active product that has history on its own
or through product lineageNoData Quality ValidationMissing values in mandatory columns validationProduct, Outbound order line, Supplementary time seriesVerifies for null/empty values in mandatory columns specified in Mandatory
columns existence checkYesData Quality ValidationMissing values in granularity columns validationProduct, Outbound order lineVerifies for null/empty values in mandatory columns specified in Granularity
columns existence checkYesData Quality ValidationDate Range validationOutboundOrderLine, SupplementaryTimeSeriesThe order\_date column in the dataset must contain dates in a sane time range:
Anywhere from 01/01/1900 00:00:00 to 12/31/2050 00:00:00. YesForecasting Eligibility ValidationTimeseries per Predictor validationOutboundOrderLine

The timeseries per predictor must not exceed 5,000,000.

"Timeseries per predictor" is calculated by taking the count of unique
values for the product\_id column and each of the forecast granularity columns
and then taking the product of all those counts.

NoForecasting Eligibility ValidationCount of active products validationProductThe number of active products with records in the OOL dataset must not exceed
800,000.NoForecasting Eligibility ValidationHistorical data sufficiency validationOutbound order line

Verifies if at least one product in the dataset has sufficient historical
demand data to generate reliable forecasts

The forecast horizon must be no greater than 1/3 the time range in the
dataset (if training a new auto predictor) or 1/4 the time range in the dataset
(if training an existing auto predictor).

There is also a global maximum forecast horizon, which is 500.

NoForecasting Eligibility ValidationRow Count validationPartitioned OutboundOrderLineThe number of records in the partitioned OOL dataset must not exceed
3,000,000,000. There are certain forecast models that have smaller limits that are
checked here as well, if those models are being used.NoForecasting Eligibility ValidationMaximum Timeseries validationPartitioned OutboundOrderLine

The number of distinct timeseries must not exceed the model's limit, if
there is one.

"Distinct timeseries" is defined as the number of distinct rows in the
dataset when product\_id + all forecast granularity columns are
considered.

NoForecasting Eligibility Validation

Data Density validation

Partitioned OutboundOrderLine

The Data density of the dataset must be at least 5.

Data density is defined as (number of distinct products in the dataset) /
(total number of rows in the dataset). In other words it is "average rows per
product".

###### Note

The rule applies only when Prophet is selected as the forecasting algorithm.

No

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Data Validation Error Export

Demand Pattern and Recommendation

All content copied from https://docs.aws.amazon.com/.
