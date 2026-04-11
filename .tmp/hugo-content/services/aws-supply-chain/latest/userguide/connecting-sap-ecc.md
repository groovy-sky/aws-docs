# Connecting to SAP ECC 6.0

To extract your data from SAP ECC 6.0, follow the procedure below.

01. On the AWS Supply Chain dashboard, on the left navigation pane, choose
     **Data Lake**.

02. On the **Data lake** page, choose **Add New Source**.

    The **Select your supply chain data source** page appears.

03. Choose **SAP ECC**.

04. Under **SAP ECC Connection Details**, enter the
     following:

- **Connection name** – Enter a name for your
connection. Connection names can only contain letters, numbers, and
dashes.

- **Connection description** – Enter a
description for your connection.

05. Under **Amazon S3 Bucket Billing**, review the Amazon S3 billing
     information, and then select **Acknowledge**.

06. Choose **Next**.

07. Under **Data Mapping**, choose **Get**
    **started**.

08. ###### Note

    The required fields are already mapped. Perform this step only if you want
    to make specific changes to the default transformation recipe.

    On the **Mapping Recipe** page, you can view the default
     transformation recipe under **Field mappings**.

    Choose **Add mapping** to map any additional destination
     field. The **Required Destination Fields** are mandatory.
     Choose **Destination field** to add an additional custom
     destination field.

09. ###### Note

    You can only use AWS Glue DataBrew to edit the recipes for transactional entities.
    Use AWS Supply Chain to download your recipes, and edit them in DataBrew. Then upload
    the recipes back into AWS Supply Chain. You can't use the AWS Supply Chain web application
    to edit the transactional data fields in a recipe.

     (Optional) Under **Recipe Actions**, you can do the
     following:

- **Download recipe file** \- Select
**Download** to edit your recipe files offline with
DataBrew.

- **Upload recipe file** \- Choose
**browse files**, or move (drag and drop) your
edited recipe files. Select **Confirm upload** to
upload the edited recipe file and modify your data field
mappings.

- **Reset to default recipe** \- Select
**Yes, reset my recipe** to remove all your custom
mappings and revert to the default recipe recommended by
AWS Supply Chain.

10. To edit your source field mappings and validate your transformation recipe,
     you can upload sample data. On the **Mapping Recipe** page,
     under **Upload sample data**, choose **browse**
    **files**, or move (drag and drop) files. The sample data file
     must contain the required parameters and include the source field
     names.

11. Choose **Accept all and continue**.

12. Under **Review and confirm**, you can view the data
     connection summary. To edit your data field mapping, choose **Go back to**
    **Data Mapping**.

13. To review the Amazon S3 paths where you must upload your SAP source data for
     ingestion, choose **Confirm and configure data ingestion**.
     Alternatively, you can choose **Confirm and configure data ingestion**
    **later**. You can view the data ingestion information anytime. From
     the AWS Supply Chain dashboard, select **Connections**. Select
     the connection dataflow that you want to ingest data, choose the vertical
     ellipsis, and select **Ingestion setup**.

14. If you're not using the Amazon S3 API to ingest data, create the Amazon S3 path manually
     on the Amazon S3 console. For more information about how to create paths, see
     [Uploading data to an Amazon S3 bucket](manually-uploading-data.md).

15. Review the following table to map the AWS Supply Chain data entity with SAP
     source.

    ###### Important

    On the **Amazon S3 path** page, you must upload the
    parent entity before the child entity. You can first upload all the parent
    entities and then upload all the child entities together.

Data entitySAP sourceHierarchyData entity action

Company – [company](organization-company-entity.md)

0COMP\_CODE\_TEXT

Parent

Replace

Geography – [geography](organization-geography-entity.md)

ADRC

Parent

Replace

Inventory – [inv\_level](inventory-mgmnt-inv-level-entity.md)

MARD

Parent

Update

MCHB

Parent

Update

VBBE

Child

Update

Inventory – [inv\_policy](planning-inv-policy-entity.md)

MARC

Parent

Replace

0MATERIAL\_ATTR

Child

Update

Outbound – [outbound\_order\_line](outbound-fulfillment-order-line-entity.md)

