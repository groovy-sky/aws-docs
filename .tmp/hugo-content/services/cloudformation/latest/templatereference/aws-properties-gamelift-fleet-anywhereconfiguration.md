This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::Fleet AnywhereConfiguration

Amazon GameLift Servers configuration options for your Anywhere fleets.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Cost" : String
}

```

### YAML

```yaml

  Cost: String

```

## Properties

`Cost`

The cost to run your fleet per hour. Amazon GameLift Servers uses the provided cost of your fleet to
balance usage in queues. For more information about queues, see [Setting\
up queues](../../../gamelift/latest/developerguide/queues-intro.md) in the _Amazon GameLift Servers Developer Guide_.

_Required_: Yes

_Type_: String

_Pattern_: `^\d{1,5}(?:\.\d{1,5})?$`

_Minimum_: `1`

_Maximum_: `11`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [Create GameLift resources using Amazon CloudFront](../../../gamelift/latest/developerguide/resources-cloudformation.md) in the _Amazon_
_GameLift Developer Guide_

- [Create an Amazon GameLift Anywhere fleet](../../../gamelift/latest/developerguide/fleets-creating-anywhere.md) in the _Amazon_
_GameLift Developer Guide_

- [LocationModel](../../../../reference/gamelift/latest/apireference/api-locationmodel.md)
in the _Amazon GameLift API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::GameLift::Fleet

CertificateConfiguration

All content copied from https://docs.aws.amazon.com/.
