This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisAnalyticsV2::Application MavenReference

The information required to specify a Maven reference. You can use Maven references to
specify dependency JAR files.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ArtifactId" : String,
  "GroupId" : String,
  "Version" : String
}

```

### YAML

```yaml

  ArtifactId: String
  GroupId: String
  Version: String

```

## Properties

`ArtifactId`

The artifact ID of the Maven reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GroupId`

The group ID of the Maven reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the Maven reference.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MappingParameters

MonitoringConfiguration

All content copied from https://docs.aws.amazon.com/.
