This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IAM::ServerCertificate Tag

A structure that represents user-provided metadata that can be associated with an IAM
resource. For more information about tagging, see [Tagging IAM resources](../../../iam/latest/userguide/id-tags.md) in the
_IAM User Guide_.

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

The key name that can be used to look up or retrieve the associated value. For example,
`Department` or `Cost Center` are common choices.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value associated with this tag. For example, tags with a key name of
`Department` could have values such as `Human Resources`,
`Accounting`, and `Support`. Tags with a key name of `Cost
        Center` might have values that consist of the number associated with the different
cost centers in your company. Typically, many resources have tags with the same key name but
with different values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IAM::ServerCertificate

AWS::IAM::ServiceLinkedRole

All content copied from https://docs.aws.amazon.com/.
