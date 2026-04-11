This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EKS::IdentityProviderConfig Tag

The metadata that you apply to a resource to help you categorize and organize them.
Each tag consists of a key and an optional value. You define them.

The following basic restrictions apply to tags:

- Maximum number of tags per resource – 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length – 128 Unicode characters in UTF-8

- Maximum value length – 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with
this prefix do not count against your tags per resource limit.

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

One part of a key-value pair that make up a tag. A `key` is a general label
that acts like a category for more specific tag values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that make up a tag. A `value` acts as
a descriptor within a tag category (key).

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

#### JSON

```json

"Tags" : [
   {
      "Key" : "keyname1",
      "Value" : "value1"
   },
   {
      "Key" : "keyname2",
      "Value" : "value2"
   }
]
```

#### YAML

```yaml

Tags:
  - Key: "keyname1"
    Value: "value1"
  - Key: "keyname2"
    Value: "value2"
```

## See also

- [Setting CloudFormation stack options](../userguide/cfn-console-add-tags.md)

- [Viewing CloudFormation stack data and resources on the AWS Management Console](../userguide/cfn-console-view-stack-data-resources.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RequiredClaim

AWS::EKS::Nodegroup

All content copied from https://docs.aws.amazon.com/.
