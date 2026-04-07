This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GreengrassV2::Deployment ComponentDeploymentSpecification

Contains information about a component to deploy.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentVersion" : String,
  "ConfigurationUpdate" : ComponentConfigurationUpdate,
  "RunWith" : ComponentRunWith
}

```

### YAML

```yaml

  ComponentVersion: String
  ConfigurationUpdate:
    ComponentConfigurationUpdate
  RunWith:
    ComponentRunWith

```

## Properties

`ComponentVersion`

The version of the component.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ConfigurationUpdate`

The configuration updates to deploy for the component. You can define reset updates and
merge updates. A reset updates the keys that you specify to the default configuration for the
component. A merge updates the core device's component configuration with the keys and values
that you specify. The AWS IoT Greengrass Core software applies reset updates before it
applies merge updates. For more information, see [Update component\
configuration](https://docs.aws.amazon.com/greengrass/v2/developerguide/update-component-configurations.html).

_Required_: No

_Type_: [ComponentConfigurationUpdate](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrassv2-deployment-componentconfigurationupdate.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RunWith`

The system user and group that the software uses to run component
processes on the core device. If you omit this parameter, the software
uses the system user and group that you configure for the core device. For more information,
see [Configure the user and group that run components](https://docs.aws.amazon.com/greengrass/v2/developerguide/configure-greengrass-core-v2.html#configure-component-user) in the _AWS IoT Greengrass V2 Developer Guide_.

_Required_: No

_Type_: [ComponentRunWith](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-greengrassv2-deployment-componentrunwith.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ComponentConfigurationUpdate

ComponentRunWith
