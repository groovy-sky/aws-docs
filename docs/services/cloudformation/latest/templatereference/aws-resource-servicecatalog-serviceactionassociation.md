This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ServiceCatalog::ServiceActionAssociation

A self-service action association consisting of the Action ID, the Product ID, and the Provisioning Artifact ID.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ServiceCatalog::ServiceActionAssociation",
  "Properties" : {
      "ProductId" : String,
      "ProvisioningArtifactId" : String,
      "ServiceActionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ServiceCatalog::ServiceActionAssociation
Properties:
  ProductId: String
  ProvisioningArtifactId: String
  ServiceActionId: String

```

## Properties

`ProductId`

The product identifier. For example, `prod-abcdzk7xy33qa`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]{1,99}\Z`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisioningArtifactId`

The identifier of the provisioning artifact. For example, `pa-4abcdjnxjj6ne`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]{1,99}\Z`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ServiceActionId`

The self-service action identifier. For example, `act-fs7abcd89wxyz`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9][a-zA-Z0-9_-]{1,99}\Z`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DefinitionParameter

AWS::ServiceCatalog::StackSetConstraint
