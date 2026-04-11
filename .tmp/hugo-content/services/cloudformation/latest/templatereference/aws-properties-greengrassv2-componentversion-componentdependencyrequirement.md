This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::ComponentVersion ComponentDependencyRequirement

Contains information about a component dependency for a Lambda function
component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DependencyType" : String,
  "VersionRequirement" : String
}

```

### YAML

```yaml

  DependencyType: String
  VersionRequirement: String

```

## Properties

`DependencyType`

The type of this dependency. Choose from the following options:

- `SOFT` – The component doesn't restart if the dependency changes
state.

- `HARD` – The component restarts if the dependency changes state.

Default: `HARD`

_Required_: No

_Type_: String

_Allowed values_: `SOFT | HARD`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VersionRequirement`

The component version requirement for the component dependency.

AWS IoT Greengrass uses semantic version constraints. For more information, see
[Semantic Versioning](https://semver.org/).

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GreengrassV2::ComponentVersion

ComponentPlatform

All content copied from https://docs.aws.amazon.com/.
