This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::TagOptionAssociation

Associate the specified TagOption with the specified portfolio or product.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::TagOptionAssociation",
  "Properties" : {
      "ResourceId" : String,
      "TagOptionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::TagOptionAssociation
Properties:
  ResourceId: String
  TagOptionId: String

```

## Properties

`ResourceId`

The resource identifier.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TagOptionId`

The TagOption identifier.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns an identifier for the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [AssociateTagOptionWithResource](../../../servicecatalog/latest/dg/api-associatetagoptionwithresource.md) in the _AWS Service Catalog_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ServiceCatalog::TagOption

Next

All content copied from https://docs.aws.amazon.com/.
