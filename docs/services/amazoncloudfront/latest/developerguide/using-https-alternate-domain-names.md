# Use alternate domain names and HTTPS

If you want to use your own domain name in the URLs for your files (for example,
`https://www.example.com/image.jpg`) and you want your viewers to use HTTPS,
you must complete the steps in the following topics. (If you use the default CloudFront distribution
domain name in your URLs, for example,
`https://d111111abcdef8.cloudfront.net/image.jpg`, follow the guidance in the
following topic instead: [Require HTTPS for communication between viewers and CloudFront](using-https-viewers-to-cloudfront.md).)

###### Important

When you add a certificate to your distribution, CloudFront immediately propagates the
certificate to all of its edge locations. As new edge locations become available,
CloudFront propagates the certificate to those locations, too. You can't restrict the edge
locations that CloudFront propagates the certificates to.

###### Topics

- [Choose how CloudFront serves HTTPS requests](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-https-dedicated-ip-or-sni.html)

- [Requirements for using SSL/TLS certificates with CloudFront](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-requirements.html)

- [Quotas on using SSL/TLS certificates with CloudFront (HTTPS between viewers and CloudFront only)](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-limits.html)

- [Configure alternate domain names and HTTPS](cnames-and-https-procedures.md)

- [Determine the size of the public key in an SSL/TLS RSA certificate](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-size-of-public-key.html)

- [Increase the quotas for SSL/TLS certificates](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/increasing-the-limit-for-ssl-tls-certificates.html)

- [Rotate SSL/TLS certificates](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-rotate-certificates.html)

- [Revert from a custom SSL/TLS certificate to the default CloudFront certificate](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-revert-to-cf-certificate.html)

- [Switch from a custom SSL/TLS certificate with dedicated IP addresses to SNI](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/cnames-and-https-switch-dedicated-to-sni.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Supported protocols and ciphers between CloudFront and the origin

Choose how CloudFront serves HTTPS requests
