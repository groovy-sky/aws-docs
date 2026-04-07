# Configuring Kerberos authentication for Amazon RDS for Oracle

You can use Kerberos authentication to authenticate users when they connect to your Amazon RDS for Oracle DB instance. In this configuration, your DB
instance works with AWS Directory Service for Microsoft Active Directory, also called AWS Managed Microsoft AD. When users authenticate with an RDS for Oracle DB instance joined to the trusting
domain, authentication requests are forwarded to the directory that you create with Directory Service.

Keeping all of your credentials in the same directory can save you time and effort. You have a centralized place
for storing and managing credentials for multiple database instances. A directory can also improve your overall
security profile.

###### Topics

- [Setting up Kerberos for Oracle DB instances](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos-setting-up.html)

- [Managing a DB instance in a domain](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos-managing.html)

- [Connecting to Oracle with Kerberos authentication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/oracle-kerberos-connecting.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Encrypting with NNE

Region and version availability
