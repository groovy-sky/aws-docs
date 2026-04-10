This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Domain RStudioServerProAppSettings

A collection of settings that configure user interaction with the
`RStudioServerPro` app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessStatus" : String,
  "UserGroup" : String
}

```

### YAML

```yaml

  AccessStatus: String
  UserGroup: String

```

## Properties

`AccessStatus`

Indicates whether the current user has access to the `RStudioServerPro`
app.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserGroup`

The level of permissions that the user has within the `RStudioServerPro` app.
This value defaults to \`User\`. The \`Admin\` value allows the user access to the RStudio
Administrative Dashboard.

_Required_: No

_Type_: String

_Allowed values_: `R_STUDIO_ADMIN | R_STUDIO_USER`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RSessionAppSettings

RStudioServerProDomainSettings

All content copied from https://docs.aws.amazon.com/.
