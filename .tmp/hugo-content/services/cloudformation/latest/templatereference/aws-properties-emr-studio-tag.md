This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::Studio Tag

A key-value pair containing user-defined metadata that you can associate with an Amazon EMR resource. Tags make it easier to associate clusters in various ways, such as
grouping clusters to track your Amazon EMR resource allocation costs. For more
information, see [Tag Clusters](../../../emr/latest/managementguide/emr-plan-tags.md).

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

A user-defined key, which is the minimum required information for a valid tag. For more
information, see [Tag](../../../emr/latest/managementguide/emr-plan-tags.md).

_Required_: Yes

_Type_: String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]+$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A user-defined value, which is optional in a tag. For more information, see [Tag\
Clusters](../../../emr/latest/managementguide/emr-plan-tags.md).

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z+-=._:/]+$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::Studio

AWS::EMR::StudioSessionMapping

All content copied from https://docs.aws.amazon.com/.
