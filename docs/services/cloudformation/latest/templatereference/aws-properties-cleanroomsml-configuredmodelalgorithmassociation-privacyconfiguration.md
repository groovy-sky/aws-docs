This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRoomsML::ConfiguredModelAlgorithmAssociation PrivacyConfiguration

Information about the privacy configuration for a configured model algorithm
association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Policies" : PrivacyConfigurationPolicies
}

```

### YAML

```yaml

  Policies:
    PrivacyConfigurationPolicies

```

## Properties

`Policies`

The privacy configuration policies for a configured model algorithm association.

_Required_: Yes

_Type_: [PrivacyConfigurationPolicies](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cleanroomsml-configuredmodelalgorithmassociation-privacyconfigurationpolicies.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

MetricsConfigurationPolicy

PrivacyConfigurationPolicies
