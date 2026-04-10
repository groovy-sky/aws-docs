This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup S3Location

`S3Location` is a property of the [CodeDeploy DeploymentGroup Revision](../userguide/aws-properties-codedeploy-deploymentgroup-deployment-revision.md) property that specifies the location
of an application revision that is stored in Amazon Simple Storage Service (Amazon S3).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "BundleType" : String,
  "ETag" : String,
  "Key" : String,
  "Version" : String
}

```

### YAML

```yaml

  Bucket: String
  BundleType: String
  ETag: String
  Key: String
  Version: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket where the application revision is
stored.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BundleType`

The file type of the application revision. Must be one of the following:

- JSON

- tar: A tar archive file.

- tgz: A compressed tar archive file.

- YAML

- zip: A zip archive file.

_Required_: No

_Type_: String

_Allowed values_: `tar | tgz | zip | YAML | JSON`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ETag`

The ETag of the Amazon S3 object that represents the bundled artifacts for the
application revision.

If the ETag is not specified as an input parameter, ETag validation of the object is
skipped.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The name of the Amazon S3 object that represents the bundled artifacts for the
application revision.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

A specific version of the Amazon S3 object that represents the bundled
artifacts for the application revision.

If the version is not specified, the system uses the most recent version by
default.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RevisionLocation

Tag

All content copied from https://docs.aws.amazon.com/.
