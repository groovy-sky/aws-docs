This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApplicationInsights::Application SubComponentTypeConfiguration

The `AWS::ApplicationInsights::Application SubComponentTypeConfiguration` property type specifies the sub-component configurations for a component.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SubComponentConfigurationDetails" : SubComponentConfigurationDetails,
  "SubComponentType" : String
}

```

### YAML

```yaml

  SubComponentConfigurationDetails:
    SubComponentConfigurationDetails
  SubComponentType: String

```

## Properties

`SubComponentConfigurationDetails`

The configuration settings of the sub-components.

_Required_: Yes

_Type_: [SubComponentConfigurationDetails](aws-properties-applicationinsights-application-subcomponentconfigurationdetails.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubComponentType`

The sub-component type.

_Required_: Yes

_Type_: String

_Allowed values_: `AWS::EC2::Instance | AWS::EC2::Volume`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubComponentConfigurationDetails

Tag

All content copied from https://docs.aws.amazon.com/.
