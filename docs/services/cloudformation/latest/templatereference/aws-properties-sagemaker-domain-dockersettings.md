This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain DockerSettings

A collection of settings that configure the domain's Docker interaction.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnableDockerAccess" : String,
  "VpcOnlyTrustedAccounts" : [ String, ... ]
}

```

### YAML

```yaml

  EnableDockerAccess: String
  VpcOnlyTrustedAccounts:
    - String

```

## Properties

`EnableDockerAccess`

Indicates whether the domain can access Docker.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcOnlyTrustedAccounts`

The list of AWS accounts that are trusted when the domain is created in
VPC-only mode.

_Required_: No

_Type_: Array of String

_Minimum_: `12 | 0`

_Maximum_: `12 | 20`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultSpaceStorageSettings

DomainSettings

All content copied from https://docs.aws.amazon.com/.
