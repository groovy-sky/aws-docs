This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::JobTemplate AbortConfig

The criteria that determine when and how a job abort takes place.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CriteriaList" : [ AbortCriteria, ... ]
}

```

### YAML

```yaml

  CriteriaList:
    - AbortCriteria

```

## Properties

`CriteriaList`

The list of criteria that determine when and how to abort the job.

_Required_: Yes

_Type_: Array of [AbortCriteria](aws-properties-iot-jobtemplate-abortcriteria.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoT::JobTemplate

AbortCriteria

All content copied from https://docs.aws.amazon.com/.
