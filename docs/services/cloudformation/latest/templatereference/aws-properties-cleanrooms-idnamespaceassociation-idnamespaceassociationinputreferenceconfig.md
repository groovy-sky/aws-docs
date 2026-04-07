This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::IdNamespaceAssociation IdNamespaceAssociationInputReferenceConfig

Provides the information for the ID namespace association input reference configuration.

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

The Amazon Resource Name (ARN) of the AWS Entity Resolution resource that is being associated to the collaboration. Valid resource ARNs are from the ID namespaces that you own.

_Required_: Yes

_Type_: String

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ManageResourcePolicies`

When `TRUE`, AWS Clean Rooms manages permissions for the ID namespace association resource.

When `FALSE`, the resource owner manages permissions for the ID namespace association resource.

_Required_: Yes

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IdMappingConfig

IdNamespaceAssociationInputReferenceProperties
