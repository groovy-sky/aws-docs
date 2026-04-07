This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::DataCatalogEncryptionSettings

Sets the security configuration for a specified catalog. After the configuration has been
set, the specified encryption is applied to every catalog write thereafter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::DataCatalogEncryptionSettings",
  "Properties" : {
      "CatalogId" : String,
      "DataCatalogEncryptionSettings" : DataCatalogEncryptionSettings
    }
}

```

### YAML

```yaml

Type: AWS::Glue::DataCatalogEncryptionSettings
Properties:
  CatalogId: String
  DataCatalogEncryptionSettings:
    DataCatalogEncryptionSettings

```

## Properties

`CatalogId`

The ID of the Data Catalog in which the settings are created.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataCatalogEncryptionSettings`

Contains configuration information for maintaining Data Catalog security.

_Required_: Yes

_Type_: [DataCatalogEncryptionSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PrincipalPrivileges

ConnectionPasswordEncryption
