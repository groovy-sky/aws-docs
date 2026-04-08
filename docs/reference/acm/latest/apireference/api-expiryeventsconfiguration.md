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

- [AWS SDK for C++](../../../goto/sdkforcpp/acm-2015-12-08/expiryeventsconfiguration.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/acm-2015-12-08/expiryeventsconfiguration.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/acm-2015-12-08/expiryeventsconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainValidationOption

ExtendedKeyUsage

All content copied from https://docs.aws.amazon.com/.
