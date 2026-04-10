This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WAFv2::WebACL BlockAction

Specifies that AWS WAF should block the request and optionally defines
additional custom handling for the response to the web request.

This is used in the context of other settings, for example to specify values for a rule
action or a web ACL default action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomResponse" : CustomResponse
}

```

### YAML

```yaml

  CustomResponse:
    CustomResponse

```

## Properties

`CustomResponse`

Defines a custom response for the web request.

For information about customizing web requests and responses,
see [Customizing web requests and responses in AWS WAF](../../../waf/latest/developerguide/waf-custom-request-response.md)
in the _AWS WAF Developer Guide_.

_Required_: No

_Type_: [CustomResponse](aws-properties-wafv2-webacl-customresponse.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

- [Set an block action](#aws-properties-wafv2-webacl-blockaction--examples--Set_an_block_action)

- [Set a block action with a custom response setting](#aws-properties-wafv2-webacl-blockaction--examples--Set_a_block_action_with_a_custom_response_setting)

### Set an block action

The following shows an example block action specification.

#### YAML

```yaml

Action:
  Block: {}

```

#### JSON

```json

"Action": {
  "Block": {}
}
```

### Set a block action with a custom response setting

The following shows an example block action specification with a custom response
setting.

#### YAML

```yaml

Block:
  CustomResponse:
    ResponseCode: 503
    CustomResponseBodyKey: CustomResponseBodyKey1
    ResponseHeaders:
      - Name: BlockActionHeader1Name
        Value: BlockActionHeader1Value
      - Name: BlockActionHeader2Name
        Value: BlockActionHeader2Value
```

#### JSON

```json

"Block": {
  "CustomResponse": {
    "ResponseCode": 503,
    "CustomResponseBodyKey": "CustomResponseBodyKey1",
    "ResponseHeaders": [
      {
        "Name": "BlockActionHeader1Name",
        "Value": "BlockActionHeader1Value"
      },
      {
        "Name": "BlockActionHeader2Name",
        "Value": "BlockActionHeader2Value"
      }
    ]
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWSManagedRulesBotControlRuleSet

Body

All content copied from https://docs.aws.amazon.com/.
