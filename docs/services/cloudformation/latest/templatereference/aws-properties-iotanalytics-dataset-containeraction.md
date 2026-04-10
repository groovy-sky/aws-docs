This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset ContainerAction

Information needed to run the "containerAction" to produce data set contents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionRoleArn" : String,
  "Image" : String,
  "ResourceConfiguration" : ResourceConfiguration,
  "Variables" : [ Variable, ... ]
}

```

### YAML

```yaml

  ExecutionRoleArn: String
  Image: String
  ResourceConfiguration:
    ResourceConfiguration
  Variables:
    - Variable

```

## Properties

`ExecutionRoleArn`

The ARN of the role which gives permission to the system to access needed resources in order
to run the "containerAction". This includes, at minimum, permission to retrieve the data set
contents which are the input to the containerized application.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Image`

The ARN of the Docker container stored in your account. The Docker container contains an
application and needed support libraries and is used to generate data set contents.

_Required_: Yes

_Type_: String

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceConfiguration`

Configuration of the resource which executes the "containerAction".

_Required_: Yes

_Type_: [ResourceConfiguration](aws-properties-iotanalytics-dataset-resourceconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The values of variables used within the context of the execution of the containerized
application (basically, parameters passed to the application). Each variable must have a
name and a value given by one of "stringValue", "datasetContentVersionValue",
or "outputFileUriValue".

_Required_: No

_Type_: Array of [Variable](aws-properties-iotanalytics-dataset-variable.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

DatasetContentDeliveryRule

All content copied from https://docs.aws.amazon.com/.
