# Renew ACM public certificates

When issuing a managed, publicly trusted certificate, AWS Certificate Manager requires you to prove
that you are the domain owner. This happens by means of either [DNS validation](dns-validation.md) or [email validation](email-validation.md). When a certificate comes up for renewal, ACM uses the
same method that you chose earlier to re-validate your ownership. The following topics
describe how the renewal process works in each case.

###### Topics

- [Renewal for domains validated by DNS](https://docs.aws.amazon.com/acm/latest/userguide/dns-renewal-validation.html)

- [Renewal for email-validated domains](https://docs.aws.amazon.com/acm/latest/userguide/email-renewal-validation.html)

- [Renewal for domains validated by HTTP](https://docs.aws.amazon.com/acm/latest/userguide/http-renewal-validation.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Managed certificate renewal

DNS-validated domains
