This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::CodeSecurityScanConfiguration ScopeSettings

The scope settings that define which repositories will be scanned. If the
`ScopeSetting` parameter is `ALL` the scan configuration applies
to all existing and future projects imported into Amazon Inspector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "projectSelectionScope" : String
}

```

### YAML

```yaml

  projectSelectionScope: String

```

## Properties

`projectSelectionScope`

The scope of projects to be selected for scanning within the integrated
repositories.

_Required_: No

_Type_: String

_Allowed values_: `ALL`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PeriodicScanConfiguration

AWS::InspectorV2::Filter
