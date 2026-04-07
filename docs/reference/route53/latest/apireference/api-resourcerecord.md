# ResourceRecord

Information specific to the resource record.

###### Note

If you're creating an alias resource record set, omit
`ResourceRecord`.

## Contents

**Value**

The current or new DNS record value, not to exceed 4,000 characters. In the case of a
`DELETE` action, if the current value does not match the actual value, an
error is returned. For descriptions about how to format `Value` for different
record types, see [Supported DNS Resource\
Record Types](../../../../services/route53/latest/developerguide/resourcerecordtypes.md) in the _Amazon Route 53 Developer_
_Guide_.

You can specify more than one value for all record types except `CNAME` and
`SOA`.

###### Note

If you're creating an alias resource record set, omit `Value`.

Type: String

Length Constraints: Maximum length of 4000.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/ResourceRecord)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/ResourceRecord)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/ResourceRecord)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

QueryLoggingConfig

ResourceRecordSet
