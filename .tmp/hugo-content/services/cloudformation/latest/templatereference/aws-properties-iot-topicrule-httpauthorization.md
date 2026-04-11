This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::TopicRule HttpAuthorization

The authorization method used to send messages.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sigv4" : SigV4Authorization
}

```

### YAML

```yaml

  Sigv4:
    SigV4Authorization

```

## Properties

`Sigv4`

Use Sig V4 authorization. For more information, see [Signature Version 4 Signing\
Process](../../../../general/latest/gr/signature-version-4.md).

_Required_: No

_Type_: [SigV4Authorization](aws-properties-iot-topicrule-sigv4authorization.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HttpActionHeader

IotAnalyticsAction

All content copied from https://docs.aws.amazon.com/.
