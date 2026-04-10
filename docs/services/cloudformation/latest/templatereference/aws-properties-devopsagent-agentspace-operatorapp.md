This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::AgentSpace OperatorApp

Configuration for the DevOps Agent web app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Iam" : IamAuthConfiguration,
  "Idc" : IdcAuthConfiguration
}

```

### YAML

```yaml

  Iam:
    IamAuthConfiguration
  Idc:
    IdcAuthConfiguration

```

## Properties

`Iam`

IAM-based authentication configuration for the DevOps Agent web app.

_Required_: No

_Type_: [IamAuthConfiguration](aws-properties-devopsagent-agentspace-iamauthconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Idc`

IAM Identity Center authentication configuration for the DevOps Agent web app.

_Required_: No

_Type_: [IdcAuthConfiguration](aws-properties-devopsagent-agentspace-idcauthconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IdcAuthConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
