# Transactional data

###### Topics

- [Forecast](#forecast)

- [Sales history or demand](#demand)

- [Inventory level](#inventory-level)

- [Inbound orders](#in-flight-orders)

## Forecast

Supply Planning uses two different sources and types of forecast. You can use
the following source systems to retrieve forecast source:

- _External_ – Supply Planning uses the data that is being ingested
into the data lake forecast entity.

- _Demand Planning_ – Supply Planning uses the forecasts from Demand
Planning.

- _None_ – Supply Planning uses the sales or demand history data from
the outbound order line.

Supply Planning supports two types of forecast: deterministic and stochastic.
Deterministic forecasts contain only the mean of the forecast. Stochastic
forecasts contain P10/P50/P90, sometimes along with mean. When mean is not
provided with stochastic forecasts, Supply Planning uses P50(median) as
mean.

Each forecast record has four fields to represent the demand forecast:

- mean(double)

- p10(double)

- p50(also known as median, double)

- p90(double)

Based on the configured inventory policy, different fields in this entity are
required. For _sl_, p10/p50/90 is required; for
_doc\_fcst_, policy p50 or mean is required. Supply
Planning uses p50 as an approximation of the mean, and for
_doc\_dem_ and _abs\_level_, none of the
forecast fields are required.

**Daily planning**

Forecasts may be different for daily planning compared to weekly planning.
Here is an example of the daily and weekly planning forecast requirement.

![Daily planning](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/daily-planning.png)

**Weekly planning**

You can use the daily planning forecast example for weekly planning, or you
can also use the following example for weekly planning.

![Weekly planning](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/weekly-planning.png)

## Sales history or demand

Inventory policy _doc\_dem_ requires demand history to compute the historical average demand. Supply Planning gets the demand history from the _outbound\_order\_line_ entity under the _Outbound_ category.
Supply Planning uses the following fields:

- _ship\_from\_site\_id_(string)

- _product\_id_(string)

- _actual\_delivery\_date_(timestamp); when missing, use
_promised\_delivery\_date_(timestamp)

As part of the calculation, Supply Planning uses historical outbound order
lines with delivery dates in the past 30 days. The target field used for
quantity is _quantity\_delivered_; when missing,
use _quantity\_promised_. If _quantity\_promised_ is missing, then _final\_quantity\_requested_ will be used. If all are
missing, then _0_ will be used.

For example, if you use Supply Planning for product “laptop” at site “TX0” on
July 1, 2023, the record in _outbound\_order\_line_ where
_product\_id=laptop_, _ship\_from\_site\_id=TX0_, and _actual\_delivery\_date_ is from June 1, 2023 to June 30, 2023.
Supply Planning adds all the records and divides by 30 days to get the daily
demand.

## Inventory level

Supply Planning requires a beginning inventory level to start the planning
process. Supply Planning searches for the inventory level under the
_entity inv\_level_ data entity. Supply Planning searches
for a record with the following fields:

- _product\_id_

- _site\_id_

Supply Planning uses _on\_hand\_inventory_ to determine the
inventory level.

## Inbound orders

Supply Planning uses _inbound\_order\_line_ to retrieve the
in-flight order quantity. If an order is delivered during the planning horizon,
the quantity is considered as part of the existing supply.

Supply Planning searches for a record under
_inbound\_order\_line_ with the following fields:

- _order\_receive\_date_; when missing, use _expected\_delivery\_date_

- _product\_id_

- _to\_site\_id_

The following are the supported Order Types: PO (Purchase), TO (Transfer), and
MO (Production or Manufacturing).

Supply Planning uses the _quantity\_received_;
when missing, use _quantity\_confirmed_ then
_quantity\_submitted_ to determine the
on-order quantity.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Planning configuration data

N-Tier Visibility

All content copied from https://docs.aws.amazon.com/.
