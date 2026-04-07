This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CertificateManager::Account

The `AWS::CertificateManager::Account` resource defines the expiry event
configuration that determines the number of days prior to expiry when ACM starts
generating EventBridge events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CertificateManager::Account",
  "Properties" : {
      "ExpiryEventsConfiguration" : ExpiryEventsConfiguration
    }
}

```

### YAML

```yaml

Type: AWS::CertificateManager::Account
Properties:
  ExpiryEventsConfiguration:
    ExpiryEventsConfiguration

```

## Properties

`ExpiryEventsConfiguration`

Object containing expiration events options associated with an AWS account. For more information, see [ExpiryEventsConfiguration](../../../../reference/acm/latest/apireference/api-expiryeventsconfiguration.md) in the API reference.

_Required_: Yes

_Type_: [ExpiryEventsConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-certificatemanager-account-expiryeventsconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the AWS account that owns the
certificate.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`AccountId`

ID of the AWS account that owns the certificate.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Certificate Manager

ExpiryEventsConfiguration
