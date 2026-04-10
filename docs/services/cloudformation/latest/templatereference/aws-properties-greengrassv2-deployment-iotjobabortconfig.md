This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment IoTJobAbortConfig

Contains a list of criteria that define when and how to cancel a configuration
deployment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CriteriaList" : [ IoTJobAbortCriteria, ... ]
}

```

### YAML

```yaml

  CriteriaList:
    - IoTJobAbortCriteria

```

## Properties

`CriteriaList`

The list of criteria that define when and how to cancel the configuration
deployment.

_Required_: Yes

_Type_: Array of [IoTJobAbortCriteria](aws-properties-greengrassv2-deployment-iotjobabortcriteria.md)

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentPolicies

IoTJobAbortCriteria

All content copied from https://docs.aws.amazon.com/.
