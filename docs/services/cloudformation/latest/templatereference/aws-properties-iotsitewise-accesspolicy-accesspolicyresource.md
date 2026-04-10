This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AccessPolicy AccessPolicyResource

The AWS IoT SiteWise Monitor resource for this access policy. Choose either a portal or a project.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Portal" : Portal,
  "Project" : Project
}

```

### YAML

```yaml

  Portal:
    Portal
  Project:
    Project

```

## Properties

`Portal`

Identifies an AWS IoT SiteWise Monitor portal.

_Required_: No

_Type_: [Portal](aws-properties-iotsitewise-accesspolicy-portal.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Project`

Identifies a specific AWS IoT SiteWise Monitor project.

_Required_: No

_Type_: [Project](aws-properties-iotsitewise-accesspolicy-project.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccessPolicyIdentity

IamRole

All content copied from https://docs.aws.amazon.com/.
