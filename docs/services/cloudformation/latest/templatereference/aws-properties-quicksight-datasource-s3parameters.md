This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::DataSource S3Parameters

The parameters for S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ManifestFileLocation" : ManifestFileLocation,
  "RoleArn" : String
}

```

### YAML

```yaml

  ManifestFileLocation:
    ManifestFileLocation
  RoleArn: String

```

## Properties

`ManifestFileLocation`

Location of the Amazon S3 manifest file. This is NULL if the manifest file was
uploaded into Quick Sight.

_Required_: Yes

_Type_: [ManifestFileLocation](aws-properties-quicksight-datasource-manifestfilelocation.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

Use the `RoleArn` structure to override an account-wide role for a specific S3 data source. For example, say an account administrator has turned off all S3 access with an account-wide role. The administrator can then use `RoleArn` to bypass the account-wide role and allow S3 access for the single S3 data source that is specified in the structure, even if the account-wide role forbidding S3 access is still active.

_Required_: No

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourcePermission

S3TablesParameters

All content copied from https://docs.aws.amazon.com/.
