package consts

const (
	/*
		common field names
	*/
	FieldPath                     = "path"
	FieldParameters               = "parameters"
	FieldMethod                   = "method"
	FieldNamespace                = "namespace"
	FieldNamespaceID              = "namespace_id"
	FieldNamespacePath            = "namespace_path"
	FieldBackend                  = "backend"
	FieldPathFQ                   = "path_fq"
	FieldData                     = "data"
	FieldMount                    = "mount"
	FieldName                     = "name"
	FieldVersion                  = "version"
	FieldMetadata                 = "metadata"
	FieldNames                    = "names"
	FieldLeaseID                  = "lease_id"
	FieldLeaseDuration            = "lease_duration"
	FieldLeaseRenewable           = "lease_renewable"
	FieldDepth                    = "depth"
	FieldDataJSON                 = "data_json"
	FieldRole                     = "role"
	FieldDescription              = "description"
	FieldTTL                      = "ttl"
	FieldDefaultLeaseTTL          = "default_lease_ttl_seconds"
	FieldMaxLeaseTTL              = "max_lease_ttl_seconds"
	FieldAuditNonHMACRequestKeys  = "audit_non_hmac_request_keys"
	FieldAuditNonHMACResponseKeys = "audit_non_hmac_response_keys"
	FieldLocal                    = "local"
	FieldSealWrap                 = "seal_wrap"
	FieldExternalEntropyAccess    = "external_entropy_access"
	FieldAWS                      = "aws"
	FieldPKCS                     = "pkcs"
	FieldAzure                    = "azure"
	FieldLibrary                  = "library"
	FieldKeyLabel                 = "key_label"
	FieldKeyID                    = "key_id"
	FieldMechanism                = "mechanism"
	FieldPin                      = "pin"
	FieldSlot                     = "slot"
	FieldTokenLabel               = "token_label"
	FieldCurve                    = "curve"
	FieldKeyBits                  = "key_bits"
	FieldForceRWSession           = "force_rw_session"
	FieldAccessKey                = "access_key"
	FieldSecretKey                = "secret_key"
	FieldEndpoint                 = "endpoint"
	FieldKeyType                  = "key_type"
	FieldKMSKey                   = "kms_key"
	FieldRegion                   = "region"
	FieldTenantID                 = "tenant_id"
	FieldClientID                 = "client_id"
	FieldClientSecret             = "client_secret"
	FieldEnvironment              = "environment"
	FieldVaultName                = "vault_name"
	FieldKeyName                  = "key_name"
	FieldResource                 = "resource"
	FieldAllowGenerateKey         = "allow_generate_key"
	FieldAllowReplaceKey          = "allow_replace_key"
	FieldAllowStoreKey            = "allow_store_key"
	FieldAnyMount                 = "any_mount"
	FieldID                       = "id"
	FieldUUID                     = "uuid"
	FieldMountAccessor            = "mount_accessor"
	FieldUsername                 = "username"
	FieldPassword                 = "password"
	FieldPasswordFile             = "password_file"
	FieldClientAuth               = "client_auth"
	FieldAuthLoginDefault         = "auth_login"
	FieldAuthLoginUserpass        = "auth_login_userpass"
	FieldAuthLoginAWS             = "auth_login_aws"
	FieldAuthLoginCert            = "auth_login_cert"
	FieldAuthLoginGCP             = "auth_login_gcp"
	FieldAuthLoginKerberos        = "auth_login_kerberos"
	FieldAuthLoginRadius          = "auth_login_radius"
	FieldAuthLoginOCI             = "auth_login_oci"
	FieldAuthLoginOIDC            = "auth_login_oidc"
	FieldAuthLoginJWT             = "auth_login_jwt"
	FieldAuthLoginAzure           = "auth_login_azure"
	FieldIAMHttpRequestMethod     = "iam_http_request_method"
	FieldIAMRequestURL            = "iam_request_url"
	FieldIAMRequestBody           = "iam_request_body"
	FieldIAMRequestHeaders        = "iam_request_headers"
	FieldAWSAccessKeyID           = "aws_access_key_id"
	FieldAWSSecretAccessKey       = "aws_secret_access_key"
	FieldAWSSessionToken          = "aws_session_token"
	FieldAWSRoleARN               = "aws_role_arn"
	FieldAWSRoleSessionName       = "aws_role_session_name"
	FieldAWSWebIdentityTokenFile  = "aws_web_identity_token_file"
	FieldAWSSTSEndpoint           = "aws_sts_endpoint"
	FieldAWSIAMEndpoint           = "aws_iam_endpoint"
	FieldAWSProfile               = "aws_profile"
	FieldAWSRegion                = "aws_region"
	FieldAWSSharedCredentialsFile = "aws_shared_credentials_file"
	FieldHeaderValue              = "header_value"
	FieldDisableRemount           = "disable_remount"
	FieldCACertFile               = "ca_cert_file"
	FieldCACertDir                = "ca_cert_dir"
	FieldCertFile                 = "cert_file"
	FieldKeyFile                  = "key_file"
	FieldSkipTLSVerify            = "skip_tls_verify"
	FieldTLSServerName            = "tls_server_name"
	FieldAddress                  = "address"
	FieldJWT                      = "jwt"
	FieldCredentials              = "credentials"
	FieldClientEmail              = "client_email"
	FieldServiceAccount           = "service_account"
	FieldAuthorization            = "authorization"
	FieldToken                    = "token"
	FieldService                  = "service"
	FieldRealm                    = "realm"
	FieldKeytabPath               = "keytab_path"
	FieldKRB5ConfPath             = "krb5conf_path"
	FieldDisableFastNegotiation   = "disable_fast_negotiation"
	FieldRemoveInstanceName       = "remove_instance_name"
	FieldAuthType                 = "auth_type"
	FieldRequestHeaders           = "request_headers"
	FieldCallbackAddress          = "callback_address"
	FieldCallbackListenerAddress  = "callback_listener_address"
	FieldScope                    = "scope"
	FieldSubscriptionID           = "subscription_id"
	FieldResourceGroupName        = "resource_group_name"
	FieldVMName                   = "vm_name"
	FieldVMSSName                 = "vmss_name"
	FieldUsernameFormat           = "username_format"
	FieldIntegrationKey           = "integration_key"
	FieldAPIHostname              = "api_hostname"
	FieldPushInfo                 = "push_info"
	FieldUsePasscode              = "use_passcode"
	FieldIssuer                   = "issuer"
	FieldPeriod                   = "period"
	FieldKeySize                  = "key_size"
	FieldQRSize                   = "qr_size"
	FieldAlgorithm                = "algorithm"
	FieldDigits                   = "digits"
	FieldSkew                     = "skew"
	FieldMaxValidationAttempts    = "max_validation_attempts"
	FieldOrgName                  = "org_name"
	FieldAPIToken                 = "api_token"
	FieldBaseURL                  = "base_url"
	FieldPrimaryEmail             = "primary_email"
	FieldSettingsFileBase64       = "settings_file_base64"
	FieldUseSignature             = "use_signature"
	FieldIdpURL                   = "idp_url"
	FieldAdminURL                 = "admin_url"
	FieldAuthenticatorURL         = "authenticator_url"
	FieldOrgAlias                 = "org_alias"
	FieldType                     = "type"
	FieldMethodID                 = "method_id"
	FieldMFAMethodIDs             = "mfa_method_ids"
	FieldAuthMethodAccessors      = "auth_method_accessors"
	FieldAuthMethodTypes          = "auth_method_types"
	FieldIdentityGroupIDs         = "identity_group_ids"
	FieldIdentityEntityIDs        = "identity_entity_ids"
	FieldWrappingAccessor         = "wrapping_accessor"
	FieldRoleName                 = "role_name"
	FieldPolicies                 = "policies"
	FieldNoParent                 = "no_parent"
	FieldNoDefaultPolicy          = "no_default_policy"
	FieldRenewable                = "renewable"
	FieldExplicitMaxTTL           = "explicit_max_ttl"
	FieldWrappingTTL              = "wrapping_ttl"
	FieldDisplayName              = "display_name"
	FieldNumUses                  = "num_uses"
	FieldRenewMinLease            = "renew_min_lease"
	FieldRenewIncrement           = "renew_increment"
	FieldLeaseStarted             = "lease_started"
	FieldClientToken              = "client_token"
	FieldWrappedToken             = "wrapped_token"
	FieldOrphan                   = "orphan"
	FieldVaultVersionOverride     = "vault_version_override"
	FieldSkipGetVaultVersion      = "skip_get_vault_version"
	FieldMemberEntityIDs          = "member_entity_ids"
	FieldMemberGroupIDs           = "member_group_ids"
	FieldExclusive                = "exclusive"
	FieldGroupID                  = "group_id"
	FieldGroupName                = "group_name"
	FieldExternal                 = "external"
	FieldInternal                 = "internal"

	/*
		common environment variables
	*/
	EnvVarVaultNamespaceImport = "TERRAFORM_VAULT_NAMESPACE_IMPORT"
	EnvVarSkipChildToken       = "TERRAFORM_VAULT_SKIP_CHILD_TOKEN"
	// EnvVarUsername to get the username for the userpass auth method
	EnvVarUsername = "TERRAFORM_VAULT_USERNAME"
	// EnvVarPassword to get the password for the userpass auth method
	EnvVarPassword = "TERRAFORM_VAULT_PASSWORD"
	// EnvVarPasswordFile to get the password for the userpass auth method
	EnvVarPasswordFile = "TERRAFORM_VAULT_PASSWORD_FILE"
	// EnvVarGCPAuthJWT to get the signed JWT for the gcp auth method
	EnvVarGCPAuthJWT = "TERRAFORM_VAULT_GCP_AUTH_JWT"
	// EnvVarVaultAuthJWT to login via the Vault jwt engine.
	EnvVarVaultAuthJWT = "TERRAFORM_VAULT_AUTH_JWT"
	// EnvVarAzureAuthJWT to login into Vault's azure auth engine.
	EnvVarAzureAuthJWT = "TERRAFORM_VAULT_AZURE_AUTH_JWT"

	EnvVarGoogleApplicationCreds = "GOOGLE_APPLICATION_CREDENTIALS"

	// EnvVarKrbSPNEGOToken to get the signed JWT for the gcp auth method
	EnvVarKrbSPNEGOToken = "KRB_SPNEGO_TOKEN"
	// EnvVarKRB5Conf path to the krb5.conf file.
	EnvVarKRB5Conf = "KRB5_CONFIG"
	// EnvVarKRBKeytab path the keytab file.
	EnvVarKRBKeytab = "KRB_KEYTAB"

	// EnvVarRadiusUsername for the Radius auth login
	EnvVarRadiusUsername = "RADIUS_USERNAME"
	// EnvVarRadiusPassword for the Radius auth login
	EnvVarRadiusPassword = "RADIUS_PASSWORD"

	/*
		common mount types
	*/
	MountTypeDatabase   = "database"
	MountTypePKI        = "pki"
	MountTypeAWS        = "aws"
	MountTypeKMIP       = "kmip"
	MountTypeRabbitMQ   = "rabbitmq"
	MountTypeNomad      = "nomad"
	MountTypeKubernetes = "kubernetes"
	MountTypeUserpass   = "userpass"
	MountTypeCert       = "cert"
	MountTypeGCP        = "gcp"
	MountTypeKerberos   = "kerberos"
	MountTypeRadius     = "radius"
	MountTypeOCI        = "oci"
	MountTypeOIDC       = "oidc"
	MountTypeJWT        = "jwt"
	MountTypeAzure      = "azure"
	MountTypeGitHub     = "github"
	MountTypeAD         = "ad"
	MountTypeConsul     = "consul"
	MountTypeTerraform  = "terraform"

	/*
		Vault version constants
	*/
	VaultVersion190 = "1.9.0"
	VaultVersion110 = "1.10.0"
	VaultVersion111 = "1.11.0"
	VaultVersion112 = "1.12.0"

	/*
		Vault auth methods
	*/
	AuthMethodAWS      = "aws"
	AuthMethodUserpass = "userpass"
	AuthMethodCert     = "cert"
	AuthMethodGCP      = "gcp"
	AuthMethodKerberos = "kerberos"
	AuthMethodRadius   = "radius"
	AuthMethodOCI      = "oci"
	AuthMethodOIDC     = "oidc"
	AuthMethodJWT      = "jwt"
	AuthMethodAzure    = "azure"

	/*
		misc. path related constants
	*/
	PathDelim      = "/"
	VaultAPIV1Root = "/v1"
)
