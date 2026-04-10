This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Step HadoopJarStepConfig

A job flow step consisting of a JAR file whose main function will be executed. The main
function submits a job for Hadoop to execute and waits for the job to finish or
fail.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Args" : [ String, ... ],
  "Jar" : String,
  "MainClass" : String,
  "StepProperties" : [ KeyValue, ... ]
}

```

### YAML

```yaml

  Args:
    - String
  Jar: String
  MainClass: String
  StepProperties:
    - KeyValue

```

## Properties

`Args`

A list of command line arguments passed to the JAR file's main function when
executed.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Jar`

A path to a JAR file run during the step.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MainClass`

The name of the main class in the specified Java file. If not specified, the JAR file
should specify a Main-Class in its manifest file.

_Required_: No

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\r\n\t]*`

_Minimum_: `0`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StepProperties`

A list of Java properties that are set when the step runs. You can use these properties to pass key value pairs to your main function.

_Required_: No

_Type_: Array of [KeyValue](aws-properties-emr-step-keyvalue.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::Step

KeyValue

All content copied from https://docs.aws.amazon.com/.
