This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::Portal PortalTypeEntry

The `PortalTypeEntry` property type specifies Property description not available. for an [AWS::IoTSiteWise::Portal](aws-resource-iotsitewise-portal.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PortalTools" : [ String, ... ]
}

```

### YAML

```yaml

  PortalTools:
    - String

```

## Properties

`PortalTools`

The array of tools associated with the specified portal type. The possible values are
`ASSISTANT` and `DASHBOARD`.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Alarms

Tag
