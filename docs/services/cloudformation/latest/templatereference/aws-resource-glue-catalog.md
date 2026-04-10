This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Catalog

The `AWS::Glue::Catalog` resource specifies a catalog object that
represents a logical grouping of databases in the AWS Glue Data Catalog
or a federated source. You can create a Redshift-federated catalog or a catalog containing
resource links to Redshift databases in another account or region.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::Catalog",
  "Properties" : {
      "AllowFullTableExternalDataAccess" : String,
      "CatalogProperties" : CatalogProperties,
      "CreateDatabaseDefaultPermissions" : [ PrincipalPermissions, ... ],
      "CreateTableDefaultPermissions" : [ PrincipalPermissions, ... ],
      "Description" : String,
      "FederatedCatalog" : FederatedCatalog,
      "Name" : String,
      "OverwriteChildResourcePermissionsWithDefault" : String,
      "Parameters" : {Key: Value, ...},
      "Tags" : [ Tag, ... ],
      "TargetRedshiftCatalog" : TargetRedshiftCatalog
    }
}

```

### YAML

```yaml

Type: AWS::Glue::Catalog
Properties:
  AllowFullTableExternalDataAccess: String
  CatalogProperties:
    CatalogProperties
  CreateDatabaseDefaultPermissions:
    - PrincipalPermissions
  CreateTableDefaultPermissions:
    - PrincipalPermissions
  Description: String
  FederatedCatalog:
    FederatedCatalog
  Name: String
  OverwriteChildResourcePermissionsWithDefault: String
  Parameters:
    Key: Value
  Tags:
    - Tag
  TargetRedshiftCatalog:
    TargetRedshiftCatalog

```

## Properties

`AllowFullTableExternalDataAccess`

Allows third-party engines to access data in Amazon S3 locations that are registered
with AWS Lake Formation.

###### Note

This property is write-only. The value is not returned by `Fn::GetAtt`
or `Ref`.

_Required_: No

_Type_: String

_Allowed values_: `True | False`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CatalogProperties`

A `CatalogProperties` object that specifies data lake access properties and other custom properties.

_Required_: No

_Type_: [CatalogProperties](aws-properties-glue-catalog-catalogproperties.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateDatabaseDefaultPermissions`

An array of `PrincipalPermissions` objects. Creates a set of default permissions on the database(s) for principals. Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations.

_Required_: No

_Type_: Array of [PrincipalPermissions](aws-properties-glue-catalog-principalpermissions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CreateTableDefaultPermissions`

An array of `PrincipalPermissions` objects. Creates a set of default permissions on the table(s) for principals. Used by AWS Lake Formation. Not used in the normal course of AWS Glue operations.

_Required_: No

_Type_: Array of [PrincipalPermissions](aws-properties-glue-catalog-principalpermissions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the catalog, not more than 2048 bytes long.

_Required_: No

_Type_: String

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederatedCatalog`

A `FederatedCatalog` object that points to an entity outside the AWS Glue Data Catalog.

_Required_: No

_Type_: [FederatedCatalog](aws-properties-glue-catalog-federatedcatalog.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the catalog. Cannot be the same as the account ID.

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OverwriteChildResourcePermissionsWithDefault`

Specifies whether to overwrite child resource permissions with the default permissions.
Valid values are `Accept` and `Deny`.

###### Note

This property is write-only. The value is not returned by `Fn::GetAtt`
or `Ref`.

_Required_: No

_Type_: String

_Allowed values_: `Accept | Deny`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Parameters`

A map array of key-value pairs that define parameters and properties of the catalog.

_Required_: No

_Type_: Object of String

_Pattern_: `^.+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-glue-catalog-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRedshiftCatalog`

A `TargetRedshiftCatalog` object that describes a target catalog for database resource linking.

_Required_: No

_Type_: [TargetRedshiftCatalog](aws-properties-glue-catalog-targetredshiftcatalog.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the catalog, such as
`arn:aws:glue:us-east-1:123456789012:catalog/my-catalog`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CatalogId`

The ID of the catalog.

`CatalogProperties.DataLakeAccessProperties.ManagedWorkgroupName`

The name of the managed workgroup associated with the catalog.

`CatalogProperties.DataLakeAccessProperties.ManagedWorkgroupStatus`

The status of the managed workgroup.

`CatalogProperties.DataLakeAccessProperties.RedshiftDatabaseName`

The name of the Redshift database associated with the catalog.

`CreateTime`

The time at which the catalog was created.

`ResourceArn`

The Amazon Resource Name (ARN) assigned to the catalog resource.

`UpdateTime`

The time at which the catalog was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Glue

CatalogProperties

All content copied from https://docs.aws.amazon.com/.
