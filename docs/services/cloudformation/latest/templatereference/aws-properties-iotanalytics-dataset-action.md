This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset Action

Information needed to run the "containerAction" to produce data set contents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActionName" : String,
  "ContainerAction" : ContainerAction,
  "QueryAction" : QueryAction
}

```

### YAML

```yaml

  ActionName: String
  ContainerAction:
    ContainerAction
  QueryAction:
    QueryAction

```

## Properties

`ActionName`

The name of the data set action by which data set contents are automatically created.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContainerAction`

Information which allows the system to run a containerized application in order to create
the data set contents. The application must be in a Docker container along with any needed
support libraries.

_Required_: No

_Type_: [ContainerAction](aws-properties-iotanalytics-dataset-containeraction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryAction`

An "SqlQueryDatasetAction" object that uses an SQL query to automatically create data set contents.

_Required_: No

_Type_: [QueryAction](aws-properties-iotanalytics-dataset-queryaction.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTAnalytics::Dataset

ContainerAction

All content copied from https://docs.aws.amazon.com/.
