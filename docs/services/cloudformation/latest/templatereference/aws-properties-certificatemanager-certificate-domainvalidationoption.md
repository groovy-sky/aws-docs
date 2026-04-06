This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CertificateManager::Certificate DomainValidationOption

`DomainValidationOption` is a property of the [AWS::CertificateManager::Certificate](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html) resource that specifies the AWS Certificate Manager (ACM) certificate domain to validate. Depending on
the chosen validation method, ACM checks the domain's DNS record for a validation CNAME,
or it attempts to send a validation email message to the domain owner.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DomainName" : String,
  "HostedZoneId" : String,
  "ValidationDomain" : String
}

```

### YAML

```yaml

  DomainName: String
  HostedZoneId: String
  ValidationDomain: String

```

## Properties

`DomainName`

A fully qualified domain name (FQDN) in the certificate request.

_Required_: Yes

_Type_: String

_Pattern_: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HostedZoneId`

The `HostedZoneId` option, which is available if you are using Route 53 as
your domain registrar, causes ACM to add your CNAME to the domain record. Your list of
`DomainValidationOptions` must contain one and only one of the
domain-validation options, and the `HostedZoneId` can be used only when
`DNS` is specified as your validation method.

Use the Route 53 `ListHostedZones` API to discover IDs for available hosted
zones.

This option is required for publicly trusted certificates.

###### Note

The `ListHostedZones` API returns IDs in the format
"/hostedzone/Z111111QQQQQQQ", but CloudFormation requires the IDs to be in the
format "Z111111QQQQQQQ".

When you change your `DomainValidationOptions`, a new resource is
created.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ValidationDomain`

The domain name to which you want ACM to send validation emails. This domain name is
the suffix of the email addresses that you want ACM to use. This must be the same as the
`DomainName` value or a superdomain of the `DomainName` value.
For example, if you request a certificate for `testing.example.com`, you can
specify `example.com` as this value. In that case, ACM sends domain
validation emails to the following five addresses:

- admin@example.com

- administrator@example.com

- hostmaster@example.com

- postmaster@example.com

- webmaster@example.com

_Required_: No

_Type_: String

_Pattern_: `(\*\.)?(((?!-)[A-Za-z0-9-]{0,62}[A-Za-z0-9])\.)+((?!-)[A-Za-z0-9-]{1,62}[A-Za-z0-9])`

_Minimum_: `1`

_Maximum_: `253`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::CertificateManager::Certificate

Tag
