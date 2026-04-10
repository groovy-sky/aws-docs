This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space OwnershipSettings

The collection of ownership settings for a space.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OwnerUserProfileName" : String
}

```

### YAML

```yaml

  OwnerUserProfileName: String

```

## Properties

`OwnerUserProfileName`

The user profile who is the owner of the space.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9](-*[a-zA-Z0-9]){0,62}`

_Maximum_: `63`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

KernelGatewayAppSettings

ResourceSpec

All content copied from https://docs.aws.amazon.com/.
