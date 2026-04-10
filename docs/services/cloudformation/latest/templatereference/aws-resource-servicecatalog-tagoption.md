This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::TagOption

Specifies a TagOption. A TagOption is a key-value pair managed by AWS Service Catalog
that serves as a template for creating an AWS tag.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::TagOption",
  "Properties" : {
      "Active" : Boolean,
      "Key" : String,
      "Value" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::TagOption
Properties:
  Active: Boolean
  Key: String
  Value: String

```

## Properties

`Active`

The TagOption active state.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Key`

The TagOption key.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The TagOption value.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the TagOption identifier.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Id`

The TagOption identifier.

## See also

- [CreateTagOption](../../../servicecatalog/latest/dg/api-createtagoption.md)
in the _AWS Service Catalog API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::StackSetConstraint

AWS::ServiceCatalog::TagOptionAssociation

All content copied from https://docs.aws.amazon.com/.
