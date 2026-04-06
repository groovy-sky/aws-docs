# DNS Firewall rule groups and rules

This section describes the settings that you can configure for your DNS Firewall rule
groups and rules, to define the DNS Firewall behavior for your VPCs. It also describes how
to manage the settings for your rule groups and rules.

When you have your rule groups configured the way you want them, you use them directly
and you can share and manage them between accounts and across your organization in
AWS Organizations.

- You can associate a rule group with multiple VPCs, to provide consistent
behavior across your organization. For information, see [Managing associations between your VPC and Resolver DNS Firewall rule group](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-dns-firewall-vpc-associating-rule-group.html).

- You can share rule groups between accounts, for consistent DNS query
management across your organization. For information, see [Sharing Resolver DNS Firewall rule groups between AWS accounts](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resolver-dns-firewall-rule-group-sharing.html).

- You can use rule groups across your organization in AWS Organizations by managing them
in AWS Firewall Manager policies. For information about Firewall Manager, see [AWS Firewall Manager](https://docs.aws.amazon.com/waf/latest/developerguide/fms-chapter.html) in the _AWS WAF, AWS Firewall Manager, and AWS Shield Advanced_
_Developer Guide_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Getting started with Resolver DNS Firewall

Rule group settings in DNS Firewall
