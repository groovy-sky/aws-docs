# Oracle native network encryption

Amazon RDS supports Oracle native network encryption (NNE). With the
`NATIVE_NETWORK_ENCRYPTION` option, you can encrypt data as it moves to and
from a DB instance. Amazon RDS supports NNE for all editions of Oracle Database.

A detailed discussion of Oracle native network encryption is beyond the scope of this guide, but you should
understand the strengths and weaknesses of each algorithm and key before you decide on a solution for your
deployment. For information about the algorithms and keys that are available through Oracle native network
encryption, see [Configuring network data encryption](http://www.oracle.com/webfolder/technetwork/tutorials/obe/db/11g/r2/prod/security/network_encrypt/ntwrkencrypt.htm) in the Oracle documentation. For more information about AWS
security, see the [AWS security center](http://aws.amazon.com/security).

###### Note

You can use Native Network Encryption
or Secure Sockets Layer,
but not both.
For more information, see
[Oracle Secure Sockets Layer](appendix-oracle-options-ssl.md).

###### Topics

- [NATIVE\_NETWORK\_ENCRYPTION option settings](oracle-options-nne-options.md)

- [Adding the NATIVE\_NETWORK\_ENCRYPTION option](oracle-options-nne-add.md)

- [Setting NNE values in the sqlnet.ora](oracle-options-nne-using.md)

- [Modifying NATIVE\_NETWORK\_ENCRYPTION option settings](oracle-options-nne-modifysettings.md)

- [Removing the NATIVE\_NETWORK\_ENCRYPTION option](oracle-options-nne-remove.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Locator

NATIVE\_NETWORK\_ENCRYPTION
settings

All content copied from https://docs.aws.amazon.com/.
