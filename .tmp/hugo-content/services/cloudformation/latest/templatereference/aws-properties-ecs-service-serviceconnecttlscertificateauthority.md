This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service ServiceConnectTlsCertificateAuthority

The certificate root authority that secures your service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsPcaAuthorityArn" : String
}

```

### YAML

```yaml

  AwsPcaAuthorityArn: String

```

## Properties

`AwsPcaAuthorityArn`

The ARN of the AWS Private Certificate Authority certificate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceConnectTestTrafficRulesHeaderValue

ServiceConnectTlsConfiguration

All content copied from https://docs.aws.amazon.com/.
