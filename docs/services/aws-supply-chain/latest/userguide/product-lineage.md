# Product lineage

_Product lineage_ refers to the relationship established between
products and their previous versions or alternate products. Demand Planning uses product lineage information to create surrogate histories for these products, which serve as forecast inputs for demand predictions.

Product lineage supports the following patterns:

- A single product has one lineage or alternate product = 1:1

![Product lineage pattern = 1:1](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/product_lineage_pattern1.png)

The following example shows an 1:1 scenario.

![Product lineage pattern = 1:1](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/1%20is%20to%201_example.png)

- A single product has more than one product as lineage or alternate = Many:1

![Product lineage pattern = Many:1](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/product_lineage_pattern2.png)

Demand Planning supports product lineage relationship modeled as both _chain_ or _flattened_ methods.

- **Chain format** – You can directly model lineage relationships like A to B and B to C. In the following example. Demand Planning will model the lineage relationship as A to B, B to C, and A to C.

PredecessorSuccessor

A

B

B

C

The following example shows an Many:1 scenario - Chain format

![Product lineage pattern = Chain format](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/chain_format.png)

- **Flattened format** – Demand Planning will continue to support lineage information in A to B and A to C format. In the following example, Demand planning will model the lineage relationship as A to B and A to C. B to C is not considered.

PredecessorSuccessor

A

B

A

C

###### Note

Chain format only supports 6 levels of lineage relationship. If you have more than 6, you can use flattened format to model the lineage relationship.

The following example shows an Many:1 scenario - Flattened format

![Product lineage pattern = Flattened format](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/1%20is%20to%20many_example.png)

- A single product can be lineage or alternate for more than 1 product = 1 : Many

![Product lineage pattern = 1:Many](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/product_lineage_pattern3.png)

To enable the product lineage feature, you can define the lineage relationship for the
different versions of the products or alternates/substitutes in the
_product\_alternate_ data entity. For more information, see [Demand Planning](required-entities.md).

If your instance was created on or after September 11, 2023, you will see
_product\_alternate_ data entity in the AWS Supply Chain data Connection
module. If your instance was created before September 11, 2023, create a new data connection
to enable the _product\_alternate_ data entity for ingestion.

To ingest data into the _product\_alternate_ data entity, follow the
guidelines below:

- _product\_id_ – The primary product to create the
forecast.

- _alternative\_product\_id_ – Previous version of the product
or alternate/substitute product.

To consider multiple _alternative\_product\_id_ for a single
_product\_id_, enter them in separate rows.

- Demand Planning will consider the data ONLY when the values are provided in the
following format.

- _alternate\_type_ is _similar\_demand\_product_.

- _status_ is _active_.

- _alternate\_product\_qty\_uom_ is the text
_percentage_.

- _alternate\_product\_qty_ – Enter the proportion of
history of the alternate product you want to use for forecasting new products in the
_alternate\_product\_qty_ data field. For example, if it is 60%,
enter 60. When you have multiple _alternative\_product\_id_ for a
single _product\_id_, the _alternate\_product\_qty_
does not have to add up to 100.

- The _eff\_start\_date_ and _eff\_end\_date_ data
fields are required. However, you can leave this field empty and Demand Planning will auto-fill with 1000 and 9999 years respectively.

When the forecast is created using product lineage data, you will see an indicator
_Forecast is based on alternate product's history_ on the Demand Planning
page when you filter by _product ID_.

The following table shows an example of how Demand Planning Product lineage feature works
based on the data ingested into the _product\_alternate_ data entity.

ColumnRequired or OptionalExample 1Example 2Example 3Example 4Example 5Example 6Example 7Example 8Example 9Example 10Example 11

product\_id

RequiredProduct 123Product 123Product 123Product 123Product 123Product 123Product 123Product 123Product 123NullProduct 123

alternative\_product\_id

RequiredProduct XYZNullProduct XYZProduct XYZProduct XYZProduct XYZProduct XYZProduct XYZProduct XYZNullProduct XYZ

alternate\_type

RequiredSimilar\_Demand\_ProductSimilar\_Demand\_ProductNull or a different valueSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_ProductSimilar\_Demand\_Product

status\*

RequiredactiveactiveactiveinactiveactiveactiveNullactiveactiveactiveactive

alternate\_product\_qty

Required10060100100Null10010010010010060

alternate\_product\_qty\_uom

RequiredpercentagepercentagepercentagepercentagepercentageNull or a different valuepercentagepercentagepercentagepercentagepercentage

eff\_start\_date

Required2023-01-01 00:00:002023-01-01 00:00:002023-01-01 00:00:002023-01-01 00:00:002023-01-01 00:00:002023-01-01 00:00:002023-01-01 00:00:00Null2023-01-01 00:00:002023-01-01 00:00:00Null

eff\_end\_date

Required2025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:592025-12-31 23:59:59Null2025-12-31 23:59:59Null

**Expected behavior**

NA100% of product XYZ's history from 1/1/2023 to 31/12/2025 will be used to
forecast product 123.Invalid mapping since alternative\_product\_id is missing.Invalid mapping since alternate \_type is not 'similar\_demand\_product'.Inactive mapping.Invalid mapping since alternate\_product\_qty is missing.Invalid mapping since alternate\_product\_qty\_uom is missing or not
percentage.Invalid mapping since status is missing.Ingestion will fail.Ingestion will fail.Invalid mapping since product\_id and alternative\_product\_id are missing.Ingestion will fail.

NANANANANANANANA

Demand Planning will auto-populate the _eff\_start\_date_ to year 1000. This scenario is valid and data ingestion will not fail.

Demand Planning will auto-populate the _eff\_end\_date_ to year 9999. This scenario is valid and ingestion will not fail.

NA

Demand Planning will auto-populate the _eff\_start\_date_ to year 1000 and _eff\_end\_date_ to year 9999. This scenario is valid and ingestion will not fail.

The following example explains how Demand Planning will interpret when the _status_ is set as _inactive_ and the product lineage is in chain format.

ColumnColumnStatus

A

B

Active

B

C

Inactive

C

D

Active

Demand planing considers the status of the first root and child mapping as the status for the entire chain.

A to B Active

A to C Active

A to D Active

B to C Inactive

B to D Inactive

C to D Active

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Demand driver recommendations

Product lifecycle

All content copied from https://docs.aws.amazon.com/.
