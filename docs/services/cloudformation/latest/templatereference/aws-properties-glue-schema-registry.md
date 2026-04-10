This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Schema Registry

Specifies a registry in the AWS Glue Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "Name" : String
}

```

### YAML

```yaml

  Arn: String
  Name: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the registry.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-(cn|us-gov|iso(-[bef])?))?:glue:.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the registry.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::Schema

SchemaVersion

All content copied from https://docs.aws.amazon.com/.
