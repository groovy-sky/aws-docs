This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::MitigationAction ReplaceDefaultPolicyVersionParams

Parameters to define a mitigation action that adds a blank policy to restrict permissions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TemplateName" : String
}

```

### YAML

```yaml

  TemplateName: String

```

## Properties

`TemplateName`

The name of the template to be applied. The only supported value is `BLANK_POLICY`.

_Required_: Yes

_Type_: String

_Allowed values_: `BLANK_POLICY | UNSET_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PublishFindingToSnsParams

Tag

All content copied from https://docs.aws.amazon.com/.
