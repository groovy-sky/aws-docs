This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Kendra::Index JsonTokenTypeConfiguration

Provides the configuration information for the JSON token type.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupAttributeField" : String,
  "UserNameAttributeField" : String
}

```

### YAML

```yaml

  GroupAttributeField: String
  UserNameAttributeField: String

```

## Properties

`GroupAttributeField`

The group attribute field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserNameAttributeField`

The user name attribute field.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentMetadataConfiguration

JwtTokenTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
