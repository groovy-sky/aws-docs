This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::IdMappingTable IdMappingTableInputReferenceConfig

Provides the input reference configuration for the ID mapping table.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InputReferenceArn" : String,
  "ManageResourcePolicies" : Boolean
}

```

### YAML

```yaml

  InputReferenceArn: String
  ManageResourcePolicies: Boolean

```

## Properties

`InputReferenceArn`

The Amazon Resource Name (ARN) of the referenced resource in AWS Entity Resolution. Valid values are ID mapping workflow ARNs.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManageResourcePolicies`

When `TRUE`, AWS Clean Rooms manages permissions for the ID mapping table resource.

When `FALSE`, the resource owner manages permissions for the ID mapping table resource.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CleanRooms::IdMappingTable

IdMappingTableInputReferenceProperties

All content copied from https://docs.aws.amazon.com/.
