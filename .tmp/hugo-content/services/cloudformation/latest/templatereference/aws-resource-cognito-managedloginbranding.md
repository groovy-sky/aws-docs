This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::ManagedLoginBranding

Creates a new set of branding settings for a user pool style and associates it with an
app client. This operation is the programmatic option for the creation of a new style in
the branding designer.

Provides values for UI customization in a `Settings` JSON object and image
files in an `Assets` array. To send the JSON object `Document`
type parameter in `Settings`, you might need to update to the most recent
version of your AWS SDK.

This operation has a 2-megabyte request-size limit and include the CSS settings and
image assets for your app client. Your branding settings might exceed 2MB in size.
Amazon Cognito doesn't require that you pass all parameters in one request and preserves
existing style settings that you don't specify. If your request is larger than 2MB,
separate it into multiple requests, each with a size smaller than the limit.

As a best practice, modify the output of [DescribeManagedLoginBrandingByClient](../../../../reference/cognito-user-identity-pools/latest/apireference/api-describemanagedloginbrandingbyclient.md) into the request parameters for this
operation. To get all settings, set `ReturnMergedResources` to
`true`. For more information, see [API and SDK operations for managed login branding](../../../cognito/latest/developerguide/managed-login-brandingdesigner.md#branding-designer-api)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Cognito::ManagedLoginBranding",
  "Properties" : {
      "Assets" : [ AssetType, ... ],
      "ClientId" : String,
      "ReturnMergedResources" : Boolean,
      "Settings" : Json,
      "UseCognitoProvidedValues" : Boolean,
      "UserPoolId" : String
    }
}

```

### YAML

```yaml

Type: AWS::Cognito::ManagedLoginBranding
Properties:
  Assets:
    - AssetType
  ClientId: String
  ReturnMergedResources: Boolean
  Settings: Json
  UseCognitoProvidedValues: Boolean
  UserPoolId: String

```

## Properties

`Assets`

An array of image files that you want to apply to roles like backgrounds, logos, and
icons. Each object must also indicate whether it is for dark mode, light mode, or
browser-adaptive mode.

_Required_: No

_Type_: Array of [AssetType](aws-properties-cognito-managedloginbranding-assettype.md)

_Minimum_: `0`

_Maximum_: `40`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ClientId`

The app client that you want to assign the branding style to. Each style is linked to
an app client until you delete it.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ReturnMergedResources`

When `true`, returns values for branding options that are unchanged from
Amazon Cognito defaults. When `false` or when you omit this parameter, returns only
values that you customized in your branding style.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Settings`

A JSON file, encoded as a `Document` type, with the the settings that you
want to apply to your style.

The following components are not currently implemented and reserved for future
use:

- `signUp`

- `instructions`

- `sessionTimerDisplay`

