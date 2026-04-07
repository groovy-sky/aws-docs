This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL DataProtectionConfig

Specifies data protection to apply to the web request data for the web ACL. This is a web ACL level data protection option.

The data protection that you configure for the web ACL alters the data that's available for any other data collection activity,
including your AWS WAF logging destinations, web ACL request sampling, and Amazon Security Lake data collection and management. Your other option for data protection is in the logging configuration, which only affects logging.

This is part of the data protection configuration for a web ACL.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataProtections" : [ DataProtect, ... ]
}

```

### YAML

```yaml

  DataProtections:
    - DataProtect

```

## Properties

`DataProtections`

An array of data protection configurations for specific web request field types. This is defined for each
web ACL. AWS WAF applies the specified protection to all web requests that the web ACL inspects.

_Required_: Yes

_Type_: Array of [DataProtect](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-wafv2-webacl-dataprotect.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DataProtect

DefaultAction
