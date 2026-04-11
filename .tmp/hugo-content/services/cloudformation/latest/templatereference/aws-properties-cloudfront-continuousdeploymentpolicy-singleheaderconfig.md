This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy SingleHeaderConfig

Determines which HTTP requests are sent to the staging distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Header" : String,
  "Value" : String
}

```

### YAML

```yaml

  Header: String
  Value: String

```

## Properties

`Header`

The request header name that you want CloudFront to send to your staging
distribution. The header must contain the prefix `aws-cf-cd-`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The request header value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1783`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SessionStickinessConfig

SingleHeaderPolicyConfig

All content copied from https://docs.aws.amazon.com/.
