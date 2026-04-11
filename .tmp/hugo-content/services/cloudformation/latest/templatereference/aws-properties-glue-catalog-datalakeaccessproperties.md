This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Catalog DataLakeAccessProperties

Input properties to configure data lake access for your catalog resource in the AWS Glue Data Catalog.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllowFullTableExternalDataAccess" : String,
  "CatalogType" : String,
  "DataLakeAccess" : Boolean,
  "DataTransferRole" : String,
  "KmsKey" : String,
  "ManagedWorkgroupName" : String,
  "ManagedWorkgroupStatus" : String,
  "RedshiftDatabaseName" : String
}

```

### YAML

```yaml

  AllowFullTableExternalDataAccess: String
  CatalogType: String
  DataLakeAccess: Boolean
  DataTransferRole: String
  KmsKey: String
  ManagedWorkgroupName: String
  ManagedWorkgroupStatus: String
  RedshiftDatabaseName: String

```

## Properties

`AllowFullTableExternalDataAccess`

Allows third-party engines to access data in Amazon S3 locations that are registered
with AWS Lake Formation.

_Required_: No

_Type_: String

_Allowed values_: `True | False`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CatalogType`

Specifies a federated catalog type for the native catalog resource. The currently supported type is `aws:redshift`.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataLakeAccess`

Turns on or off data lake access for Apache Spark applications that access Amazon Redshift
databases in the Data Catalog from any non-Redshift engine, such as Amazon Athena, Amazon EMR,
or AWS Glue ETL.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataTransferRole`

A role that will be assumed by AWS Glue for transferring data into/out of the staging bucket during a query.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(-(cn|us-gov|iso(-[bef])?))?:iam::[0-9]{12}:role/.+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKey`

An encryption key that will be used for the staging bucket that will be created along with the catalog.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedWorkgroupName`

The name of the managed workgroup associated with the catalog.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManagedWorkgroupStatus`

The status of the managed workgroup.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RedshiftDatabaseName`

The name of the Redshift database.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CatalogProperties

DataLakePrincipal

All content copied from https://docs.aws.amazon.com/.
