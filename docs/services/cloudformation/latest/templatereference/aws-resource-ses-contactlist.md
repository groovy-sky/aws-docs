This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ContactList

A list that contains contacts that have subscribed to a particular topic or
topics.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::ContactList",
  "Properties" : {
      "ContactListName" : String,
      "Description" : String,
      "Tags" : [ Tag, ... ],
      "Topics" : [ Topic, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::SES::ContactList
Properties:
  ContactListName: String
  Description: String
  Tags:
    - Tag
  Topics:
    - Topic

```

## Properties

`ContactListName`

The name of the contact list.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9_-]{1,64}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

A description of what the contact list is about.

_Required_: No

_Type_: String

_Maximum_: `500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags associated with a contact list.

_Required_: No

_Type_: Array of [Tag](aws-properties-ses-contactlist-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Topics`

An interest group, theme, or label within a list. A contact list can have multiple
topics.

_Required_: No

_Type_: Array of [Topic](aws-properties-ses-contactlist-topic.md)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic Ref function, Ref
returns the resource name. For example:

`{ "Ref" : "ContactListName" }`

For the Amazon SES ContactList, `Ref` returns the name of the contact
list.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnsDestination

Tag

All content copied from https://docs.aws.amazon.com/.
