This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataProvider

Provides information that defines a data provider.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::DataProvider",
  "Properties" : {
      "DataProviderIdentifier" : String,
      "DataProviderName" : String,
      "Description" : String,
      "Engine" : String,
      "ExactSettings" : Boolean,
      "Settings" : Settings,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DMS::DataProvider
Properties:
  DataProviderIdentifier: String
  DataProviderName: String
  Description: String
  Engine: String
  ExactSettings: Boolean
  Settings:
    Settings
  Tags:
    - Tag

```

## Properties

`DataProviderIdentifier`

The identifier of the data provider. Identifiers must begin with a letter
and must contain only ASCII letters, digits, and hyphens. They can't end with
a hyphen, or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataProviderName`

The name of the data provider.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the data provider. Descriptions can have up to 31 characters.
A description can contain only ASCII letters, digits, and hyphens ('-'). Also, it can't
end with a hyphen or contain two consecutive hyphens, and can only begin with a letter.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The type of database engine for the data provider. Valid values include `"aurora"`,
`"aurora-postgresql"`, `"mysql"`, `"oracle"`, `"postgres"`,
`"sqlserver"`, `redshift`, `mariadb`, `mongodb`, `db2`, `db2-zos`, `docdb`, and `sybase`. A value of `"aurora"` represents Amazon Aurora MySQL-Compatible Edition.

_Required_: Yes

_Type_: String

_Allowed values_: `aurora | aurora_postgresql | mysql | oracle | postgres | sqlserver | redshift | mariadb | mongodb | docdb | db2 | db2_zos | sybase`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExactSettings`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

The settings in JSON format for a data provider.

_Required_: No

_Type_: [Settings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-dataprovider-settings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-dataprovider-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DataProviderArn`

The Amazon Resource Name (ARN) string that uniquely identifies the data provider.

`DataProviderCreationTime`

The time the data provider was created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DocDbSettings
