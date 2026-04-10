This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Job JobCommand

Specifies code executed when a job is run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "PythonVersion" : String,
  "Runtime" : String,
  "ScriptLocation" : String
}

```

### YAML

```yaml

  Name: String
  PythonVersion: String
  Runtime: String
  ScriptLocation: String

```

## Properties

`Name`

The name of the job command. For an Apache Spark ETL job, this must be
`glueetl`. For a Python shell job, it must be `pythonshell`.
For an Apache Spark streaming ETL job, this must be `gluestreaming`. For a Ray job,
this must be `glueray`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PythonVersion`

The Python version being used to execute a Python shell job. Allowed values are 3 or 3.9. Version 2 is deprecated.

_Required_: No

_Type_: String

_Pattern_: `^([2-3]|3[.]9)$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Runtime`

In Ray jobs, Runtime is used to specify the versions of Ray, Python and additional libraries available in your environment. This field is not used in other job types. For supported runtime environment values, see [Working with Ray jobs](../../../glue/latest/dg/ray-jobs-section.md) in the AWS Glue Developer Guide.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScriptLocation`

Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that executes
a job (required).

_Required_: No

_Type_: String

_Maximum_: `400000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExecutionProperty

NotificationProperty

All content copied from https://docs.aws.amazon.com/.
