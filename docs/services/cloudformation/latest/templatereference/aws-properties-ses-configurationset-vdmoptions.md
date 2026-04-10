This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::ConfigurationSet VdmOptions

The Virtual Deliverability Manager (VDM) options that apply to a configuration
set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DashboardOptions" : DashboardOptions,
  "GuardianOptions" : GuardianOptions
}

```

### YAML

```yaml

  DashboardOptions:
    DashboardOptions
  GuardianOptions:
    GuardianOptions

```

## Properties

`DashboardOptions`

Specifies additional settings for your VDM configuration as applicable to the
Dashboard.

_Required_: No

_Type_: [DashboardOptions](aws-properties-ses-configurationset-dashboardoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardianOptions`

Specifies additional settings for your VDM configuration as applicable to the
Guardian.

_Required_: No

_Type_: [GuardianOptions](aws-properties-ses-configurationset-guardianoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ValidationOptions

AWS::SES::ConfigurationSetEventDestination

All content copied from https://docs.aws.amazon.com/.
