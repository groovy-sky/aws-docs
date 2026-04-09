# Example of modeling relational data in DynamoDB

This example describes how to model relational data in Amazon DynamoDB. The DynamoDB table design
corresponds to the relational order entry schema that is shown in [Relational modeling](bp-relational-modeling.md). This design
uses multiple specialized tables rather than a single adjacency list, providing clear operational
boundaries while leveraging strategic GSIs to serve all access patterns efficiently.

The design approach uses aggregate-oriented principles, grouping data based on access patterns
rather than rigid entity boundaries. Key design decisions include using separate tables for entities
with low access correlation, embedding related data when always accessed together, and using
item collections for identifying relationships.

The following tables and their accompanying indexes support the relational order entry schema:

## Employee Table Design

The Employee table stores employee information as a single entity per item, optimized for direct employee lookups and supporting multiple query patterns through strategic GSIs. This table demonstrates the principle of designing separate tables for entities with independent operational characteristics and low cross-entity access correlation.

The table uses a simple partition key (employee\_id) without a sort key, as each employee is a distinct entity. Four GSIs enable efficient querying by different attributes:

- _EmployeeByName GSI_ \- Uses INCLUDE projection with all employee attributes to support complete employee detail retrieval by name, handling potential duplicate names with employee\_id as sort key

- _EmployeeByWarehouse GSI_ \- Uses INCLUDE projection with only essential attributes (name, job\_title, hire\_date) to minimize storage costs while supporting warehouse-based queries

- _EmployeeByJobTitle GSI_ \- Enables role-based queries with INCLUDE projection for reporting and organizational analysis

- _EmployeeByHireDate GSI_ \- Uses a static partition key value "EMPLOYEE" with hire\_date as sort key to enable efficient date range queries for recent hires. Since employee additions/updates are typically under 1,000 WCU, a single partition can handle the write load without hot partition issues

Employee Table - Base Table Structureemployee\_id (PK)namephone\_numberswarehouse\_idjob\_titlehire\_dateentity\_typeemp\_001John Smith\["+1-555-0101"\]wh\_seaManager2024-03-15EMPLOYEEemp\_002Jane Doe\["+1-555-0102", "+1-555-0103"\]wh\_seaAssociate2025-01-10EMPLOYEEemp\_003Bob Wilson\["+1-555-0104"\]wh\_pdxAssociate2025-06-20EMPLOYEEemp\_004Alice Brown\["+1-555-0105"\]wh\_pdxSupervisor2023-11-05EMPLOYEEemp\_005Charlie Davis\["+1-555-0106"\]wh\_seaAssociate2025-12-01EMPLOYEE

EmployeeByName GSI - Supporting Employee Name Queriesname (GSI-PK)employee\_id (GSI-SK)phone\_numberswarehouse\_idjob\_titlehire\_dateAlice Brownemp\_004\["+1-555-0105"\]wh\_pdxSupervisor2023-11-05Bob Wilsonemp\_003\["+1-555-0104"\]wh\_pdxAssociate2025-06-20Charlie Davisemp\_005\["+1-555-0106"\]wh\_seaAssociate2025-12-01Jane Doeemp\_002\["+1-555-0102", "+1-555-0103"\]wh\_seaAssociate2025-01-10John Smithemp\_001\["+1-555-0101"\]wh\_seaManager2024-03-15

EmployeeByWarehouse GSI - Supporting Warehouse Querieswarehouse\_id (GSI-PK)employee\_id (GSI-SK)namejob\_titlehire\_datewh\_pdxemp\_003Bob WilsonAssociate2025-06-20wh\_pdxemp\_004Alice BrownSupervisor2023-11-05wh\_seaemp\_001John SmithManager2024-03-15wh\_seaemp\_002Jane DoeAssociate2025-01-10wh\_seaemp\_005Charlie DavisAssociate2025-12-01

EmployeeByJobTitle GSI - Supporting Job Title Queriesjob\_title (GSI-PK)employee\_id (GSI-SK)namewarehouse\_idhire\_dateAssociateemp\_002Jane Doewh\_sea2025-01-10Associateemp\_003Bob Wilsonwh\_pdx2025-06-20Associateemp\_005Charlie Daviswh\_sea2025-12-01Manageremp\_001John Smithwh\_sea2024-03-15Supervisoremp\_004Alice Brownwh\_pdx2023-11-05

EmployeeByHireDate GSI - Supporting Recent Hire Queriesentity\_type (GSI-PK)hire\_date (GSI-SK)employee\_idnamewarehouse\_idEMPLOYEE2023-11-05emp\_004Alice Brownwh\_pdxEMPLOYEE2024-03-15emp\_001John Smithwh\_seaEMPLOYEE2025-01-10emp\_002Jane Doewh\_seaEMPLOYEE2025-06-20emp\_003Bob Wilsonwh\_pdxEMPLOYEE2025-12-01emp\_005Charlie Daviswh\_sea

## Customer Table Design

The Customer table maintains customer information with strategic denormalization of account\_rep\_id to enable efficient account representative queries. This design choice trades slight storage overhead for query performance, eliminating the need for joins between customer and account representative data.

