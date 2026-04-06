# MetadataCatalogDetail

Describes the metadata catalog, metadata table, and data partitions that Amazon AppFlow used for the associated flow run.

## Contents

**catalogType**

The type of metadata catalog that Amazon AppFlow used for the associated flow run.
This parameter returns the following value:

GLUE

The metadata catalog is provided by the AWS Glue Data Catalog. AWS Glue
includes the AWS Glue Data Catalog as a component.

Type: String

Valid Values: `GLUE`

Required: No

**partitionRegistrationOutput**

Describes the status of the attempt from Amazon AppFlow to register the data
partitions with the metadata catalog. The data partitions organize the flow output into a
hierarchical path, such as a folder path in an S3 bucket. Amazon AppFlow creates the
partitions (if they don't already exist) based on your flow configuration.

Type: [RegistrationOutput](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_RegistrationOutput.html) object

Required: No

**tableName**

The name of the table that stores the metadata for the associated flow run. The table
stores metadata that represents the data that the flow transferred. Amazon AppFlow stores
the table in the metadata catalog.

Type: String

Length Constraints: Maximum length of 2048.

Pattern: `.*`

Required: No

**tableRegistrationOutput**

Describes the status of the attempt from Amazon AppFlow to register the metadata
table with the metadata catalog. Amazon AppFlow creates or updates this table for the
associated flow run.

Type: [RegistrationOutput](https://docs.aws.amazon.com/appflow/1.0/APIReference/API_RegistrationOutput.html) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/MetadataCatalogDetail)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/MetadataCatalogDetail)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/MetadataCatalogDetail)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetadataCatalogConfig

OAuth2Credentials
