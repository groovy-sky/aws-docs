This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset VersioningConfiguration

Information about the versioning of dataset contents.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxVersions" : Integer,
  "Unlimited" : Boolean
}

```

### YAML

```yaml

  MaxVersions: Integer
  Unlimited: Boolean

```

## Properties

`MaxVersions`

How many versions of dataset contents are kept. The `unlimited` parameter must
be `false`.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unlimited`

If true, unlimited versions of dataset contents are kept.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Variable

AWS::IoTAnalytics::Datastore
