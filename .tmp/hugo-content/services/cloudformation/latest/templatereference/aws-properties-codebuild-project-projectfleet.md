This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeBuild::Project ProjectFleet

Information about the compute fleet of the build project. For more
information, see [Working \
with reserved capacity in AWS CodeBuild](../../../codebuild/latest/userguide/fleets.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FleetArn" : String
}

```

### YAML

```yaml

  FleetArn: String

```

## Properties

`FleetArn`

Specifies the compute fleet ARN for the build project.

_Required_: No

_Type_: String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProjectFileSystemLocation

ProjectSourceVersion

All content copied from https://docs.aws.amazon.com/.
