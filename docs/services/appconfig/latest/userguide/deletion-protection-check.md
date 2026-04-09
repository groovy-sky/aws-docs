# Bypassing or forcing a deletion protection check

To help you manage deletion protection, the [DeleteEnvironment](../../../../reference/appconfig/2019-10-09/apireference/api-deleteenvironment.md)
and [DeleteConfigurationProfile](../../../../reference/appconfig/2019-10-09/apireference/api-deleteconfigurationprofile.md) APIs include a parameter called
`DeletionProtectionCheck`. This parameter supports the following values:

- `BYPASS`: Instructs AWS AppConfig to bypass the deletion protection
check and delete a configuration profile even if deletion protection would have otherwise
prevented it.

- `APPLY`: Instructs the deletion protection check to run, even if deletion
protection is disabled at the account level. `APPLY` also forces the deletion
protection check to run against resources created in the past hour, which are normally
excluded from deletion protection checks.

- `ACCOUNT_DEFAULT`: The default setting, which instructs AWS AppConfig
to implement the deletion protection value specified in the
`UpdateAccountSettings` API.

###### Note

By default, `DeletionProtectionCheck` skips configuration profiles and
environments created in the past hour. The default configuration is intended to prevent
deletion protection from interferring with tests and demos that create short-lived
resources. You can override this behavior by passing
`DeletionProtectionCheck=APPLY` when calling `DeleteEnvironment` or
`DeleteConfigurationProfile`.

The following CLI walkthrough uses sample commands to illustrate how to use the
`DeletionProtectionCheck` parameter. Replace `ID` in the
following commands with the ID for your AWS AppConfig artifacts.

1. Call [GetLatestConfiguration](../../../../reference/appconfig/2019-10-09/apireference/api-appconfigdata-getlatestconfiguration.md) on a deployed configuration.

```nohighlight

aws appconfigdata get-latest-configuration --configuration-token $(aws appconfigdata start-configuration-session --application-identifier ID --environment-identifier ID --configuration-profile-identifier ID --query InitialConfigurationToken) outfile.txt
```

2. Wait 60 seconds for AWS AppConfig to register that the configuration is active.

3. Run the following command to call [DeleteEnvironment](../../../../reference/appconfig/2019-10-09/apireference/api-deleteenvironment.md) and apply deletion protection on the environment.

```nohighlight

aws appconfig delete-environment --environment-id ID --application-id ID --deletion-protection-check APPLY
```

The command should return the following error.

```nohighlight

An error occurred (BadRequestException) when calling the DeleteEnvironment operation: Environment Beta is actively being used in your application and cannot be deleted.
```

4. Run the following command to bypass deletion protection and delete the
    environment.

```nohighlight

aws appconfig delete-environment --environment-id ID --application-id ID --deletion-protection-check BYPASS
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Deletion protection

Security

All content copied from https://docs.aws.amazon.com/.
