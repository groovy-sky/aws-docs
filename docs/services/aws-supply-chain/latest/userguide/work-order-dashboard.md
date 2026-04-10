# Orders

You can view all the orders that are at-risk, delivered, early, late, on time, or watch. You can expand the order to view the materials under each order.

In the left navigation pane on the AWS Supply Chain dashboard, choose
**Order Planning and Tracking**. The **Order Planning and Tracking** page appears.

![Orders](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/Work_order.png)

Choose **Filters** to filter the orders based on **Country/Location**, **Campaign**, **Revision**
, **Main Work Center**, **Process Name**, and **Planner Group**. Once you set your filters, choose **Apply**. You can also choose **Save filter group** to save your filters.

You can also filter the orders by **All**, **On Time/Early**, **Watch**, **At Risk**, **Late**, **Delivered**, and **Site Delivery Forecast** status. For example, if you choose
**Late**, you will see all the orders that are currently late or delayed.

You can use the **Search** field to search by order or material number and use the _Sort_ option to sort the orders. You can sort them by any of the headers but by default, the orders are sorted first by
**Site Delivery Forecast** and second by **Order Priority**.

The **Orders** page, displays the following from your ERP or source system:

Orders columnDescriptionData entityColumn

Order

Display the order number. You can select the order to view your ERP or source system. You can expand each order to view the materials in the order.

process\_header

process\_id

Campaign/Revision

Displays the campaign and/or the revision of the order.

process\_header

program\_group

process\_header

revision

Main Work Center

Displays the main work center defined in the source system.

process\_header

execution\_group

Planner Group

Displays the planning group for each order.

process\_header

planning\_group

Order Description

Displays a brief reasoning of the order.

process\_header

description

Order End Date

Displays the date by which the order should me completed.

process\_header

planned\_completion\_date

Order Priority

Displays the priority of the order. AWS Supply Chain will only accept a numerical value for this field. For example, 1,2,3, and so on. If your ERP system doesn't contain a numerical value for this field, you will not be able to sort the order by priority.

process\_header

priority

Planned Start Date

The date when all the materials are required on-site before starting the work.

process\_header

planned\_start\_date

Flex 1 to 5

Custom fields that can be renamed and populated with any data.

process\_header

flex\_1, flex\_2, flex\_3, flex\_4, flex\_5

Recommendation

Displays all actionable items and is linked to a milestone. For example, if the order is blocked with a PO blocked milestone, the recommendation text will display to look for alternate products.

Calculated by Order Planning and TrackingCalculated by Order Planning and TrackingSite Delivery Forecast

Displays one of the following:

- **At risk** – Displayed when the material with the latest arrival date has a process that is either delayed or is in a blocked milestone. This item can still make the
required date and is displayed in Yellow.

- **Delivered** – Displayed after the last milestone of the last process is initiated indicating the completion of the process.

- **Early** – Displayed in green when all the order lines are early and includes the count of days of the earliest line.

- **Late** – Displayed when the order is running late due to the underlying order material with the latest delivery date estimated to arrive late. This item is displayed in Red.

- **On-time** – Displayed when the materials under the order is reaching the site within the required on-site date. This item is displayed in Green.

- **Watch** – Displayed when the material with the latest date is either blocked or late in a current supply chain process.

## Viewing order materials

You can view all the materials related to a order.

1. In the left navigation pane on the AWS Supply Chain dashboard, choose
    **Order Planning and Tracking**.

The **Order Planning and Tracking** page appears.

2. Use the expandable **Comments** feature to do the following:

- Add a comment (under 400 characters).

- Edit or delete a comment.

- See other users' comments.

![Comments feature on the Orders page](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/comments-orders.png)

3. Expand the order you would like to view.

The **Materials in Order** page appears.

Order LinesDescriptionData entityColumn

Material

Displays the material number.

process\_product

product\_id

Material Description

Provides a description of the material.

product

description

Quantity/UoM

Lists the quantity of the material. If UoM is available, UoM value is displayed. For example, 2 eaches.

reservation

quantity

quantity\_uom

Material Source

Displays if the material is in inventory or direct purchase.

site

description

inbound\_order

tpartner\_id

trading\_partner

description

Required on Site

Displays the date on which the material is required on-site.

process\_header

planned\_start\_date

process\_product

requested\_availability\_date

Brand name

Provides a name of the brand.

product

brand\_name

Product status

Provides the status of the product.

process\_product

status

Product type

Provides the type of the product.

process\_product

type

Reservation type

Provides the type of the reservation.

reservation

reservation\_type

Process product allocation type

Displays the allocation type for the product.

.

process\_product

overallocation

Process product allocation status

Displays the allocation status for the product.

.

process\_product

allocation\_status

Product flexible field 1 to 5

Custom fields that can be renamed and populated with any data.

process\_product

flex\_1, flex\_2, flex\_3, flex\_4, flex\_5

Reservation flexible field 1 to 5