The table supports multiple phone numbers per customer using a list attribute, demonstrating DynamoDB's schema flexibility. The single GSI enables account representative workflows:

- _CustomerByAccountRep GSI_ \- Uses INCLUDE projection with name and email attributes to support account rep customer management without requiring full customer record retrieval

Customer Table - Base Table Structurecustomer\_id (PK)namephone\_numbersemailaccount\_rep\_idcust\_001Acme Corp\["+1-555-1001"\]contact@acme.comrep\_001cust\_002TechStart Inc\["+1-555-1002", "+1-555-1003"\]info@techstart.comrep\_001cust\_003Global Traders\["+1-555-1004"\]sales@globaltraders.comrep\_002cust\_004BuildRight LLC\["+1-555-1005"\]orders@buildright.comrep\_002cust\_005FastShip Co\["+1-555-1006"\]support@fastship.comrep\_003

CustomerByAccountRep GSI - Supporting Account Rep Queriesaccount\_rep\_id (GSI-PK)customer\_id (GSI-SK)nameemailrep\_001cust\_001Acme Corpcontact@acme.comrep\_001cust\_002TechStart Incinfo@techstart.comrep\_002cust\_003Global Traderssales@globaltraders.comrep\_002cust\_004BuildRight LLCorders@buildright.comrep\_003cust\_005FastShip Cosupport@fastship.com

## Order Table Design

The Order table uses vertical partitioning with separate items for order headers and order items. This design enables efficient product-based queries while maintaining all order components within the same partition for efficient access. Each order consists of multiple items:

- _Order Header_ \- Contains order metadata with PK=order\_id, SK=order\_id

- _Order Items_ \- Individual line items with PK=order\_id, SK=product\_id, enabling direct product queries

###### Note

This vertical partitioning approach trades the simplicity of embedded order items for enhanced query flexibility. Each order item becomes a separate DynamoDB item, enabling efficient product-based queries while maintaining all order data within the same partition for efficient retrieval in a single request.

The table includes strategic denormalization of account\_rep\_id (duplicated from Customer table) to enable direct account representative queries without requiring customer lookups. For high-throughput write scenarios, OPEN orders include status and shard attributes to enable write sharding across multiple partitions.

Four GSIs support different query patterns with optimized projections:

- _OrderByCustomerDate GSI_ \- Uses INCLUDE projection with order summary and item details to support customer order history with date range filtering

