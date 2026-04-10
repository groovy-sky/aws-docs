This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SDB::Domain

Use the `AWS::SDB::Domain` resource to declare a SimpleDB domain.
When you specify `AWS::SDB::Domain` as an argument in a `Ref` function,
AWS CloudFormation returns the value of the `DomainName`.

###### Important

The `AWS::SDB::Domain` resource does not allow any updates, including
metadata updates.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SDB::Domain",
  "Properties" : {
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::SDB::Domain
Properties:
  Description: String

```

## Properties

`Description`

Information about the SimpleDB domain.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of the `DomainName`, such as:
`Domain1-123456789012`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon SimpleDB

Next

All content copied from https://docs.aws.amazon.com/.
