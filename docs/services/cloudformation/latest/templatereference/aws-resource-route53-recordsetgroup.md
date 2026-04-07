This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53::RecordSetGroup

A complex type that contains an optional comment, the name and ID of the hosted zone that you want to make changes in,
and values for the records that you want to create.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53::RecordSetGroup",
  "Properties" : {
      "Comment" : String,
      "HostedZoneId" : String,
      "HostedZoneName" : String,
      "RecordSets" : [ RecordSet, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53::RecordSetGroup
Properties:
  Comment: String
  HostedZoneId: String
  HostedZoneName: String
  RecordSets:
    - RecordSet

```

## Properties

`Comment`

_Optional:_ Any comments you want to include about a change batch
request.

_Required_: No

_Type_: String

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HostedZoneId`

The ID of the hosted zone that you want to create records in.

Specify either `HostedZoneName` or `HostedZoneId`, but not both. If you have multiple hosted zones
with the same domain name, you must specify the hosted zone using `HostedZoneId`.

_Required_: No

_Type_: String

_Maximum_: `32`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostedZoneName`

The name of the hosted zone that you want to create records in. You must include a trailing dot (for example, `www.example.com.`) as part
of the `HostedZoneName`.

When you create a stack using an `AWS::Route53::RecordSet` that specifies `HostedZoneName`,
AWS CloudFormation attempts to find a hosted zone whose name matches the `HostedZoneName`. If AWS CloudFormation
can't find a hosted zone with a matching domain name, or if there is more than one hosted zone with the specified domain name,
AWS CloudFormation will not create the stack.

Specify either `HostedZoneName` or `HostedZoneId`, but not both. If you have multiple hosted zones
with the same domain name, you must specify the hosted zone using `HostedZoneId`.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RecordSets`

A complex type that contains one `RecordSet` element for each record that you want to create.

_Required_: No

_Type_: Array of [RecordSet](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53-recordsetgroup-recordset.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the record set group, for example, `MyRecordSetGroup`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

This element contains an ID that you use when performing a `GetChange` action to get detailed information about the change.

## Examples

For more examples, see
[Route 53 Template Snippets](../userguide/quickref-route53.md).

### Creating records for a mail server

The following example shows how to create three records for a mail server:

- An A record that specifies the IP address for the mail server.

- An MX record that routes email to that server.

- A TXT record that contains an SPF string, which is used to identify the sender of email messages.
SPF records are no longer recommended. For more information, see
[SPF Record Type](../../../route53/latest/developerguide/resourcerecordtypes.md#SPFFormat)
in the _Amazon Route 53 Developer Guide_.

#### JSON

```json

{
   "myExampleDotComEmailServer": {
      "Type": "AWS::Route53::RecordSetGroup",
      "Properties": {
         "Comment": "Creating records for mail server",
         "HostedZoneId": "Z1PA6795UKMFR9",
         "RecordSets": [
            {
               "Name": "mail.example.com.",
               "Type": "A",
               "TTL": "900",
               "ResourceRecords": [
                  "192.0.2.44"
               ]
            },
            {
               "Name": "mail.example.com.",
               "Type": "MX",
               "TTL": "900",
               "ResourceRecords": [
                  "10 mail.example.com"
               ]
            },
            {
               "Name": "mail.example.com.",
               "Type": "TXT",
               "TTL": "900",
               "ResourceRecords": [
                  "\"v=spf1 ip4:203.0.113.0/30 -all\""
               ]
            }
         ]
      }
   }
}
```

#### YAML

```yaml

myExampleDotComEmailServer:
  Type: AWS::Route53::RecordSetGroup
  Properties:
    Comment: Creating records for mail server
    HostedZoneId: Z1PA6795UKMFR9
    RecordSets:
    - Name: mail.example.com.
      ResourceRecords:
      - 192.0.2.44
      TTL: '900'
      Type: A
    - Name: mail.example.com.
      ResourceRecords:
      - '10 mail.example.com'
      TTL: '900'
      Type: MX
    - Name: mail.example.com.
      ResourceRecords:
      - '"v=spf1 ip4:203.0.113.0/30 -all"'
      TTL: '900'
      Type: TXT
```

## See also

- For `AWS::Route53::RecordSetGroup` examples, see [ChangeResourceRecordSets](../../../../reference/route53/latest/apireference/api-changeresourcerecordsets.md) in the _Amazon Route 53 API Reference_

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GeoProximityLocation

AliasTarget
