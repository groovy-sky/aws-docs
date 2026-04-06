# Import certificates into AWS Certificate Manager

In addition to requesting SSL/TLS certificates provided by AWS Certificate Manager (ACM), you can
import certificates that you obtained outside of AWS. You might do this because you
already have a certificate from a third-party certificate authority (CA), or because you
have application-specific requirements that are not met by ACM issued certificates.

You can use an imported certificate with any [AWS service\
that is integrated with ACM](acm-services.md). The certificates that you import work the same as
those provided by ACM, with one important exception: ACM does not provide [managed renewal](https://docs.aws.amazon.com/acm/latest/userguide/managed-renewal.html) for imported certificates.

To renew an imported certificate, you can obtain a new certificate from your certificate
issuer and then manually [reimport](https://docs.aws.amazon.com/acm/latest/userguide/import-reimport.html#reimport-certificate-api) it into
ACM. This action preserves the certificate's association and its Amazon Resource name
(ARN). Alternatively, you can import a completely new certificate. Multiple certificates
with the same domain name can be imported, but they must be imported one at a time.

###### Important

You are responsible for monitoring the expiration date of your imported certificates
and for renewing them before they expire. You can simplify this task by using Amazon CloudWatch Events
to send notices when your imported certificates approach expiration. For more
information, see [Using Amazon EventBridge](https://docs.aws.amazon.com/acm/latest/userguide/cloudwatch-events.html).

All certificates in ACM are regional resources, including the certificates that you
import. To use the same certificate with Elastic Load Balancing load balancers in different AWS
Regions, you must import the certificate into each Region where you want to use it. To use a
certificate with Amazon CloudFront, you must import it into the US East (N. Virginia) Region. For more
information, see [Supported Regions](acm-overview.md#acm-regions).

For information about how to import certificates into ACM, see the following topics. If
you encounter problems importing a certificate, see [Certificate import problems](https://docs.aws.amazon.com/acm/latest/userguide/troubleshoot-import.html).

###### Topics

- [Prerequisites for importing ACM certificates](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate-prerequisites.html)

- [Certificate and key format for importing](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate-format.html)

- [Import a certificate](https://docs.aws.amazon.com/acm/latest/userguide/import-certificate-api-cli.html)

- [Reimport a certificate](https://docs.aws.amazon.com/acm/latest/userguide/import-reimport.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Export certificate

Prerequisites
