This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::DataSource FormInput

The details of a metadata form.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "FormName" : String,
  "TypeIdentifier" : String,
  "TypeRevision" : String
}

```

### YAML

```yaml

  Content: String
  FormName: String
  TypeIdentifier: String
  TypeRevision: String

```

## Properties

`Content`

The content of the metadata form.

_Required_: No

_Type_: String

_Maximum_: `75000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FormName`

The name of the metadata form.

_Required_: Yes

_Type_: String

_Pattern_: `^(?![0-9_])\w+$|^_\w*[a-zA-Z0-9]\w*$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeIdentifier`

The ID of the metadata form type.

_Required_: No

_Type_: String

_Pattern_: `^(?!\.)[\w\.]*\w$`

_Minimum_: `1`

_Maximum_: `385`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeRevision`

The revision of the metadata form type.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FilterExpression

GlueRunConfigurationInput

All content copied from https://docs.aws.amazon.com/.
