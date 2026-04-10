This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::TableOptimizer VpcConfiguration

An object that describes the VPC configuration for a table optimizer. This configuration is necessary to perform optimization on tables that are in a customer VPC.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueConnectionName" : String
}

```

### YAML

```yaml

  GlueConnectionName: String

```

## Properties

`GlueConnectionName`

The name of the AWS Glue connection used for the VPC for the table optimizer.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TableOptimizerConfiguration

AWS::Glue::Trigger

All content copied from https://docs.aws.amazon.com/.
