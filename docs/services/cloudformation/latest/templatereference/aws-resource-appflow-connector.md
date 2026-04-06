This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Connector

Creates a new connector profile associated with your AWS account. There is
a soft quota of 100 connector profiles per AWS account. If you need more
connector profiles than this quota allows, you can submit a request to the Amazon AppFlow
team through the Amazon AppFlow support channel. In each connector profile that you
create, you can provide the credentials and properties for only one connector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::AppFlow::Connector",
  "Properties" : {
      "ConnectorLabel" : String,
      "ConnectorProvisioningConfig" : ConnectorProvisioningConfig,
      "ConnectorProvisioningType" : String,
      "Description" : String
    }
}

```

### YAML

```yaml

Type: AWS::AppFlow::Connector
Properties:
  ConnectorLabel: String
  ConnectorProvisioningConfig:
    ConnectorProvisioningConfig
  ConnectorProvisioningType: String
  Description: String

```

## Properties

`ConnectorLabel`

The label used for registering the connector.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9][\w!@#.-]+`

_Maximum_: `512`

_Update requires_: [Replacement](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-replacement)

`ConnectorProvisioningConfig`

The configuration required for registering the connector.

_Required_: Yes

_Type_: [ConnectorProvisioningConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-appflow-connector-connectorprovisioningconfig.html)

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`ConnectorProvisioningType`

The provisioning type used to register the connector.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9][\w!@#.-]+`

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

`Description`

A description about the connector runtime setting.

_Required_: No

_Type_: String

_Pattern_: `[\s\w/!@#+=.-]*`

_Maximum_: `2048`

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ConnectorArn`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon AppFlow

ConnectorProvisioningConfig
