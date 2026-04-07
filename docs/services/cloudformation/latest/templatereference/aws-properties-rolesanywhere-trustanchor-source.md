This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RolesAnywhere::TrustAnchor Source

Object representing the TrustAnchor type and its related certificate data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceData" : SourceData,
  "SourceType" : String
}

```

### YAML

```yaml

  SourceData:
    SourceData
  SourceType: String

```

## Properties

`SourceData`

A union object representing the data field of the TrustAnchor depending on its type

_Required_: Yes

_Type_: [SourceData](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rolesanywhere-trustanchor-sourcedata.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceType`

The type of the TrustAnchor.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS_ACM_PCA | CERTIFICATE_BUNDLE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotificationSetting

SourceData
