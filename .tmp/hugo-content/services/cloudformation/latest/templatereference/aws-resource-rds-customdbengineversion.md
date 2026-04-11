This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::CustomDBEngineVersion

Creates a custom DB engine version (CEV).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::CustomDBEngineVersion",
  "Properties" : {
      "DatabaseInstallationFilesS3BucketName" : String,
      "DatabaseInstallationFilesS3Prefix" : String,
      "Description" : String,
      "Engine" : String,
      "EngineVersion" : String,
      "ImageId" : String,
      "KMSKeyId" : String,
      "Manifest" : String,
      "SourceCustomDbEngineVersionIdentifier" : String,
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "UseAwsProvidedLatestImage" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::RDS::CustomDBEngineVersion
Properties:
  DatabaseInstallationFilesS3BucketName: String
  DatabaseInstallationFilesS3Prefix: String
  Description: String
  Engine: String
  EngineVersion: String
  ImageId: String
  KMSKeyId: String
  Manifest: String
  SourceCustomDbEngineVersionIdentifier: String
  Status: String
  Tags:
    - Tag
  UseAwsProvidedLatestImage: Boolean

```

## Properties

`DatabaseInstallationFilesS3BucketName`

The name of an Amazon S3 bucket that contains database installation files for your CEV. For example, a valid
bucket name is `my-custom-installation-files`.

_Required_: No

_Type_: String

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DatabaseInstallationFilesS3Prefix`

The Amazon S3 directory that contains the database installation files for your CEV. For example, a valid
bucket name is `123456789012/cev1`. If this setting isn't specified, no prefix is assumed.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

An optional description of your CEV.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The database engine to use for your custom engine version (CEV).

Valid values:

- `custom-oracle-ee`

- `custom-oracle-ee-cdb`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `35`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EngineVersion`

The name of your CEV. The name format is `major version.customized_string`.
For example, a valid CEV name is `19.my_cev1`. This setting is required for RDS
Custom for Oracle, but optional for Amazon RDS. The combination of `Engine`
and `EngineVersion` is unique per customer per Region.

_Constraints:_ Minimum length is 1. Maximum length is 60.

_Pattern:_ `^[a-z0-9_.-]{1,60$`}

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `60`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ImageId`

A value that indicates the ID of the AMI.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KMSKeyId`

The AWS KMS key identifier for an encrypted CEV. A symmetric encryption KMS key is required for
RDS Custom, but optional for Amazon RDS.

If you have an existing symmetric encryption KMS key in your account, you can use it with RDS Custom.
No further action is necessary. If you don't already have a symmetric encryption KMS key in your account,
follow the instructions in [Creating a symmetric encryption KMS key](../../../kms/latest/developerguide/create-keys.md#create-symmetric-cmk) in the _AWS Key Management Service_
_Developer Guide_.

You can choose the same symmetric encryption key when you create a CEV and a DB instance, or choose different keys.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Manifest`

The CEV manifest, which is a JSON document that describes the installation .zip files stored in Amazon S3.
Specify the name/value pairs in a file or a quoted string. RDS Custom applies the patches in the order in which
they are listed.

The following JSON fields are valid:

MediaImportTemplateVersion

Version of the CEV manifest. The date is in the format `YYYY-MM-DD`.

databaseInstallationFileNames

Ordered list of installation files for the CEV.

opatchFileNames

Ordered list of OPatch installers used for the Oracle DB engine.

psuRuPatchFileNames

The PSU and RU patches for this CEV.

OtherPatchFileNames

The patches that are not in the list of PSU and RU patches.
Amazon RDS applies these patches after applying the PSU and RU patches.

For more information, see [Creating the CEV manifest](../../../amazonrds/latest/userguide/custom-cev.md#custom-cev.preparing.manifest) in the _Amazon RDS User Guide_.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `51000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SourceCustomDbEngineVersionIdentifier`

The ARN of a CEV to use as a source for creating a new CEV. You can specify a different
Amazon Machine Imagine (AMI) by using either `Source` or
`UseAwsProvidedLatestImage`. You can't specify a different JSON manifest
when you specify `SourceCustomDbEngineVersionIdentifier`.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Status`

A value that indicates the status of a custom engine version (CEV).

_Required_: No

_Type_: String

_Allowed values_: `available | inactive | inactive-except-restore`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags.
For more information, see [Tagging \
Amazon RDS Resources](../../../amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide._

_Required_: No

_Type_: Array of [Tag](aws-properties-rds-customdbengineversion-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseAwsProvidedLatestImage`

Specifies whether to use the latest service-provided Amazon Machine Image (AMI) for
the CEV. If you specify `UseAwsProvidedLatestImage`, you can't also specify
`ImageId`.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref
returns the resource name, such as `mystack-abc1d2efg3h4.` For more
information about using the Ref function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

GetAtt returns a value for a specified attribute of this type. For more information, see
[Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md). The following attributes are available.

`DBEngineVersionArn`

The ARN of the custom engine version.

## Examples

The following example specifies a custom DB engine version.

### Authentication details

#### JSON

```json

"RDSCustomDBEngineVersion":{
   "Type":"AWS::RDS::CustomDBEngineVersion",
   "Properties":{
      "DatabaseInstallationFilesS3BucketName": "1-custom-installation-files",
      "DatabaseInstallationFilesS3Prefix": "123456789012/cev1",
      "Description": "CEV description",
      "Engine": "custom-oracle-ee",
      "EngineVersion": "19.cev1",
      "KMSKeyId": "12ab3c4d-5678-90e1-2fg3-45h6ijklmnops",
      "Manifest": "%7B%22mediaImportTemplateVersion%22%3A%222020-08--14%22%2C%22databaseInstallationFileNames%22%3A%5B%22V982063-01.zip%22%5D%2C%22opatchFileNames%22%3A%5B%22p6880880_190000_Linux-x86-64.zip%22%5D%2C%22psuRuPatchFileNames%22%3A%5B%22p31720396_190000_Linux-x86-64.zip%22%2C%22p29213893_199000DBRU_Generic.zip%22%2C%22p28730253_190000_Linux-x86-64.zip%22%2C%22p29374604_199000DBRU_Linux-x86-64.zip%22%2C%22p28852325_190000_Linux-x86-64.zip%22%2C%22p29997937_190000_Linux-x86-64.zip%22%2C%22p31335037_190000_Linux-x86-64.zip%22%2C%22p31335142_190000_Generic.zip%22%5D%7D"
  }
}
```

#### YAML

```yaml

CustomDBEngineVersion:
  Type: AWS::RDS::CustomDBEngineVersion
  Properties:
      DatabaseInstallationFilesS3BucketName: 1-custom-installation-files
      DatabaseInstallationFilesS3Prefix: "123456789012/cev1
      Description: CEV description
      Engine: custom-oracle-ee
      EngineVersion: 19.cev1
      KMSKeyId: 12ab3c4d-5678-90e1-2fg3-45h6ijklmnops
      Manifest: "%7B%22mediaImportTemplateVersion%22%3A%222020-08--14%22%2C%22databaseInstallationFileNames%22%3A%5B%22V982063-01.zip%22%5D%2C%22opatchFileNames%22%3A%5B%22p6880880_190000_Linux-x86-64.zip%22%5D%2C%22psuRuPatchFileNames%22%3A%5B%22p31720396_190000_Linux-x86-64.zip%22%2C%22p29213893_199000DBRU_Generic.zip%22%2C%22p28730253_190000_Linux-x86-64.zip%22%2C%22p29374604_199000DBRU_Linux-x86-64.zip%22%2C%22p28852325_190000_Linux-x86-64.zip%22%2C%22p29997937_190000_Linux-x86-64.zip%22%2C%22p31335037_190000_Linux-x86-64.zip%22%2C%22p31335142_190000_Generic.zip%22%5D%7D
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Relational Database Service

Tag

All content copied from https://docs.aws.amazon.com/.
