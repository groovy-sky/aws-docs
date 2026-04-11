This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMR::WALWorkspace

The `AWS::EMR::WALWorkspace` resource Property description not available. for EMR.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EMR::WALWorkspace",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "WALWorkspaceName" : String
    }
}

```

### YAML

```yaml

Type: AWS::EMR::WALWorkspace
Properties:
  Tags:
    - Tag
  WALWorkspaceName: String

```

## Properties

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-emr-walworkspace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WALWorkspaceName`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::EMR::StudioSessionMapping

Tag

All content copied from https://docs.aws.amazon.com/.
