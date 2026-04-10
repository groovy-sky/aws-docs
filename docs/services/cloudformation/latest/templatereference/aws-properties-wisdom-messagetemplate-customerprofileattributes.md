This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate CustomerProfileAttributes

The customer profile attributes that are used with the message template.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountNumber" : String,
  "AdditionalInformation" : String,
  "Address1" : String,
  "Address2" : String,
  "Address3" : String,
  "Address4" : String,
  "BillingAddress1" : String,
  "BillingAddress2" : String,
  "BillingAddress3" : String,
  "BillingAddress4" : String,
  "BillingCity" : String,
  "BillingCountry" : String,
  "BillingCounty" : String,
  "BillingPostalCode" : String,
  "BillingProvince" : String,
  "BillingState" : String,
  "BirthDate" : String,
  "BusinessEmailAddress" : String,
  "BusinessName" : String,
  "BusinessPhoneNumber" : String,
  "City" : String,
  "Country" : String,
  "County" : String,
  "Custom" : {Key: Value, ...},
  "EmailAddress" : String,
  "FirstName" : String,
  "Gender" : String,
  "HomePhoneNumber" : String,
  "LastName" : String,
  "MailingAddress1" : String,
  "MailingAddress2" : String,
  "MailingAddress3" : String,
  "MailingAddress4" : String,
  "MailingCity" : String,
  "MailingCountry" : String,
  "MailingCounty" : String,
  "MailingPostalCode" : String,
  "MailingProvince" : String,
  "MailingState" : String,
  "MiddleName" : String,
  "MobilePhoneNumber" : String,
  "PartyType" : String,
  "PhoneNumber" : String,
  "PostalCode" : String,
  "ProfileARN" : String,
  "ProfileId" : String,
  "Province" : String,
  "ShippingAddress1" : String,
  "ShippingAddress2" : String,
  "ShippingAddress3" : String,
  "ShippingAddress4" : String,
  "ShippingCity" : String,
  "ShippingCountry" : String,
  "ShippingCounty" : String,
  "ShippingPostalCode" : String,
  "ShippingProvince" : String,
  "ShippingState" : String,
  "State" : String
}

```

### YAML

```yaml

  AccountNumber: String
  AdditionalInformation: String
  Address1: String
  Address2: String
  Address3: String
  Address4: String
  BillingAddress1: String
  BillingAddress2: String
  BillingAddress3: String
  BillingAddress4: String
  BillingCity: String
  BillingCountry: String
  BillingCounty: String
  BillingPostalCode: String
  BillingProvince: String
  BillingState: String
  BirthDate: String
  BusinessEmailAddress: String
  BusinessName: String
  BusinessPhoneNumber: String
  City: String
  Country: String
  County: String
  Custom:
    Key: Value
  EmailAddress: String
  FirstName: String
  Gender: String
  HomePhoneNumber: String
  LastName: String
  MailingAddress1: String
  MailingAddress2: String
  MailingAddress3: String
  MailingAddress4: String
  MailingCity: String
  MailingCountry: String
  MailingCounty: String
  MailingPostalCode: String
  MailingProvince: String
  MailingState: String
  MiddleName: String
  MobilePhoneNumber: String
  PartyType: String
  PhoneNumber: String
  PostalCode: String
  ProfileARN: String
  ProfileId: String
  Province: String
  ShippingAddress1: String
  ShippingAddress2: String
  ShippingAddress3: String
  ShippingAddress4: String
  ShippingCity: String
  ShippingCountry: String
  ShippingCounty: String
  ShippingPostalCode: String
  ShippingProvince: String
  ShippingState: String
  State: String

```

## Properties

`AccountNumber`

A unique account number that you have given to the customer.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AdditionalInformation`

Any additional information relevant to the customer's profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Address1`

The first line of a customer address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Address2`

The second line of a customer address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Address3`

The third line of a customer address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Address4`

The fourth line of a customer address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingAddress1`

The first line of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingAddress2`

The second line of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingAddress3`

The third line of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingAddress4`

The fourth line of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingCity`

The city of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingCountry`

The country of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingCounty`

The county of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingPostalCode`

The postal code of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingProvince`

The province of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BillingState`

The state of a customer’s billing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BirthDate`

The customer's birth date.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessEmailAddress`

The customer's business email address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessName`

The name of the customer's business.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BusinessPhoneNumber`

The customer's business phone number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`City`

The city in which a customer lives.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Country`

The country in which a customer lives.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`County`

The county in which a customer lives.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Custom`

The custom attributes in customer profile attributes.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EmailAddress`

The customer's email address, which has not been specified as a personal or business address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FirstName`

The customer's first name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Gender`

The customer's gender.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`HomePhoneNumber`

The customer's mobile phone number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LastName`

The customer's last name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingAddress1`

The first line of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingAddress2`

The second line of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingAddress3`

The third line of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingAddress4`

The fourth line of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingCity`

The city of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingCountry`

The country of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingCounty`

The county of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingPostalCode`

The postal code of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingProvince`

The province of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MailingState`

The state of a customer’s mailing address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MiddleName`

The customer's middle name.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MobilePhoneNumber`

The customer's mobile phone number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PartyType`

The customer's party type.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PhoneNumber`

The customer's phone number, which has not been specified as a mobile, home, or business number.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PostalCode`

The postal code of a customer address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileARN`

The ARN of a customer profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ProfileId`

The unique identifier of a customer profile.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Province`

The province in which a customer lives.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingAddress1`

The first line of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingAddress2`

The second line of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingAddress3`

The third line of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingAddress4`

The fourth line of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingCity`

The city of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingCountry`

The country of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingCounty`

The county of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingPostalCode`

The postal code of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingProvince`

The province of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ShippingState`

The state of a customer’s shipping address.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`State`

The state in which a customer lives.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `32767`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Content

EmailMessageTemplateContent

All content copied from https://docs.aws.amazon.com/.
