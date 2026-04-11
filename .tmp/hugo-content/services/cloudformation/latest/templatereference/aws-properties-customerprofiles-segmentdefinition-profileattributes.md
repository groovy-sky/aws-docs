This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::SegmentDefinition ProfileAttributes

The object used to segment on attributes within the customer profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountNumber" : ProfileDimension,
  "AdditionalInformation" : ExtraLengthValueProfileDimension,
  "Address" : AddressDimension,
  "Attributes" : {Key: Value, ...},
  "BillingAddress" : AddressDimension,
  "BirthDate" : DateDimension,
  "BusinessEmailAddress" : ProfileDimension,
  "BusinessName" : ProfileDimension,
  "BusinessPhoneNumber" : ProfileDimension,
  "EmailAddress" : ProfileDimension,
  "FirstName" : ProfileDimension,
  "GenderString" : ProfileDimension,
  "HomePhoneNumber" : ProfileDimension,
  "LastName" : ProfileDimension,
  "MailingAddress" : AddressDimension,
  "MiddleName" : ProfileDimension,
  "MobilePhoneNumber" : ProfileDimension,
  "PartyTypeString" : ProfileDimension,
  "PersonalEmailAddress" : ProfileDimension,
  "PhoneNumber" : ProfileDimension,
  "ProfileType" : ProfileTypeDimension,
  "ShippingAddress" : AddressDimension
}

```

### YAML

```yaml

  AccountNumber:
    ProfileDimension
  AdditionalInformation:
    ExtraLengthValueProfileDimension
  Address:
    AddressDimension
  Attributes:
    Key: Value
  BillingAddress:
    AddressDimension
  BirthDate:
    DateDimension
  BusinessEmailAddress:
    ProfileDimension
  BusinessName:
    ProfileDimension
  BusinessPhoneNumber:
    ProfileDimension
  EmailAddress:
    ProfileDimension
  FirstName:
    ProfileDimension
  GenderString:
    ProfileDimension
  HomePhoneNumber:
    ProfileDimension
  LastName:
    ProfileDimension
  MailingAddress:
    AddressDimension
  MiddleName:
    ProfileDimension
  MobilePhoneNumber:
    ProfileDimension
  PartyTypeString:
    ProfileDimension
  PersonalEmailAddress:
    ProfileDimension
  PhoneNumber:
    ProfileDimension
  ProfileType:
    ProfileTypeDimension
  ShippingAddress:
    AddressDimension

```

## Properties

`AccountNumber`

A field to describe values to segment on within account number.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AdditionalInformation`

A field to describe values to segment on within additional information.

_Required_: No

_Type_: [ExtraLengthValueProfileDimension](aws-properties-customerprofiles-segmentdefinition-extralengthvalueprofiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Address`

A field to describe values to segment on within address.

_Required_: No

_Type_: [AddressDimension](aws-properties-customerprofiles-segmentdefinition-addressdimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Attributes`

A field to describe values to segment on within attributes.

_Required_: No

_Type_: Object of [AttributeDimension](aws-properties-customerprofiles-segmentdefinition-attributedimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BillingAddress`

A field to describe values to segment on within billing address.

_Required_: No

_Type_: [AddressDimension](aws-properties-customerprofiles-segmentdefinition-addressdimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BirthDate`

A field to describe values to segment on within birthDate.

_Required_: No

_Type_: [DateDimension](aws-properties-customerprofiles-segmentdefinition-datedimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BusinessEmailAddress`

A field to describe values to segment on within business email address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BusinessName`

A field to describe values to segment on within business name.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`BusinessPhoneNumber`

A field to describe values to segment on within business phone number.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EmailAddress`

A field to describe values to segment on within email address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FirstName`

A field to describe values to segment on within first name.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GenderString`

A field to describe values to segment on within genderString.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`HomePhoneNumber`

A field to describe values to segment on within home phone number.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LastName`

A field to describe values to segment on within last name.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MailingAddress`

A field to describe values to segment on within mailing address.

_Required_: No

_Type_: [AddressDimension](aws-properties-customerprofiles-segmentdefinition-addressdimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MiddleName`

A field to describe values to segment on within middle name.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MobilePhoneNumber`

A field to describe values to segment on within mobile phone number.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PartyTypeString`

A field to describe values to segment on within partyTypeString.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PersonalEmailAddress`

A field to describe values to segment on within personal email address.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PhoneNumber`

A field to describe values to segment on within phone number.

_Required_: No

_Type_: [ProfileDimension](aws-properties-customerprofiles-segmentdefinition-profiledimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProfileType`

The type of profile.

_Required_: No

_Type_: [ProfileTypeDimension](aws-properties-customerprofiles-segmentdefinition-profiletypedimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShippingAddress`

A field to describe values to segment on within shipping address.

_Required_: No

_Type_: [AddressDimension](aws-properties-customerprofiles-segmentdefinition-addressdimension.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Group

ProfileDimension

All content copied from https://docs.aws.amazon.com/.