2LIS\_11\_VAITM

Parent

Update

0BP\_DEF\_ADDRESS\_ATTR

Child

Update

0MATERIAL\_ATTR

Child

Update

2LIS\_11\_VAHDR

Child

Update

Outbound – [outbound\_shipment](outbound-fulfillment-shipment-entity.md)

2LIS\_08TRTLP

Parent

Update

2LIS\_08TRFKP

Child

Update

2LIS\_08TRTK

Child

Update

2LIS\_12\_VCITM

Child

Update

Product – [product](product-product-entity.md)

0MATERIAL\_ATTR

Parent

Replace

0MATERIAL\_TEXT

Child

Update

Product – [product\_hierarchy](product-hierarchy-entity.md)

T179

Parent

Replace

Purchase order – [inbound\_order](replenishment-inbound-order-entity.md)

2LIS\_02\_HDR

Parent

Update

CDHDR

Child

Update

EKKO

Child

Update

Purchase order – [inbound\_order\_line](replenishment-inbound-order-line-entity.md)

2LIS\_02\_ITM

Parent

Update

0MATERIAL\_ATTR

Child

Update

2LIS\_03\_BF

Child

Update

EKPO

Child

Update

LIPS

Child

Update

LIKP

Child

Update

INB-SHIPMENT

Child

Update

Purchase order – [inbound\_order\_line\_schedule](replenishment-inbound-order-line-schedule-entity.md)

2LIS\_02\_SCL

Parent

Update

2LIS\_02\_SCN

Child

Update

Production order – [inbound\_order](replenishment-inbound-order-entity.md)

2LIS\_04\_P\_MATNR

Parent

Update

Production order – [inbound\_order\_line](replenishment-inbound-order-line-entity.md)

2LIS\_04\_P\_MATNR

Parent

Update

0CO\_PC\_ACT\_05

Child

Update

0MATERIAL\_ATTR

Child

Update

Reference – [reference\_field](reference-fields-entity.md)

0PURCH\_ORG\_TEXT

Parent

Update

MDRP\_NODTT

Parent

Update

T005T

Parent

Update

T141T

Parent

Update

T173T

Parent

Update

T179T

Parent

Update

T370U

Parent

Update

T618T

Parent

Update

Shipment – [shipment](replenishment-shipment-entity.md)

INB-SHIPMENT

Parent

Replace

EQUI

Parent

Replace

LIKP

Parent

Replace

LIPS

Parent

Replace

0MATERIAL\_TEXT

Parent

Replace

0MAT\_VEND\_ATTR

Parent

Replace

0MATERIAL\_ATTR

Parent

Replace

EKPO

Parent

Replace

T001W

Parent

Replace

ADRC

Parent

Replace

0VENDOR\_ATTR

Parent

Replace

BUT021\_FS

Parent

Replace

Site – [site](network-site-entity.md)

T001W

Parent

Replace

ADRC

Child

Update

GEOLOC

Child

Update

Trading partner – [trading\_partner](organization-trading-partner-entity.md)

0BPARTNER\_ATTR

Parent

Update

0BPARTNER\_TEXT

Child

Update

0VENDOR\_ATTR

Child

Update

0CUSTOMER\_ATTR

Child

Update

0BP\_DEF\_ADDRESS\_ATTR

Child

Update

Transfer order – [inbound\_order\_line](replenishment-inbound-order-line-entity.md)

2LIS\_03\_BF

Parent

Update

0MATERIAL\_ATTR

Child

Update

Transportation – [transportation\_lane](network-transporation-lane-entity.md)

TVRO

Parent

Replace

TVRAB

Child

Update

VALW

Child

Update

Vendor management – [vendor\_lead\_time](vendor-management-lead-time-entity.md)

EINA

Parent

Replace

EINE

Child

Update

0MATERIAL\_ATTR

Child

Update

Vendor management – [vendor\_product](vendor-management-product-entity.md)

EINA

Parent

Replace

0MATERIAL\_ATTR

Child

Update

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Connecting to S/4 HANA

Adding a new outbound source for Supply Planning

All content copied from https://docs.aws.amazon.com/.
