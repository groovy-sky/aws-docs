# Connecting to an EDI

To ingest data from an EDI data source, follow the procedure below.

01. On the AWS Supply Chain dashboard, on the left navigation pane, choose
     **Data Lake**.

02. On the **Data lake** page, choose **Add New**
    **Source**.

    The **Select your supply chain data source** page
     appears.

03. Choose **EDI**.

04. In the **EDI Connection Details** page, under
     **Name your connection**, enter a name for your
     connection.

05. (Optional) Under **Connection description**, enter a
     description for your connection.

06. Under **Amazon S3 Bucket Billing**, review the Amazon S3 billing
     information, and then select **Acknowledge**.

07. Choose **Next**.

08. Under **Data Mapping**, choose **Get**
    **started**.

09. ###### Note

    EDI 850, EDI 860, and EDI 856 are supported in
    AWS Supply Chain.

    ###### Note

    The required fields are already mapped. Perform this step only if you
    want to make specific changes to the default transformation
    recipe.

    On the **Mapping Recipe** page, you can view the default
     transformation recipe under **Field mappings**.

    Choose **Add mapping**, to map any additional destination
     field. The **Required Destination Fields** are mandatory.
     Choose **Destination field** to add an additional custom
     destination field.

    ###### Note

    Review all the entities (for example, Inbound Order, Inbound Order
    Line, and Inbound Order Line Schedule for EDI 850 Entity Group) under
    each Entity Group.

10. To view the source field values and data mappings from the transformation
     recipe, you can upload sample data. On the **Mapping**
    **Recipe** page, under **Upload sample data**,
     choose **browse files**, or drag and drop files. The sample
     data file must contain the required parameters and include the source field
     names.

11. Choose **Accept all and continue**.

12. Under **Review and confirm**, you can view the data
     connection summary. To edit your data field mapping, choose **Go**
    **back to Data Mapping**.

13. Choose **Confirm and configure data ingestion** to review
     the Amazon S3 paths where your source data must be uploaded to start the
     ingestion process.

14. Choose **Confirm and configure data ingestion later** if
     you want to ingest data later. You can ingest data anytime after creating
     the connection from the AWS Supply Chain dashboard.

15. On the AWS Supply Chain dashboard, choose **Open**
    **Connections**. Select the connection dataflow that you want to
     ingest data, choose the vertical ellipsis, and select **Ingestion**
    **setup**.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Uploading subsequent files to an existing source

Connecting to S/4 HANA

All content copied from https://docs.aws.amazon.com/.