- _OpenOrdersByDate GSI (Sparse, Sharded)_ \- Uses multi-attribute partition key (status + shard) with 5 shards to distribute 5,000 WPS (writes per second) across partitions (1,000 WPS each, matching DynamoDB's 1,000 WCU per partition limit). Only indexes OPEN orders (20% of total), which can help reduce GSI storage costs. Requires parallel queries across all 5 shards with client-side result merging

- _OrderByAccountRep GSI_ \- Uses INCLUDE projection with order summary attributes to support account representative workflows without full order details

- _ProductInOrders GSI_ \- Created from OrderItem records (PK=order\_id, SK=product\_id), this GSI enables queries to find all orders containing a specific product. Uses INCLUDE projection with order context (customer\_id, order\_date, quantity) for product demand analysis

Order Table - Base Table Structure (Vertical Partitioning)PKSKcustomer\_idorder\_datestatusaccount\_rep\_idquantitypriceshardord\_001ord\_001cust\_0012025-11-15CLOSEDrep\_001ord\_001prod\_100525.00ord\_002ord\_002cust\_0012025-12-20OPENrep\_0010ord\_002prod\_1011015.00ord\_003ord\_003cust\_0022026-01-05OPENrep\_0012ord\_003prod\_100325.00

OrderByCustomerDate GSI - Supporting Customer Order Queriescustomer\_id (GSI-PK)order\_date (GSI-SK)order\_idstatustotal\_amountorder\_itemsshardcust\_0012025-11-15ord\_001CLOSED225.00\[{product\_id: "prod\_100", qty: 5}\]cust\_0012025-12-20ord\_002OPEN150.00\[{product\_id: "prod\_101", qty: 10}\]0cust\_0022026-01-05ord\_003OPEN175.00\[{product\_id: "prod\_100", qty: 3}\]2cust\_0032025-10-10ord\_004CLOSED250.00\[{product\_id: "prod\_101", qty: 5}\]cust\_0042026-01-03ord\_005OPEN200.00\[{product\_id: "prod\_100", qty: 20}\]1

OpenOrdersByDate GSI (Sparse, Sharded) - Supporting High-Throughput Open Order Queriesstatus (GSI-PK-1)shard (GSI-PK-2)order\_date (SK)order\_idcustomer\_idaccount\_rep\_idorder\_itemstotal\_amountOPEN02025-12-20ord\_002cust\_001rep\_001\[{product\_id: "prod\_101", qty: 10}\]150.00OPEN12026-01-03ord\_005cust\_004rep\_002\[{product\_id: "prod\_100", qty: 20}\]200.00OPEN22026-01-05ord\_003cust\_002rep\_001\[{product\_id: "prod\_100", qty: 3}\]175.00

OrderByAccountRep GSI - Supporting Account Rep Order Queriesaccount\_rep\_id (GSI-PK)order\_date (GSI-SK)order\_idcustomer\_idstatustotal\_amountrep\_0012025-11-15ord\_001cust\_001CLOSED225.00rep\_0012025-12-20ord\_002cust\_001OPEN150.00rep\_0012026-01-05ord\_003cust\_002OPEN175.00rep\_0022025-10-10ord\_004cust\_003CLOSED250.00rep\_0022026-01-03ord\_005cust\_004OPEN200.00

ProductInOrders GSI - Supporting Product Order Queriesproduct\_id (GSI-PK)order\_id (GSI-SK)customer\_idorder\_datequantityprod\_100ord\_001cust\_0012025-11-155prod\_100ord\_003cust\_0022026-01-053prod\_101ord\_002cust\_0012025-12-2010

## Product Table Design

The Product table uses the item collection pattern to store both product metadata and inventory data within the same partition. This design leverages the identifying relationship between products and inventory - inventory cannot exist without a parent product. Using PK=product\_id with SK=product\_id for product metadata and SK=warehouse\_id for inventory items eliminates the need for a separate Inventory table and GSI, reducing costs by approximately 50%.

This pattern enables efficient queries for both individual warehouse inventory (GetItem with composite key) and all warehouse inventory for a product (Query on partition key). The total\_inventory attribute in the product metadata item provides denormalized aggregation for quick total inventory lookups.

Product Table - Base Table Structure (Item Collection Pattern)product\_id (PK)warehouse\_id (SK)product\_namecategoryunit\_priceinventory\_quantitytotal\_inventoryprod\_100prod\_100Widget AHardware25.00500prod\_100wh\_sea200prod\_100wh\_pdx150prod\_100wh\_atl150prod\_101prod\_101Gadget BElectronics50.00300prod\_101wh\_sea100prod\_101wh\_pdx200

Each table is designed with specific Global Secondary Indexes (GSIs) to support the required
access patterns efficiently. The design uses aggregate-oriented principles with strategic denormalization
and sparse indexing to optimize both performance and cost.

Key design optimizations include:

- _Sparse GSI_ \- OpenOrdersByDate only indexes OPEN orders (20% of total), which can help reduce GSI storage costs

- _Item Collection Pattern_ \- Product table stores inventory using PK=product\_id, SK=warehouse\_id to eliminate separate inventory table

- _Order + OrderItems Aggregation_ \- Embedded as single item due to 100% access correlation

- _Strategic Denormalization_ \- account\_rep\_id duplicated in Order table for efficient queries

Finally, you can revisit the access patterns that were defined earlier. The following table shows
how each access pattern is efficiently supported using the multi-table design with strategic GSIs.
Each pattern uses either direct key lookups or single GSI queries, avoiding expensive scans and
providing consistent performance at any scale.

S. No.Access patternsQuery conditions

1

Look up Employee Details by Employee ID

Employee Table: GetItem(employee\_id="emp\_001")

2

Query Employee Details by Employee Name

EmployeeByName GSI: Query(name="John Smith")

3

Find an Employee's Phone Number(s)

Employee Table: GetItem(employee\_id="emp\_001")

4

Find a Customer's Phone Number(s)

Customer Table: GetItem(customer\_id="cust\_001")

5

Get Orders for Customer within Date Range

OrderByCustomerDate GSI: Query(customer\_id="cust\_001", order\_date BETWEEN "2025-01-01" AND "2025-12-31")

6

Show all Open Orders within Date Range

OpenOrdersByDate GSI: Query 5 shards in parallel with multi-attribute PK (status="OPEN" + shard=0-4), SK=order\_date BETWEEN "2025-01-01" AND "2025-12-31", merge results

7

See all Employees hired recently

EmployeeByHireDate GSI: Query(entity\_type="EMPLOYEE", hire\_date >= "2025-01-01")

8

Find all Employees in Warehouse

EmployeeByWarehouse GSI: Query(warehouse\_id="wh\_sea")

9

Get all Items on Order for Product

ProductInOrders GSI: Query(product\_id="prod\_100")

10

Get Inventories for Product at all Warehouses

Product Table: Query(product\_id="prod\_100")

11

Get Customers by Account Rep

CustomerByAccountRep GSI: Query(account\_rep\_id="rep\_001")

12

Get Orders by Account Rep

OrderByAccountRep GSI: Query(account\_rep\_id="rep\_001")

13

Get Employees with Job Title

EmployeeByJobTitle GSI: Query(job\_title="Manager")

14

Get Inventory by Product and Warehouse

Product Table: GetItem(product\_id="prod\_100", warehouse\_id="wh\_sea")

15

Get Total Product Inventory

Product Table: GetItem(product\_id="prod\_100", warehouse\_id="prod\_100")

[Document Conventions](../../../../general/latest/gr/docconventions.md)

First steps

Migrating to DynamoDB

All content copied from https://docs.aws.amazon.com/.