- `languageSelector` (for localization, see [Managed login localization)](../../../cognito/latest/developerguide/cognito-user-pools-managed-login.md#managed-login-localization)

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseCognitoProvidedValues`

When true, applies the default branding style options. This option reverts to default
style options that are managed by Amazon Cognito. You can modify them later in the branding
editor.

When you specify `true` for this option, you must also omit values for
`Settings` and `Assets` in the request.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserPoolId`

The user pool where the branding style is assigned.

_Required_: Yes

_Type_: String

_Pattern_: `[\w-]+_[0-9a-zA-Z]+`

_Minimum_: `1`

_Maximum_: `55`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the style ID, for example
`a1b2c3d4-5678-90ab-cdef-EXAMPLE22222`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ManagedLoginBrandingId`

The ID of the managed login branding style.

## Examples

### Creating a new branding style for an app client

The following example creates a branding style for the requested app
client.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Description": "Provision a managed login branding style to a user pool app client\n",
    "Resources": {
        "ManagedLoginBranding": {
            "Properties": {
                "Assets": [
                    {
                        "Bytes": "PHN2ZyB3aWR0aD0iMjAwMDAiIGhlaWdodD0iNDAwIiB2aWV3Qm94PSIwIDAgMjAwMDAgNDAwIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgo8ZyBjbGlwLXBhdGg9InVybCgjY2xpcDBfMTcyNTlfMjM2Njc0KSI+CjxyZWN0IHdpZHRoPSIyMDAwMCIgaGVpZ2h0PSI0MDAiIGZpbGw9InVybCgjcGFpbnQwX2xpbmVhcl8xNzI1OV8yMzY2NzQpIi8+CjxwYXRoIGQ9Ik0wIDBIMjAwMDBWNDAwSDBWMFoiIGZpbGw9IiMxMjIwMzciIGZpbGwtb3BhY2l0eT0iMC41Ii8+CjwvZz4KPGRlZnM+CjxsaW5lYXJHcmFkaWVudCBpZD0icGFpbnQwX2xpbmVhcl8xNzI1OV8yMzY2NzQiIHgxPSItODk0LjI0OSIgeTE9IjE5OS45MzEiIHgyPSIxODAzNC41IiB5Mj0iLTU4OTkuNTciIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj4KPHN0b3Agc3RvcC1jb2xvcj0iI0JGODBGRiIvPgo8c3RvcCBvZmZzZXQ9IjEiIHN0b3AtY29sb3I9IiNGRjhGQUIiLz4KPC9saW5lYXJHcmFkaWVudD4KPGNsaXBQYXRoIGlkPSJjbGlwMF8xNzI1OV8yMzY2NzQiPgo8cmVjdCB3aWR0aD0iMjAwMDAiIGhlaWdodD0iNDAwIiBmaWxsPSJ3aGl0ZSIvPgo8L2NsaXBQYXRoPgo8L2RlZnM+Cjwvc3ZnPgo=",
                        "Category": "PAGE_FOOTER_BACKGROUND",
                        "ColorMode": "DARK",
                        "Extension": "SVG"
                    }
                ],
                "ClientId": "1example23456789",
                "ReturnMergedResources": false,
                "Settings": {
                    "categories": {
                        "auth": {
                            "authMethodOrder": [
                                [
                                    {
                                        "display": "BUTTON",
                                        "type": "FEDERATED"
                                    },
                                    {
                                        "display": "INPUT",
                                        "type": "USERNAME_PASSWORD"
                                    }
                                ]
                            ],
                            "federation": {
                                "interfaceStyle": "BUTTON_LIST",
                                "order": [
                                ]
                            }
                        },
                        "form": {
                            "displayGraphics": true,
                            "instructions": {
                                "enabled": false
                            },
                            "languageSelector": {
                                "enabled": false
                            },
                            "location": {
                                "horizontal": "CENTER",
                                "vertical": "CENTER"
                            },
                            "sessionTimerDisplay": "NONE"
                        },
                        "global": {
                            "colorSchemeMode": "LIGHT",
                            "pageFooter": {
                                "enabled": false
                            },
                            "pageHeader": {
                                "enabled": false
                            },
                            "spacingDensity": "REGULAR"
                        },
                        "signUp": {
                            "acceptanceElements": [
                                {
                                    "enforcement": "NONE",
                                    "textKey": "en"
                                }
                            ]
                        }
                    },
                    "componentClasses": {
                        "buttons": {
                            "borderRadius": 8.0
                        },
                        "divider": {
                            "darkMode": {
                                "borderColor": "232b37ff"
                            },
                            "lightMode": {
                                "borderColor": "ebebf0ff"
                            }
                        },
                        "dropDown": {
                            "borderRadius": 8.0,
                            "darkMode": {
                                "defaults": {
                                    "itemBackgroundColor": "192534ff"
                                },
                                "hover": {
                                    "itemBackgroundColor": "081120ff",
                                    "itemBorderColor": "5f6b7aff",
                                    "itemTextColor": "e9ebedff"
                                },
                                "match": {
                                    "itemBackgroundColor": "d1d5dbff",
                                    "itemTextColor": "89bdeeff"
                                }
                            },
                            "lightMode": {
                                "defaults": {
                                    "itemBackgroundColor": "ffffffff"
                                },
                                "hover": {
                                    "itemBackgroundColor": "f4f4f4ff",
                                    "itemBorderColor": "7d8998ff",
                                    "itemTextColor": "000716ff"
                                },
                                "match": {
                                    "itemBackgroundColor": "414d5cff",
                                    "itemTextColor": "0972d3ff"
                                }
                            }
                        },
                        "focusState": {
                            "darkMode": {
                                "borderColor": "539fe5ff"
                            },
                            "lightMode": {
                                "borderColor": "0972d3ff"
                            }
                        },
                        "idpButtons": {
                            "icons": {
                                "enabled": true
                            }
                        },
                        "input": {
                            "borderRadius": 8.0,
                            "darkMode": {
                                "defaults": {
                                    "backgroundColor": "0f1b2aff",
                                    "borderColor": "5f6b7aff"
                                },
                                "placeholderColor": "8d99a8ff"
                            },
                            "lightMode": {
                                "defaults": {
                                    "backgroundColor": "ffffffff",
                                    "borderColor": "7d8998ff"
                                },
                                "placeholderColor": "5f6b7aff"
                            }
                        },
                        "inputDescription": {
                            "darkMode": {
                                "textColor": "8d99a8ff"
                            },
                            "lightMode": {
                                "textColor": "5f6b7aff"
                            }
                        },
                        "inputLabel": {
                            "darkMode": {
                                "textColor": "d1d5dbff"
                            },
                            "lightMode": {
                                "textColor": "000716ff"
                            }
                        },
                        "link": {
                            "darkMode": {
                                "defaults": {
                                    "textColor": "539fe5ff"
                                },
                                "hover": {
                                    "textColor": "89bdeeff"
                                }
                            },
                            "lightMode": {
                                "defaults": {
                                    "textColor": "0972d3ff"
                                },
                                "hover": {
                                    "textColor": "033160ff"
                                }
                            }
                        },
                        "optionControls": {
                            "darkMode": {
                                "defaults": {
                                    "backgroundColor": "0f1b2aff",
                                    "borderColor": "7d8998ff"
                                },
                                "selected": {
                                    "backgroundColor": "539fe5ff",
                                    "foregroundColor": "000716ff"
                                }
                            },
                            "lightMode": {
                                "defaults": {
                                    "backgroundColor": "ffffffff",
                                    "borderColor": "7d8998ff"
                                },
                                "selected": {
                                    "backgroundColor": "0972d3ff",
                                    "foregroundColor": "ffffffff"
                                }
                            }
                        },
                        "statusIndicator": {
                            "darkMode": {
                                "error": {
                                    "backgroundColor": "1a0000ff",
                                    "borderColor": "eb6f6fff",
                                    "indicatorColor": "eb6f6fff"
                                },
                                "pending": {
                                    "indicatorColor": "AAAAAAAA"
                                },
                                "success": {
                                    "backgroundColor": "001a02ff",
                                    "borderColor": "29ad32ff",
                                    "indicatorColor": "29ad32ff"
                                },
                                "warning": {
                                    "backgroundColor": "1d1906ff",
                                    "borderColor": "e0ca57ff",
                                    "indicatorColor": "e0ca57ff"
                                }
                            },
                            "lightMode": {
                                "error": {
                                    "backgroundColor": "fff7f7ff",
                                    "borderColor": "d91515ff",
                                    "indicatorColor": "d91515ff"
                                },
                                "pending": {
                                    "indicatorColor": "AAAAAAAA"
                                },
                                "success": {
                                    "backgroundColor": "f2fcf3ff",
                                    "borderColor": "037f0cff",
                                    "indicatorColor": "037f0cff"
                                },
                                "warning": {
                                    "backgroundColor": "fffce9ff",
                                    "borderColor": "8d6605ff",
                                    "indicatorColor": "8d6605ff"
                                }
                            }
                        }
                    },
                    "components": {
                        "alert": {
                            "borderRadius": 12.0,
                            "darkMode": {
                                "error": {
                                    "backgroundColor": "1a0000ff",
                                    "borderColor": "eb6f6fff"
                                }
                            },
                            "lightMode": {
                                "error": {
                                    "backgroundColor": "fff7f7ff",
                                    "borderColor": "d91515ff"
                                }
                            }
                        },
                        "favicon": {
                            "enabledTypes": [
                                "ICO",
                                "SVG"
                            ]
                        },
                        "form": {
                            "backgroundImage": {
                                "enabled": false
                            },
                            "borderRadius": 8.0,
                            "darkMode": {
                                "backgroundColor": "0f1b2aff",
                                "borderColor": "424650ff"
                            },
                            "lightMode": {
                                "backgroundColor": "ffffffff",
                                "borderColor": "c6c6cdff"
                            },
                            "logo": {
                                "enabled": false,
                                "formInclusion": "IN",
                                "location": "CENTER",
                                "position": "TOP"
                            }
                        },
                        "idpButton": {
                            "custom": {
                            },
                            "standard": {
                                "darkMode": {
                                    "active": {
                                        "backgroundColor": "354150ff",
                                        "borderColor": "89bdeeff",
                                        "textColor": "89bdeeff"
                                    },
                                    "defaults": {
                                        "backgroundColor": "0f1b2aff",
                                        "borderColor": "c6c6cdff",
                                        "textColor": "c6c6cdff"
                                    },
                                    "hover": {
                                        "backgroundColor": "192534ff",
                                        "borderColor": "89bdeeff",
                                        "textColor": "89bdeeff"
                                    }
                                },
                                "lightMode": {
                                    "active": {
                                        "backgroundColor": "d3e7f9ff",
                                        "borderColor": "033160ff",
                                        "textColor": "033160ff"
                                    },
                                    "defaults": {
                                        "backgroundColor": "ffffffff",
                                        "borderColor": "424650ff",
                                        "textColor": "424650ff"
                                    },
                                    "hover": {
                                        "backgroundColor": "f2f8fdff",
                                        "borderColor": "033160ff",
                                        "textColor": "033160ff"
                                    }
                                }
                            }
                        },
                        "pageBackground": {
                            "darkMode": {
                                "color": "0f1b2aff"
                            },
                            "image": {
                                "enabled": true
                            },
                            "lightMode": {
                                "color": "ffffffff"
                            }
                        },
                        "pageFooter": {
                            "backgroundImage": {
                                "enabled": false
                            },
                            "darkMode": {
                                "background": {
                                    "color": "0f141aff"
                                },
                                "borderColor": "424650ff"
                            },
                            "lightMode": {
                                "background": {
                                    "color": "fafafaff"
                                },
                                "borderColor": "d5dbdbff"
                            },
                            "logo": {
                                "enabled": false,
                                "location": "START"
                            }
                        },
                        "pageHeader": {
                            "backgroundImage": {
                                "enabled": false
                            },
                            "darkMode": {
                                "background": {
                                    "color": "0f141aff"
                                },
                                "borderColor": "424650ff"
                            },
                            "lightMode": {
                                "background": {
                                    "color": "fafafaff"
                                },
                                "borderColor": "d5dbdbff"
                            },
                            "logo": {
                                "enabled": false,
                                "location": "START"
                            }
                        },
                        "pageText": {
                            "darkMode": {
                                "bodyColor": "b6bec9ff",
                                "descriptionColor": "b6bec9ff",
                                "headingColor": "d1d5dbff"
                            },
                            "lightMode": {
                                "bodyColor": "414d5cff",
                                "descriptionColor": "414d5cff",
                                "headingColor": "000716ff"
                            }
                        },
                        "phoneNumberSelector": {
                            "displayType": "TEXT"
                        },
                        "primaryButton": {
                            "darkMode": {
                                "active": {
                                    "backgroundColor": "539fe5ff",
                                    "textColor": "000716ff"
                                },
                                "defaults": {
                                    "backgroundColor": "539fe5ff",
                                    "textColor": "000716ff"
                                },
                                "disabled": {
                                    "backgroundColor": "ffffffff",
                                    "borderColor": "ffffffff"
                                },
                                "hover": {
                                    "backgroundColor": "89bdeeff",
                                    "textColor": "000716ff"
                                }
                            },
                            "lightMode": {
                                "active": {
                                    "backgroundColor": "033160ff",
                                    "textColor": "ffffffff"
                                },
                                "defaults": {
                                    "backgroundColor": "0972d3ff",
                                    "textColor": "ffffffff"
                                },
                                "disabled": {
                                    "backgroundColor": "ffffffff",
                                    "borderColor": "ffffffff"
                                },
                                "hover": {
                                    "backgroundColor": "033160ff",
                                    "textColor": "ffffffff"
                                }
                            }
                        },
                        "secondaryButton": {
                            "darkMode": {
                                "active": {
                                    "backgroundColor": "354150ff",
                                    "borderColor": "89bdeeff",
                                    "textColor": "89bdeeff"
                                },
                                "defaults": {
                                    "backgroundColor": "0f1b2aff",
                                    "borderColor": "539fe5ff",
                                    "textColor": "539fe5ff"
                                },
                                "hover": {
                                    "backgroundColor": "192534ff",
                                    "borderColor": "89bdeeff",
                                    "textColor": "89bdeeff"
                                }
                            },
                            "lightMode": {
                                "active": {
                                    "backgroundColor": "d3e7f9ff",
                                    "borderColor": "033160ff",
                                    "textColor": "033160ff"
                                },
                                "defaults": {
                                    "backgroundColor": "ffffffff",
                                    "borderColor": "0972d3ff",
                                    "textColor": "0972d3ff"
                                },
                                "hover": {
                                    "backgroundColor": "f2f8fdff",
                                    "borderColor": "033160ff",
                                    "textColor": "033160ff"
                                }
                            }
                        }
                    }
                },
                "UseCognitoProvidedValues": false,
                "UserPoolId": "ca-central-1_EXAMPLE"
            },
            "Type": "AWS::Cognito::ManagedLoginBranding"
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"

Description: |
  Provision a managed login branding style to a user pool app client

Resources:
  ManagedLoginBranding:
    Type: AWS::Cognito::ManagedLoginBranding
    Properties:
      Assets:
        - Bytes: PHN2ZyB3aWR0aD0iMjAwMDAiIGhlaWdodD0iNDAwIiB2aWV3Qm94PSIwIDAgMjAwMDAgNDAwIiBmaWxsPSJub25lIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciPgo8ZyBjbGlwLXBhdGg9InVybCgjY2xpcDBfMTcyNTlfMjM2Njc0KSI+CjxyZWN0IHdpZHRoPSIyMDAwMCIgaGVpZ2h0PSI0MDAiIGZpbGw9InVybCgjcGFpbnQwX2xpbmVhcl8xNzI1OV8yMzY2NzQpIi8+CjxwYXRoIGQ9Ik0wIDBIMjAwMDBWNDAwSDBWMFoiIGZpbGw9IiMxMjIwMzciIGZpbGwtb3BhY2l0eT0iMC41Ii8+CjwvZz4KPGRlZnM+CjxsaW5lYXJHcmFkaWVudCBpZD0icGFpbnQwX2xpbmVhcl8xNzI1OV8yMzY2NzQiIHgxPSItODk0LjI0OSIgeTE9IjE5OS45MzEiIHgyPSIxODAzNC41IiB5Mj0iLTU4OTkuNTciIGdyYWRpZW50VW5pdHM9InVzZXJTcGFjZU9uVXNlIj4KPHN0b3Agc3RvcC1jb2xvcj0iI0JGODBGRiIvPgo8c3RvcCBvZmZzZXQ9IjEiIHN0b3AtY29sb3I9IiNGRjhGQUIiLz4KPC9saW5lYXJHcmFkaWVudD4KPGNsaXBQYXRoIGlkPSJjbGlwMF8xNzI1OV8yMzY2NzQiPgo8cmVjdCB3aWR0aD0iMjAwMDAiIGhlaWdodD0iNDAwIiBmaWxsPSJ3aGl0ZSIvPgo8L2NsaXBQYXRoPgo8L2RlZnM+Cjwvc3ZnPgo=
          Category: PAGE_FOOTER_BACKGROUND
          ColorMode: DARK
          Extension: SVG
      ClientId: 1example23456789
      ReturnMergedResources: false
      Settings:
        categories:
          auth:
            authMethodOrder:
              - - display: BUTTON
                  type: FEDERATED
                - display: INPUT
                  type: USERNAME_PASSWORD
            federation:
              interfaceStyle: BUTTON_LIST
              order: []
          form:
            displayGraphics: true
            instructions:
              enabled: false
            languageSelector:
              enabled: false
            location:
              horizontal: CENTER
              vertical: CENTER
            sessionTimerDisplay: NONE
          global:
            colorSchemeMode: LIGHT
            pageFooter:
              enabled: false
            pageHeader:
              enabled: false
            spacingDensity: REGULAR
          signUp:
            acceptanceElements:
              - enforcement: NONE
                textKey: en
        componentClasses:
          buttons:
            borderRadius: 8.0
          divider:
            darkMode:
              borderColor: 232b37ff
            lightMode:
              borderColor: ebebf0ff
          dropDown:
            borderRadius: 8.0
            darkMode:
              defaults:
                itemBackgroundColor: 192534ff
              hover:
                itemBackgroundColor: 081120ff
                itemBorderColor: 5f6b7aff
                itemTextColor: e9ebedff
              match:
                itemBackgroundColor: d1d5dbff
                itemTextColor: 89bdeeff
            lightMode:
              defaults:
                itemBackgroundColor: ffffffff
              hover:
                itemBackgroundColor: f4f4f4ff
                itemBorderColor: 7d8998ff
                itemTextColor: 000716ff
              match:
                itemBackgroundColor: 414d5cff
                itemTextColor: 0972d3ff
          focusState:
            darkMode:
              borderColor: 539fe5ff
            lightMode:
              borderColor: 0972d3ff
          idpButtons:
            icons:
              enabled: true
          input:
            borderRadius: 8.0
            darkMode:
              defaults:
                backgroundColor: 0f1b2aff
                borderColor: 5f6b7aff
              placeholderColor: 8d99a8ff
            lightMode:
              defaults:
                backgroundColor: ffffffff
                borderColor: 7d8998ff
              placeholderColor: 5f6b7aff
          inputDescription:
            darkMode:
              textColor: 8d99a8ff
            lightMode:
              textColor: 5f6b7aff
          inputLabel:
            darkMode:
              textColor: d1d5dbff
            lightMode:
              textColor: 000716ff
          link:
            darkMode:
              defaults:
                textColor: 539fe5ff
              hover:
                textColor: 89bdeeff
            lightMode:
              defaults:
                textColor: 0972d3ff
              hover:
                textColor: 033160ff
          optionControls:
            darkMode:
              defaults:
                backgroundColor: 0f1b2aff
                borderColor: 7d8998ff
              selected:
                backgroundColor: 539fe5ff
                foregroundColor: 000716ff
            lightMode:
              defaults:
                backgroundColor: ffffffff
                borderColor: 7d8998ff
              selected:
                backgroundColor: 0972d3ff
                foregroundColor: ffffffff
          statusIndicator:
            darkMode:
              error:
                backgroundColor: 1a0000ff
                borderColor: eb6f6fff
                indicatorColor: eb6f6fff
              pending:
                indicatorColor: AAAAAAAA
              success:
                backgroundColor: 001a02ff
                borderColor: 29ad32ff
                indicatorColor: 29ad32ff
              warning:
                backgroundColor: 1d1906ff
                borderColor: e0ca57ff
                indicatorColor: e0ca57ff
            lightMode:
              error:
                backgroundColor: fff7f7ff
                borderColor: d91515ff
                indicatorColor: d91515ff
              pending:
                indicatorColor: AAAAAAAA
              success:
                backgroundColor: f2fcf3ff
                borderColor: 037f0cff
                indicatorColor: 037f0cff
              warning:
                backgroundColor: fffce9ff
                borderColor: 8d6605ff
                indicatorColor: 8d6605ff
        components:
          alert:
            borderRadius: 12.0
            darkMode:
              error:
                backgroundColor: 1a0000ff
                borderColor: eb6f6fff
            lightMode:
              error:
                backgroundColor: fff7f7ff
                borderColor: d91515ff
          favicon:
            enabledTypes:
              - ICO
              - SVG
          form:
            backgroundImage:
              enabled: false
            borderRadius: 8.0
            darkMode:
              backgroundColor: 0f1b2aff
              borderColor: 424650ff
            lightMode:
              backgroundColor: ffffffff
              borderColor: c6c6cdff
            logo:
              enabled: false
              formInclusion: IN
              location: CENTER
              position: TOP
          idpButton:
            custom: {}
            standard:
              darkMode:
                active:
                  backgroundColor: 354150ff
                  borderColor: 89bdeeff
                  textColor: 89bdeeff
                defaults:
                  backgroundColor: 0f1b2aff
                  borderColor: c6c6cdff
                  textColor: c6c6cdff
                hover:
                  backgroundColor: 192534ff
                  borderColor: 89bdeeff
                  textColor: 89bdeeff
              lightMode:
                active:
                  backgroundColor: d3e7f9ff
                  borderColor: 033160ff
                  textColor: 033160ff
                defaults:
                  backgroundColor: ffffffff
                  borderColor: 424650ff
                  textColor: 424650ff
                hover:
                  backgroundColor: f2f8fdff
                  borderColor: 033160ff
                  textColor: 033160ff
          pageBackground:
            darkMode:
              color: 0f1b2aff
            image:
              enabled: true
            lightMode:
              color: ffffffff
          pageFooter:
            backgroundImage:
              enabled: false
            darkMode:
              background:
                color: 0f141aff
              borderColor: 424650ff
            lightMode:
              background:
                color: fafafaff
              borderColor: d5dbdbff
            logo:
              enabled: false
              location: START
          pageHeader:
            backgroundImage:
              enabled: false
            darkMode:
              background:
                color: 0f141aff
              borderColor: 424650ff
            lightMode:
              background:
                color: fafafaff
              borderColor: d5dbdbff
            logo:
              enabled: false
              location: START
          pageText:
            darkMode:
              bodyColor: b6bec9ff
              descriptionColor: b6bec9ff
              headingColor: d1d5dbff
            lightMode:
              bodyColor: 414d5cff
              descriptionColor: 414d5cff
              headingColor: 000716ff
          phoneNumberSelector:
            displayType: TEXT
          primaryButton:
            darkMode:
              active:
                backgroundColor: 539fe5ff
                textColor: 000716ff
              defaults:
                backgroundColor: 539fe5ff
                textColor: 000716ff
              disabled:
                backgroundColor: ffffffff
                borderColor: ffffffff
              hover:
                backgroundColor: 89bdeeff
                textColor: 000716ff
            lightMode:
              active:
                backgroundColor: 033160ff
                textColor: ffffffff
              defaults:
                backgroundColor: 0972d3ff
                textColor: ffffffff
              disabled:
                backgroundColor: ffffffff
                borderColor: ffffffff
              hover:
                backgroundColor: 033160ff
                textColor: ffffffff
          secondaryButton:
            darkMode:
              active:
                backgroundColor: 354150ff
                borderColor: 89bdeeff
                textColor: 89bdeeff
              defaults:
                backgroundColor: 0f1b2aff
                borderColor: 539fe5ff
                textColor: 539fe5ff
              hover:
                backgroundColor: 192534ff
                borderColor: 89bdeeff
                textColor: 89bdeeff
            lightMode:
              active:
                backgroundColor: d3e7f9ff
                borderColor: 033160ff
                textColor: 033160ff
              defaults:
                backgroundColor: ffffffff
                borderColor: 0972d3ff
                textColor: 0972d3ff
              hover:
                backgroundColor: f2f8fdff
                borderColor: 033160ff
                textColor: 033160ff
      UseCognitoProvidedValues: false
      UserPoolId: ca-central-1_EXAMPLE
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3Configuration

AssetType

All content copied from https://docs.aws.amazon.com/.
