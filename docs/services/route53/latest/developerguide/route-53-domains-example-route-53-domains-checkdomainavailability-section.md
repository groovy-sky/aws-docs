# Use `CheckDomainAvailability` with an AWS SDK or CLI

The following code examples show how to use `CheckDomainAvailability`.

Action examples are code excerpts from larger programs and must be run in context. You can see this action in
context in the following code example:

- [Learn the basics](route-53-domains-example-route-53-domains-scenario-getstartedroute53domains-section.md)

.NET

**SDK for .NET**

###### Note

There's more on GitHub. Find the complete example and learn how to set up and run in the
[AWS Code\
Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/dotnetv3/Route53).

```csharp

    /// <summary>
    /// Check the availability of a domain name.
    /// </summary>
    /// <param name="domain">The domain to check for availability.</param>
    /// <returns>An availability result string.</returns>
    public async Task<string> CheckDomainAvailability(string domain)
    {
        var result = await _amazonRoute53Domains.CheckDomainAvailabilityAsync(
            new CheckDomainAvailabilityRequest
            {
                DomainName = domain
            }
        );
        return result.Availability.Value;
    }

```

- For API details, see
[CheckDomainAvailability](https://docs.aws.amazon.com/goto/DotNetSDKV3/route53domains-2014-05-15/CheckDomainAvailability)
in _AWS SDK for .NET API Reference_.

CLI

**AWS CLI**

**To determine whether you can register a domain name with Route 53**

The following `check-domain-availability` command returns information about whether the domain name `example.com`
is available to be registered using Route 53.

This command runs only in the `us-east-1` Region. If your default region is set to `us-east-1`, you can omit the `region` parameter.

```nohighlight

aws route53domains check-domain-availability \
    --region us-east-1 \
    --domain-name example.com

```

Output:

```nohighlight

{
    "Availability": "UNAVAILABLE"
}
```

Route 53 supports a large number of top-level domains (TLDs), such as `.com` and `.jp`, but we don't support all available TLDs. If you check the availability of a domain and Route 53 doesn't support the TLD, `check-domain-availability` returns the following message.

```nohighlight

An error occurred (UnsupportedTLD) when calling the CheckDomainAvailability operation: <top-level domain> tld is not supported.
```

For a list of the TLDs that you can use when registering a domain with Route 53, see [Domains That You Can Register with Amazon Route 53](registrar-tld-list.md) in the _Amazon Route 53 Developer Guide_.
For more information about registering domains with Amazon Route 53, see [Registering a New Domain](domain-register.md) in the _Amazon Route 53 Developer Guide_.

- For API details, see
[CheckDomainAvailability](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/route53domains/check-domain-availability.html)
in _AWS CLI Command Reference_.

Java

**SDK for Java 2.x**

###### Note

There's more on GitHub. Find the complete example and learn how to set up and run in the
[AWS Code\
Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/javav2/example_code/route53).

```java

    public static void checkDomainAvailability(Route53DomainsClient route53DomainsClient, String domainSuggestion) {
        try {
            CheckDomainAvailabilityRequest availabilityRequest = CheckDomainAvailabilityRequest.builder()
                    .domainName(domainSuggestion)
                    .build();

            CheckDomainAvailabilityResponse response = route53DomainsClient
                    .checkDomainAvailability(availabilityRequest);
            System.out.println(domainSuggestion + " is " + response.availability().toString());

        } catch (Route53Exception e) {
            System.err.println(e.getMessage());
            System.exit(1);
        }
    }

```

- For API details, see
[CheckDomainAvailability](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53domains-2014-05-15/CheckDomainAvailability)
in _AWS SDK for Java 2.x API Reference_.

Kotlin

**SDK for Kotlin**

###### Note

There's more on GitHub. Find the complete example and learn how to set up and run in the
[AWS Code\
Examples Repository](https://github.com/awsdocs/aws-doc-sdk-examples/tree/main/kotlin/services/route53).

```kotlin

suspend fun checkDomainAvailability(domainSuggestion: String) {
    val availabilityRequest =
        CheckDomainAvailabilityRequest {
            domainName = domainSuggestion
        }
    Route53DomainsClient.fromEnvironment { region = "us-east-1" }.use { route53DomainsClient ->
        val response = route53DomainsClient.checkDomainAvailability(availabilityRequest)
        println("$domainSuggestion is ${response.availability}")
    }
}

```

- For API details, see
[CheckDomainAvailability](https://sdk.amazonaws.com/kotlin/api/latest/index.html)
in _AWS SDK for Kotlin API reference_.

For a complete list of AWS SDK developer guides and code examples, see
[Using Route 53 with an AWS SDK](../../../../reference/route53/latest/developerguide/sdk-general-information-section.md).
This topic also includes information about getting started and details about previous SDK versions.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Actions

CheckDomainTransferability
