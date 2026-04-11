This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RedshiftServerless::Workgroup ConfigParameter

A array of parameters to set for more control over a serverless database.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterKey" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterKey: String
  ParameterValue: String

```

## Properties

`ParameterKey`

The key of the parameter. The options are `auto_mv`, `datestyle`, `enable_case_sensitive_identifier`,
`enable_user_activity_logging`, `query_group`, `search_path`, `require_ssl`, `use_fips_ssl`,
and query monitoring metrics that let you define performance boundaries. For more information about query monitoring rules and available metrics, see
[Query monitoring metrics for Amazon Redshift Serverless](../../../redshift/latest/dg/cm-c-wlm-query-monitoring-rules.md#cm-c-wlm-query-monitoring-metrics-serverless).

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The value of the parameter to set.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `15000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RedshiftServerless::Workgroup

Endpoint

All content copied from https://docs.aws.amazon.com/.
