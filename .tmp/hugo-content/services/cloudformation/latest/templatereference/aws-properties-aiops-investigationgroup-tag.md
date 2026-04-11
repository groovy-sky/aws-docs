This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AIOps::InvestigationGroup Tag

A list of key-value pairs to associate with the investigation group. You can associate
as many as 50 tags with an investigation group. To be able to associate tags when you
create the investigation group, you must have the `cloudwatch:TagResource`
permission.

Tags can help you organize and categorize your resources. You can also use them to scope
user permissions by granting a user permission to access or change only resources with
certain tag values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

Assigns one or more tags (key-value pairs) to the specified resource.

Tags can help you organize and categorize your resources. You can also use them to scope
user permissions by granting a user permission to access or change only resources with
certain tag values.

Tags don't have any semantic meaning to AWS and are interpreted strictly
as strings of characters.

You can associate as many as 50 tags with a resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A list of key-value pairs to associate with the investigation group. You can associate
as many as 50 tags with an investigation group. To be able to associate tags when you
create the investigation group, you must have the `cloudwatch:TagResource`
permission.

Tags can help you organize and categorize your resources. You can also use them to scope
user permissions by granting a user permission to access or change only resources with
certain tag values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfigMap

Next

All content copied from https://docs.aws.amazon.com/.
