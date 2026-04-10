This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMPool SourceResource

The resource used to provision CIDRs to a resource planning pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResourceId" : String,
  "ResourceOwner" : String,
  "ResourceRegion" : String,
  "ResourceType" : String
}

```

### YAML

```yaml

  ResourceId: String
  ResourceOwner: String
  ResourceRegion: String
  ResourceType: String

```

## Properties

`ResourceId`

The source resource ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceOwner`

The source resource owner.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceRegion`

The source resource Region.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

The source resource type.

_Required_: Yes

_Type_: String

_Allowed values_: `vpc`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedCidr

Tag

All content copied from https://docs.aws.amazon.com/.
