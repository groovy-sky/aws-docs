This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::MigrationProject

Provides information that defines a migration project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::MigrationProject",
  "Properties" : {
      "Description" : String,
      "InstanceProfileArn" : String,
      "InstanceProfileIdentifier" : String,
      "InstanceProfileName" : String,
      "MigrationProjectIdentifier" : String,
      "MigrationProjectName" : String,
      "SchemaConversionApplicationAttributes" : SchemaConversionApplicationAttributes,
      "SourceDataProviderDescriptors" : [ DataProviderDescriptor, ... ],
      "Tags" : [ Tag, ... ],
      "TargetDataProviderDescriptors" : [ DataProviderDescriptor, ... ],
      "TransformationRules" : String
    }
}

```

### YAML

```yaml

Type: AWS::DMS::MigrationProject
Properties:
  Description: String
  InstanceProfileArn: String
  InstanceProfileIdentifier: String
  InstanceProfileName: String
  MigrationProjectIdentifier: String
  MigrationProjectName: String
  SchemaConversionApplicationAttributes:
    SchemaConversionApplicationAttributes
  SourceDataProviderDescriptors:
    - DataProviderDescriptor
  Tags:
    - Tag
  TargetDataProviderDescriptors:
    - DataProviderDescriptor
  TransformationRules: String

```

## Properties

`Description`

A user-friendly description of the migration project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileArn`

The Amazon Resource Name (ARN) of the instance profile for your migration project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileIdentifier`

The identifier of the instance profile for your migration project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceProfileName`

The name of the associated instance profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MigrationProjectIdentifier`

The identifier of the migration project. Identifiers must begin with a letter
and must contain only ASCII letters, digits, and hyphens. They can't end with
a hyphen, or contain two consecutive hyphens.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MigrationProjectName`

The name of the migration project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SchemaConversionApplicationAttributes`

The schema conversion application attributes, including the Amazon S3 bucket name and Amazon S3 role ARN.

_Required_: No

_Type_: [SchemaConversionApplicationAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-migrationproject-schemaconversionapplicationattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceDataProviderDescriptors`

Information about the source data provider, including the name or ARN, and AWS Secrets Manager parameters.

_Required_: No

_Type_: Array of [DataProviderDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-migrationproject-dataproviderdescriptor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-migrationproject-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetDataProviderDescriptors`

Information about the target data provider, including the name or ARN, and AWS Secrets Manager parameters.

_Required_: No

_Type_: Array of [DataProviderDescriptor](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-migrationproject-dataproviderdescriptor.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransformationRules`

The settings in JSON format for migration rules. Migration rules make it possible for you to change
the object names according to the rules that you specify. For example, you can change an object name
to lowercase or uppercase, add or remove a prefix or suffix, or rename objects.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MigrationProjectArn`

The ARN string that uniquely identifies the migration project.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DataProviderDescriptor
