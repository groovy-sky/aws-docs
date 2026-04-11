This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Omics::WorkflowVersion RegistryMapping

If you are using the ECR pull through cache feature, the registry mapping
maps between the ECR repository and the upstream registry where container images
are pulled and synchronized.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EcrAccountId" : String,
  "EcrRepositoryPrefix" : String,
  "UpstreamRegistryUrl" : String,
  "UpstreamRepositoryPrefix" : String
}

```

### YAML

```yaml

  EcrAccountId: String
  EcrRepositoryPrefix: String
  UpstreamRegistryUrl: String
  UpstreamRepositoryPrefix: String

```

## Properties

`EcrAccountId`

Account ID of the account that owns the upstream container image.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]+$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EcrRepositoryPrefix`

The repository prefix to use in the ECR private repository.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpstreamRegistryUrl`

The URI of the upstream registry.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `1`

_Maximum_: `750`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpstreamRepositoryPrefix`

The repository prefix of the corresponding repository in the upstream registry.

_Required_: No

_Type_: String

_Pattern_: `^[\p{L}||\p{M}||\p{Z}||\p{S}||\p{N}||\p{P}]+$`

_Minimum_: `2`

_Maximum_: `30`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageMapping

SourceReference

All content copied from https://docs.aws.amazon.com/.
