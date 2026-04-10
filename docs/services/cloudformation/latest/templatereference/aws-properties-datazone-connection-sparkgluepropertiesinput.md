This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Connection SparkGluePropertiesInput

The Spark AWS Glue properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalArgs" : SparkGlueArgs,
  "GlueConnectionName" : String,
  "GlueVersion" : String,
  "IdleTimeout" : Number,
  "JavaVirtualEnv" : String,
  "NumberOfWorkers" : Number,
  "PythonVirtualEnv" : String,
  "WorkerType" : String
}

```

### YAML

```yaml

  AdditionalArgs:
    SparkGlueArgs
  GlueConnectionName: String
  GlueVersion: String
  IdleTimeout: Number
  JavaVirtualEnv: String
  NumberOfWorkers:
    Number
  PythonVirtualEnv: String
  WorkerType: String

```

## Properties

`AdditionalArgs`

The additional args in the Spark AWS Glue properties.

_Required_: No

_Type_: [SparkGlueArgs](aws-properties-datazone-connection-sparkglueargs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueConnectionName`

The AWS Glue connection name in the Spark AWS Glue
properties. Specify either `glueConnectionName` or
`glueConnectionNames`, but not both.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlueVersion`

The AWS Glue version in the Spark AWS Glue
properties.

_Required_: No

_Type_: String

_Pattern_: `^\w+\.\w+$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IdleTimeout`

The idle timeout in the Spark AWS Glue properties.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `3000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JavaVirtualEnv`

The Java virtual env in the Spark AWS Glue properties.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumberOfWorkers`

The number of workers in the Spark AWS Glue properties.

_Required_: No

_Type_: Number

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PythonVirtualEnv`

The Python virtual env in the Spark AWS Glue properties.

_Required_: No

_Type_: String

_Pattern_: `^[\S]*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WorkerType`

The worker type in the Spark AWS Glue properties.

_Required_: No

_Type_: String

_Pattern_: `^[G|Z].*$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SparkGlueArgs

UsernamePassword

All content copied from https://docs.aws.amazon.com/.
