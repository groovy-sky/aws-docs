This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FMS::ResourceSet

A set of resources to include in a policy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FMS::ResourceSet",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Resources" : [ String, ... ],
      "ResourceTypeList" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FMS::ResourceSet
Properties:
  Description: String
  Name: String
  Resources:
    - String
  ResourceTypeList:
    - String
  Tags:
    - Tag

```

## Properties

`Description`

A description of the resource set.

_Required_: No

_Type_: String

_Pattern_: `^([a-zA-Z0-9_.:/=+\-@\s]*)$`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The descriptive name of the resource set. You can't change the name of a resource set after you create it.

_Required_: Yes

_Type_: String

_Pattern_: `^([a-zA-Z0-9_.:/=+\-@\s]+)$`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

Property description not available.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceTypeList`

Determines the resources that can be associated to the resource set. Depending on
your setting for max results and the number of resource sets, a single call might not
return the full list.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fms-resourceset-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Id`

A unique identifier for the resource set. This ID is returned in the responses to create and list commands. You provide it to operations like update and delete.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ThirdPartyFirewallPolicy

Tag
