# Mutual TLS authentication with CloudFront (Viewer mTLS)

Mutual TLS Authentication (Mutual Transport Layer Security Authentication — mTLS) is a security protocol that extends standard TLS authentication by requiring bidirectional certificate-based authentication, where both client and server must prove their identity before establishing a secure connection. Using mutual TLS, you can ensure that only clients presenting trusted TLS certificates gain access to your CloudFront distributions.

## How it works

In a standard TLS handshake, only the server presents a certificate to prove its identity to the client. With mutual TLS, the authentication process becomes bidirectional. When a client attempts to connect to your CloudFront distribution, CloudFront requests a client certificate during the TLS handshake. The client must present a valid X.509 certificate that CloudFront validates against your configured trust store before establishing the secure connection.

CloudFront performs this certificate validation at AWS edge locations, offloading the authentication complexity from your origin servers while maintaining CloudFront's global performance benefits. You can configure mTLS in two modes: verify mode (which requires all clients to present valid certificates) or optional mode (which validates certificates when presented but also allows connections without certificates).

## Use cases

Mutual TLS authentication with CloudFront addresses several critical security scenarios where traditional authentication methods are insufficient:

- **Device authentication with content caching** \- You can authenticate gaming consoles, IoT devices, or corporate hardware before allowing access to firmware updates, game downloads, or internal resources. Each device contains a unique certificate that proves its authenticity while benefiting from CloudFront's caching capabilities.

- **API-to-API authentication** \- You can secure machine-to-machine communication between trusted business partners, payment systems, or micro-services. Certificate-based authentication eliminates the need for shared secrets or API keys while providing strong identity verification for automated data exchanges.

###### Topics

- [Trust stores and certificate management](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/trust-stores-certificate-management.html)

- [Enable mutual TLS for CloudFront distributions](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/enable-mtls-distributions.html)

- [Associate a CloudFront Connection Function](connection-functions.md)

- [Configuring additional settings](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/configuring-additional-settings.html)

- [Viewer mTLS headers for cache policies and forwarded to origin](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/viewer-mtls-headers.html)

- [Revocation using CloudFront Connection Function and KVS](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/revocation-connection-function-kvs.html)

- [Observability using connection logs](connection-logs.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Switch from a custom SSL/TLS certificate with dedicated IP addresses to SNI

Trust stores and certificate management
