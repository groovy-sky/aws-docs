---
title: "AWS::ObservabilityAdmin::OrganizationTelemetryRule LogDeliveryParameters"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::ObservabilityAdmin::OrganizationTelemetryRule LogDeliveryParameters
<a name="aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters"></a>

Configuration parameters for Amazon Bedrock AgentCore logging, including `logType` settings.

## Syntax
<a name="aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-syntax.json"></a>

```
{
  "[LogTypes](#cfn-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-logtypes)" : {{[ String, ... ]}}
}
```

### YAML
<a name="aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-syntax.yaml"></a>

```
  [LogTypes](#cfn-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-logtypes): {{
    - String}}
```

## Properties
<a name="aws-properties-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-properties"></a>

`LogTypes`  <a name="cfn-observabilityadmin-organizationtelemetryrule-logdeliveryparameters-logtypes"></a>
The type of log that the source is sending.
*Required*: No
*Type*: Array of String
*Allowed values*: `SECURITY_FINDING_LOGS`
*Minimum*: `1`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
