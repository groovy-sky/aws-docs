# ExpiryEventsConfiguration

Object containing expiration events options associated with an AWS account.

## Contents

###### Note

In the following list, the required parameters are described first.

**DaysBeforeExpiry**

Specifies the number of days prior to certificate expiration when ACM starts
generating `EventBridge` events. ACM sends one event per day per
certificate until the certificate expires. By default, accounts receive events starting
45 days before certificate expiration.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/acm-2015-12-08/ExpiryEventsConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/acm-2015-12-08/ExpiryEventsConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/acm-2015-12-08/ExpiryEventsConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DomainValidationOption

ExtendedKeyUsage
