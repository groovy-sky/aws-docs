This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MWAAServerless::Workflow S3Location

Specifies the Amazon S3 location of a workflow definition file. This structure contains the bucket name, object key, and optional version ID for the workflow definition. Amazon Managed Workflows for Apache Airflow Serverless takes a snapshot of the definition file at the time of workflow creation or update, ensuring that the workflow behavior remains consistent even if the source file is modified. The definition must be a valid YAML file that uses supported AWS operators and Amazon Managed Workflows for Apache Airflow Serverless syntax.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "ObjectKey" : String,
  "VersionId" : String
}

```

### YAML

```yaml

  Bucket: String
  ObjectKey: String
  VersionId: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket that contains the workflow definition file.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ObjectKey`

The key (name) of the workflow definition file within the S3 bucket.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionId`

Optional. The version ID of the workflow definition file in Amazon S3. If not specified, the latest version is used.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NetworkConfiguration

ScheduleConfiguration

All content copied from https://docs.aws.amazon.com/.
