This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::KinesisFirehose::DeliveryStream SnowflakeVpcConfiguration

Configure a Snowflake VPC

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PrivateLinkVpceId" : String
}

```

### YAML

```yaml

  PrivateLinkVpceId: String

```

## Properties

`PrivateLinkVpceId`

The VPCE ID for Firehose to privately connect with Snowflake. The ID format is
com.amazonaws.vpce.\[region\].vpce-svc-<\[id\]>. For more information, see [Amazon PrivateLink & Snowflake](https://docs.snowflake.com/en/user-guide/admin-security-privatelink)

_Required_: Yes

_Type_: String

_Pattern_: `([a-zA-Z0-9\-\_]+\.){2,3}vpce\.[a-zA-Z0-9\-]*\.vpce-svc\-[a-zA-Z0-9\-]{17}$`

_Minimum_: `47`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SnowflakeRoleConfiguration

SplunkBufferingHints

All content copied from https://docs.aws.amazon.com/.
