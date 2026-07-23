---
title: "AWS::ObservabilityAdmin::TelemetryRule LogDeliveryParameters"
---

This is the new *CloudFormation Template Reference Guide*. Please update your bookmarks and links. For help getting started with CloudFormation, see the [AWS CloudFormation User Guide](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/Welcome.html).

# AWS::ObservabilityAdmin::TelemetryRule LogDeliveryParameters
<a name="aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters"></a>

Configuration parameters for Amazon Bedrock AgentCore logging, including `logType` settings.

## Syntax
<a name="aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters-syntax"></a>

To declare this entity in your CloudFormation template, use the following syntax:

### JSON
<a name="aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters-syntax.json"></a>

```
{
  "[LogTypes](#cfn-observabilityadmin-telemetryrule-logdeliveryparameters-logtypes)" : {{[ String, ... ]}}
}
```

### YAML
<a name="aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters-syntax.yaml"></a>

```
  [LogTypes](#cfn-observabilityadmin-telemetryrule-logdeliveryparameters-logtypes): {{
    - String}}
```

## Properties
<a name="aws-properties-observabilityadmin-telemetryrule-logdeliveryparameters-properties"></a>

`LogTypes`  <a name="cfn-observabilityadmin-telemetryrule-logdeliveryparameters-logtypes"></a>
The type of log that the source is sending.
*Required*: No
*Type*: Array of String
*Allowed values*: `APPLICATION_LOGS | USAGE_LOGS | SECURITY_FINDING_LOGS`
*Minimum*: `1`
*Update requires*: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

All content copied from https://docs.aws.amazon.com/.
