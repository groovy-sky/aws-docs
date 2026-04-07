This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::AnalysisTemplate ErrorMessageConfiguration

A structure that defines the level of detail included in error messages returned by PySpark jobs. This configuration allows you to control the verbosity of error messages to help with troubleshooting PySpark jobs while maintaining appropriate security controls.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The level of detail for error messages returned by the PySpark job. When set to DETAILED, error messages include more information to help troubleshoot issues with your PySpark job.

Because this setting may expose sensitive data, it is recommended for development and testing environments.

_Required_: Yes

_Type_: String

_Allowed values_: `DETAILED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ColumnClassificationDetails

Hash
