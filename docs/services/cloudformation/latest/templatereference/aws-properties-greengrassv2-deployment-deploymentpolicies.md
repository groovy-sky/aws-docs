This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment DeploymentPolicies

Contains information about policies that define how a deployment updates components and
handles failure.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentUpdatePolicy" : DeploymentComponentUpdatePolicy,
  "ConfigurationValidationPolicy" : DeploymentConfigurationValidationPolicy,
  "FailureHandlingPolicy" : String
}

```

### YAML

```yaml

  ComponentUpdatePolicy:
    DeploymentComponentUpdatePolicy
  ConfigurationValidationPolicy:
    DeploymentConfigurationValidationPolicy
  FailureHandlingPolicy: String

```

## Properties

`ComponentUpdatePolicy`

The component update policy for the configuration deployment. This policy defines when
it's safe to deploy the configuration to devices.

_Required_: No

_Type_: [DeploymentComponentUpdatePolicy](aws-properties-greengrassv2-deployment-deploymentcomponentupdatepolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationValidationPolicy`

The configuration validation policy for the configuration deployment. This policy defines
how long each component has to validate its configure updates.

_Required_: No

_Type_: [DeploymentConfigurationValidationPolicy](aws-properties-greengrassv2-deployment-deploymentconfigurationvalidationpolicy.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FailureHandlingPolicy`

The failure handling policy for the configuration deployment. This policy defines what to
do if the deployment fails.

Default: `ROLLBACK`

_Required_: No

_Type_: String

_Allowed values_: `ROLLBACK | DO_NOTHING`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentIoTJobConfiguration

IoTJobAbortConfig

All content copied from https://docs.aws.amazon.com/.
