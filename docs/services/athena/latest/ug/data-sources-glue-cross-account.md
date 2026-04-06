# Register a Data Catalog from another account

You can use Athena's cross-account AWS Glue catalog feature to register an AWS Glue catalog from
an account other than your own. After you configure the required IAM permissions for AWS Glue
and register the catalog as an Athena `DataCatalog` resource, you can use Athena to
run cross-account queries. For information about configuring the required permissions, see
[Configure cross-account access to AWS Glue data catalogs](https://docs.aws.amazon.com/athena/latest/ug/security-iam-cross-account-glue-catalog-access.html).

The following procedure shows you how to use the Athena to configure an AWS Glue Data Catalog in an
Amazon Web Services account other than your own as a data source.

## Register from console

01. Follow the steps in [Configure cross-account access to AWS Glue data catalogs](https://docs.aws.amazon.com/athena/latest/ug/security-iam-cross-account-glue-catalog-access.html) to ensure that
     you have permissions to query the data catalog in the other account.

02. Open the Athena console at
     [https://console.aws.amazon.com/athena/](https://console.aws.amazon.com/athena/home).

03. If the console navigation pane is not visible, choose the expansion menu
     on the left.

    ![Choose the expansion menu.](https://docs.aws.amazon.com/images/athena/latest/ug/images/nav-pane-expansion.png)

04. Choose **Data sources and catalogs**.

05. On the upper right, choose **Create data source**.

06. On the **Choose a data source** page, for **Data**
    **sources**, choose **S3 - AWS Glue Data Catalog**, and then
     choose **Next**.

07. On the **Enter data source details** page, in the
     **AWS Glue Data Catalog** section, for **Choose an**
    **AWS Glue Data Catalog**, choose **AWS Glue Data Catalog in another**
    **account**.

08. For **Data source details**, enter the following
     information:

- **Data source name** – Enter the name that you
want to use in your SQL queries to refer to the data catalog in the other
account.

- **Description** – (Optional) Enter a description
of the data catalog in the other account.

- **Catalog ID** – Enter the 12-digit Amazon Web Services
account ID of the account to which the data catalog belongs. The Amazon Web Services
account ID is the catalog ID.

09. (Optional) For **Tags**, enter key-value pairs that you want to
     associate with the data source. For more information about tags, see [Tag Athena resources](tags.md).

10. Choose **Next**.

11. On the **Review and create** page, review the information that
     you provided, and then choose **Create data source**. The
     **Data source details** page lists the databases and tags for
     the data catalog that you registered.

12. Choose **Data sources and catalogs**. The data catalog that you registered is
     listed in the **Data source name** column.

13. To view or edit information about the data catalog, choose the catalog, and then
     choose **Actions**, **Edit**.

14. To delete the new data catalog, choose the catalog, and then choose
     **Actions**, **Delete**.

## Register using API operations

1. The following `CreateDataCatalog` request body registers an AWS Glue catalog
    for cross-account access:

```json

# Example CreateDataCatalog request to register a cross-account Glue catalog:
{
       "Description": "Cross-account Glue catalog",
       "Name": "ownerCatalog",
       "Parameters": {"catalog-id" : "<catalogid>"  # Owner's account ID
       },
       "Type": "GLUE"
}
```

2. The following sample code uses a Java client to create the `DataCatalog`
    object.

```java

# Sample code to create the DataCatalog through Java client
CreateDataCatalogRequest request = new CreateDataCatalogRequest()
       .withName("ownerCatalog")
       .withType(DataCatalogType.GLUE)
       .withParameters(ImmutableMap.of("catalog-id", "<catalogid>"));

athenaClient.createDataCatalog(request);
```

After these steps, the borrower should see
    `ownerCatalog` when it calls the [ListDataCatalogs](https://docs.aws.amazon.com/athena/latest/APIReference/API_ListDataCatalogs.html) API operation.

## Register using AWS CLI

Use the followig example CLI command to register an AWS Glue Data Catalog from another account

```nohighlight

aws athena create-data-catalog \
  --name cross_account_catalog \
  --type GLUE \
  --description "Cross Account Catalog" \
  --parameters catalog-id=<catalogid>
```

For more information, see [Query\
cross-account AWS Glue Data Catalogs using Amazon Athena](https://aws.amazon.com/blogs/big-data/query-cross-account-aws-glue-data-catalogs-using-amazon-athena) in the _AWS Big Data_
_Blog_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Query AWS Glue data catalogs in Athena

Control access to data catalogs
