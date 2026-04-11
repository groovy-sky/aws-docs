This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Organizations::Policy Tag

A custom key-value pair associated with a resource within your organization.

You can attach tags to any of the following organization resources.

- AWS account

- Organizational unit (OU)

- Organization root

- Policy

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

The key identifier, or name, of the tag.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The string value that's associated with the key of the tag. You can set the value of a
tag to an empty string, but you can't set the value of a tag to null.

_Required_: Yes

_Type_: String

_Pattern_: `[\s\S]*`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Organizations::Policy

AWS::Organizations::ResourcePolicy

All content copied from https://docs.aws.amazon.com/.
