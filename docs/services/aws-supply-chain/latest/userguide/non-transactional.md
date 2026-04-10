# Planning configuration data

This section lists all the required fields used by Supply Planning and describes how
each field is used. For information on data fields required for Supply Planning, see
[Supply Planning](entities-supply-planning.md).

###### Topics

- [Product](#product)

- [Site](#site)

- [Trading partner](#trading-partners)

- [Vendor product](#vendor-product)

- [Vendor lead time](#vendor-leadtime)

- [Sourcing rule](#sourcing-rule)

- [Inventory policy](#inventory-policy)

- [Sourcing schedule](#sourcing-schedule)

- [Bill of Material (BOM)](#product-bom)

- [Production process](#production-process)

- [Supply planning parameters](#production-process2)

- [Transactional data](transactional.md)

## Product

The product entity defines the list of items or products that must be included
in the planning. The purchase order requests use _unit\_cost_
_field_ from the _Product_ entity to determine
the order value or amount. The _Product_ entity also contains
the product group corresponding to a specific product, which is a foreign key
into a _product\_hierarchy_ entity. Product groups can be used
in configuring inventory policies, sourcing schedules, lead times, and so on, at
the aggregate level.

## Site

The _Site_ entity defines the list of sites or locations
that must be included in the planning. The _Site_ entity also
contains Regions corresponding to a specific site, which is a foreign key into a
Geography entity. Regions can be used in configuring inventory policies, sourcing
schedules, lead times, and so on, at the aggregate level.

## Trading partner

The _Trading\_partner_ entity defines the list of suppliers.
_tpartner\_type_ should be set to
_Vendor_ when uploading supplier information.

## Vendor product

Products supplied by each supplier are defined in the
_vendor\_product_ entity. This entity also contains
vendor-specific cost information.

## Vendor lead time

Vendor lead time is the time period between placing an order to a vendor and
receiving the order. This data is defined in the _VendorMgmt_
category under the _vendor\_lead\_time_ data entity. Vendor
lead time follows the following override logic:

- Product level vendor lead time overrides product group level vendor lead time.

- Site level vendor lead time overrides region level vendor lead time.

- Region level vendor lead time overrides company level vendor lead time.

To look for a record, Supply Planning uses the following fields:

- company\_id

- region\_id

- site\_id

- product\_group\_id

- product\_id

The following is an example of the override logic:

![Override logic example](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/override_logic.png)

The following is an example of how Supply Planning calculates vendor lead time:

![Vendor lead time calculation](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/vendor_lead_time.png)

Prioritization order is _product_ \> _product\_group_ \> _site_ \> _dest\_geo (region)_ \> _product segment_ \> _company_.

## Sourcing rule

Supply Planning generates a plan based on the supply chain network topology defined under the _sourcing\_rules_ entity.

The supported sourcing rule types are transfer, buy, and manufacture.

Sourcing rules follow the _product\_id_ >
_product\_group\_id_ \> _company\_id_
override logic.

Supply Planning retrieves the transportation lead time by referencing _transportation\_lane\_id_ and
accessing _transit\_time_ in _transportation\_lane_.
There are two steps to retrieve the transfer lead time.

1. Find _transportation\_lane\_id_ in _sourcing\_rules_. Only the sourcing rules that have both
    _to\_site\_id_ and _from\_site\_id_ are eligible for retrieving
    _transfer\_lead\_time_.

2. Use _transportation\_lane\_id_ to look up _transportation\_lane_.

When there are multiple records with the same _to\_site\_id_ and _product\_id_
( _product\_group\_id_) in the _sourcing\_rule_ entity, only the records with the highest
priority (the smallest number) will be used.

Sourcing rules example:

Based on the preceding definition, Supply Planning selects the following
sourcing rule SR1: Laptop at site `TX0` is sourced from site `IL0`
via `transportation_lane_9`.

sourcing\_rule\_id

product\_id

product\_group\_id

sourcing\_rule\_type

from\_site\_id

to\_site\_id

sourcing\_priority

transportation\_lane\_id

SR1

laptop

electronics

transfer

IL0

TX0

1

transportation\_lane\_9

SR2

laptop

electronics

transfer

NJ1

TX0

2

transportation\_lane\_21

SR3

laptop

electronics

transfer

IL0

TX0

1

transportation\_lane\_11

When multiple records with the same priority exist for the same combination of _to\_site\_id_,
_product\_id_ (or _product\_group\_id_), the reorder quantity will be distributed among the available
sourcing options based on the _sourcing\_ratio_ field.
Note that multiple sourcing is currently only supported for the `buy` sourcing rule type.

Multi-sourcing example:

sourcing\_rule\_id

product\_id

product\_group\_id

sourcing\_rule\_type

tpartner\_id

to\_site\_id

sourcing\_priority

sourcing\_ratio

SR1

laptop

electronics

buy

supplier1

TX0

1

4

SR2

laptop

electronics

buy

supplier2

TX0

1

6

Both sourcing rules, SR1 and SR2, are selected,
and the order quantity will be allocated between Supplier 1 and Supplier 2 in a 4:6 ratio.

## Inventory policy

Supply Planning searches for a record in the dataset by using the following
fields:

- _site\_id_

- _geodesic_

- _company\_id_

- _product\_id_

- _product\_group\_id_

- _segment\_id_

Supply Planning uses _ss\_policy_ to determine the inventory
policy. The override logic uses the following priority:
_product\_id_ \> _product\_group\_id_ >
_site\_id_ \>  and
_dest\_geo\_id_ \> _segment\_id_ >
_company\_id_.

The supported _ss\_policy_ values are
_abs\_level_, _doc\_dem_, _doc\_fcst_, and
_sl_.

The following example displays the override priority logic.

![Override logic](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/override1.png)

The following is an example of the _ss\_policy_ value based
on the override logic.

![Override ride logic example for ss_policy value](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/override2.png)

## Sourcing schedule

###### Note

Sourcing schedule is an optional entity. If this entity is not provided, Supply Planning uses
a continuous review process to generate _required\_date_
based on when products are needed.

Supply Planning uses sourcing schedule to generate purchase plans by using the
following steps:

- Find _sourcing\_schedule\_id_ in _sourcing\_schedule_.

- Find the schedule by _using sourcing\_schedule\_id_ in
_sourcing\_schedule\_details_.

Supply Planning searches for the following fields in _sourcing\_schedule\_id_ under _sourcing\_schedule_.

- _to\_site\_id_

- _tpartner\_id_ or _from\_site\_id_

Based on the sourcing path in sourcing rules, Supply Planning determines
whether to use _from\_site\_id_ or _tpartner\_id_. Supply Planning reads the value in the
_sourcing\_schedule\_id_ field to determine the next
step.

Supply Planning reads the schedule details under _sourcing\_schedule\_details_ with the following fields:

- _sourcing\_schedule\_id_

- _company\_id_

- _product\_group\_id_

- _product\_id_

_sourcing\_schedule\_details_ follows the
override logic, _product\_id_ \> _product\_group\_id_ \> _company\_id_.

The following is an example of the override logic in _sourcing\_schedule\_details_.

![Sourcing schedule override logic](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/sourcing_schedule2.png)

The following are the selected schedules after applying the override
logic.

![Sourcing schedule override logic](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/sourcing_schedule3.png)

The actual schedule can be from one row to multiple rows, based on the
complexity of the schedule. For the field _week\_of\_month_, only one number is allowed in each row. For
multiple weeks of the month, multiple records are required (see the following
example). For the field _day\_of\_week_, both
integer and name of day are allowed (Sun: 0, Mon: 1, Tue: 2, Wed: 3, Thu: 4,
Fri: 5, Sat: 6). In the sourcing schedule details, weekly planning requires
_week\_of\_month_. While in daily planning,
_week\_of\_month_ can be empty, which means
every week. See the following examples.

![Sourcing schedule override logic](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/sourcing_schedule4.png)

Note that for weekly planning, _week\_of\_month_ is required if _day\_of\_week_ is provided.

The following example shows the dates that can be used for daily
planning.

DateDay of the weekWeek of the month

8/1/2023

NA

NA

8/12/2023

NA

NA

NA

2

NA

NA

5

NA

The following example can be used for both daily and weekly planning.

DateDay of the weekWeek of the month

8/1/2023

NA

NA

8/12/2023

NA

NA

NA

2

1

NA

2

2

NA

2

3

NA

2

4

NA

2

5

NA

5

1

NA

5

2

NA

5

3

NA

5

4

NA

5

5

## Bill of Material (BOM)

Product BOM is used in Manufacturing Plans when _sourcing\_rule_ is set to Manufacture. For information on how to ingest Product BOM, see the AWS Supply Chain API Reference document.

## Production process

_production\_process\_id_ is referenced in the
_sourcing\_rule_ and _product\_bom_
entities. These fields are used to consume lead time information to make or
assemble a BOM.

## Supply planning parameters

In _supply\_planning\_parameters_ entity,
_planner\_name_ of the supply planner can be assigned at
_product\_id_ level. Planner name will be displayed on the
planned orders generated by the supply planning engine.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Business workflow

Transactional data

All content copied from https://docs.aws.amazon.com/.