Displays the reservation type of the product.

reservation

flex\_1, flex\_2, flex\_3, flex\_4, flex\_5

Revision

Displays the material revision.

process\_header

revision

Order type

Displays the order type.

process\_header

type

Current Process

Displays the current supply chain process for the order material.

Calculated by order planning and tracking.

Calculated by order planning and tracking.

Recommendation

Displays all actionable items and is linked to a milestone.

Site Delivery Forecast

Displays the site delivery forecast and status.

4. Choose the **Material** you would like to view in-detail. The
    **Material Summary** page appears and displays the
    summary of the material. You can use the same **Comments**
    feature mentioned in step 2 to add, update, and view comments.

![Order material summary - working forwards process](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/Insights_order.png)

![Order material summary - forecasted completion](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/working_backwards.png)

![Order material summary - working backwards process](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/working_backwards_new.png)

![Comments feature on the Procurement page](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/comments1.PNG)

![Comments feature on the Procurement page](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/comments2.PNG)

You can view the current milestone for the material and the recommendation AWS Supply Chain provides for each milestone.

MaterialDescriptionData entityColumn

Material name

Displays the name of the material.

product

description

Material

Provides a description of the material.

process\_product

product\_id

Quantity/UoM

Lists the quantity of the material. If UoM is available, UoM value is displayed. For example, 2 eaches.

reservation

quantity

reservation

quantity\_uom

Required on Site

Displays the date on which the material is required on-site.

process\_header

planned\_start\_date

process\_product

requested\_availability\_date

Vendor

Display the vendor from which the material is being procured.

inbound\_order

tpartner\_id

trading\_partner

description

PO Delivery Date

Displays the purchase order delivery date.

inbound\_order\_line

expected\_delivery\_date

Site Delivery Forecast

Displays the site delivery forecast and status.

Calculated by order planning and tracking.

Updated PO Delivery Date

Displays the updated PO delivery date.

Update Quantity

Displays the updated product quantity.

Supplier Delivery Date Confirmation

Displays the delivery date confirmation from the supplier.

Process product allocation type

Displays the allocation type for the product.

.

process\_product

allocation\_type

Process product allocation status

Displays the allocation status for the product.

.

process\_product

allocation\_status

Inventory Location

Displays the inventory location.

site

description

Inco Terms

Displays the incoterm code.

inbound\_order\_line

incoterm

Reservation Type

Displays the type of reservation.

reservation

reservation\_type

Brand Name

Displays the brand name of the product.

product

brand\_name

Product Status

Displays the product status.

process\_product

status

Product Type

Displays the product type.

process\_product

type

Campaign

Displays the campaign of the order.

process\_header

program\_group

Order

Display the order number. You can select the order to view your ERP or source system.

process\_product

process\_id

process\_header

process\_url

PR/Line Number

You can select the procurement or line number to view in your ERP or source system.

reservation

requisition\_id

reservation

requisition\_line\_id

inbound\_order\_line

inbound\_order\_line\_url

PO/Line Number

You can select the purchase order (PO) or line number to view in your ERP or source system.

reservation

order\_id

reservation

order\_line\_id

inbound\_order\_line

inbound\_order\_line\_url

STO/Line Number

You can select the STO or line number to view in your ERP or source system.

reservation

stock\_transfer\_1\_order\_id

reservation

stock\_transfer\_1\_order\_line\_id

reservation

stock\_transfer\_2\_order\_id

reservation

stock\_transfer\_2\_order\_line\_id

inbound\_order\_line

inbound\_order\_line\_url

RFQ/Line Number

You can select the RFQ or line number to view in your ERP or source system.

reservation

rfq\_id

reservation

rfq\_line\_id

inbound\_order\_line

inbound\_order\_line\_url

Product Type

Displays the type of the product.

product

product\_type

Currency UOM

Displays the currency unit of measure for the price and other economic variables of this product.

.

process\_product

currency\_uom

Danger

Displays the products that are hazardous.

product

un\_id

Hazmat Class

Displays the products that contain hazardous materials.

un\_details

un\_class

UN Class

Displays the products that are under the hazardous category.

un\_details

hazmat\_class

UN Description

Displays the description of the products that are under the hazardous category.

un\_details

un\_description

Image

Displays an image of the products that are under the hazardous category.

un\_details

image\_url

5. Choose **Copy shareable link to clipboard** to share the material summary dashboard.

6. Choose the **Edit** icon to edit the material summary view. Slide the data entity button to view the data field on the material summary page.

![Edit material summary page](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/edit_material_summary.png)

You can drag and drop the data entities to rearrange the date entity view on the material summary page.

7. Choose **Save Changes**.

8. Slide the **Show Completed Milestones** button to view all the completed milestones for a material.

![Viewing completed milestones](https://docs.aws.amazon.com/images/aws-supply-chain/latest/userguide/images/completed_milestones.png)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Organization Labels

Procurement

All content copied from https://docs.aws.amazon.com/.
