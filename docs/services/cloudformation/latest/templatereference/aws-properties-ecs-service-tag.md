This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::Service Tag

The metadata that you apply to a resource to help you categorize and organize them. Each tag consists
of a key and an optional value. You define them.

The following basic restrictions apply to tags:

- Maximum number of tags per resource - 50

- For each resource, each tag key must be unique, and each tag key can have only
one value.

- Maximum key length - 128 Unicode characters in UTF-8

- Maximum value length - 256 Unicode characters in UTF-8

- If your tagging schema is used across multiple services and resources,
remember that other services may have restrictions on allowed characters.
Generally allowed characters are: letters, numbers, and spaces representable in
UTF-8, and the following characters: + - = . \_ : / @.

- Tag keys and values are case-sensitive.

- Do not use `aws:`, `AWS:`, or any upper or lowercase
combination of such as a prefix for either keys or values as it is reserved for
AWS use. You cannot edit or delete tag keys or values with this prefix. Tags with
this prefix do not count against your tags per resource limit.

In order to tag a service that has the following ARN format, you need to migrate the
service to the long ARN. You must use the API, CLI or console to migrate the service ARN.
For more information, see [Migrate an Amazon ECS short service ARN to a long ARN](../../../amazonecs/latest/developerguide/service-arn-migration.md) in the _Amazon Elastic Container Service Developer Guide_.

`arn:aws:ecs:region:aws_account_id:service/service-name`

After the migration is complete, the following are true:

- The service ARN is:
`arn:aws:ecs:region:aws_account_id:service/cluster-name/service-name`

- You can use CloudFormation to tag the service as you would a service with a long ARN format.

- When the `PhysicalResourceId` in the
CloudFormation stack represents a service, the value does not change and will be the
short service ARN.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

One part of a key-value pair that make up a tag. A `key` is a general label
that acts like a category for more specific tag values.

_Required_: No

_Type_: String

_Pattern_: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The optional part of a key-value pair that make up a tag. A `value` acts as
a descriptor within a tag category (key).

_Required_: No

_Type_: String

_Pattern_: `([\p{L}\p{Z}\p{N}_.:/=+\-@]*)`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Associate an Application Load Balancer with a service](../userguide/aws-resource-ecs-service.md#aws-resource-ecs-service--examples--Associate_an_Application_Load_Balancer_with_a_service)

- [Migrate an\
Amazon ECS short service ARN to a long ARN](../../../amazonecs/latest/developerguide/service-arn-migration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ServiceVolumeConfiguration

TimeoutConfiguration

All content copied from https://docs.aws.amazon.com/.
