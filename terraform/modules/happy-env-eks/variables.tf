
variable "tags" {
  description = "Standard tags. Typically generated by fogg"
  type = object({
    env : string,
    owner : string,
    project : string,
    service : string,
    managedBy : string,
  })
}

variable "base_zone_id" {
  description = "The base zone all happy stacks and infrastructure will build on top of"
  type        = string
}

variable "ecr_repos" {
  description = "Map of ECR repositories to create. These should map exactly to the service names of your docker-compose"
  type = map(object({
    name       = string,
    read_arns  = list(string),
    write_arns = list(string),
  }))
  default = {}
}

variable "rds_dbs" {
  description = "Map of DB's to create for your happy applications. If an engine_version is not provided, the default_db_engine_version is used"
  type = map(object({
    name           = string,
    username       = string,
    instance_class = string,
    engine_version = string,
  }))
  default = {}
}

variable "s3_buckets" {
  description = "Map of S3 buckets to create for your happy applications"
  type        = map(object({ name = string }))
  default     = {}
}

variable "additional_secrets" {
  description = "Any extra secret key/value pairs to make available to services"
  type        = any
  default     = {}
}

variable "oauth_dns_prefix" {
  description = "DNS prefix for oauth-proxied stacks. Leave this empty if we don't need a prefix!"
  type        = string
  default     = ""
}

variable "oauth_bypass_paths" {
  description = "Bypass these paths in the oauth proxy"
  type        = list(string)
  default     = []
}

variable "default_db_engine_version" {
  description = "The default Aurora Postgres engine version if one is not specified in rds_dbs"
  type        = string
  default     = "14.3"
}

variable "cloud-env" {
  type = object({
    public_subnets        = list(string)
    private_subnets       = list(string)
    database_subnets      = list(string)
    database_subnet_group = string
    vpc_id                = string
    vpc_cidr_block        = string
  })
}

variable "eks-cluster" {
  type = object({
    cluster_id : string,
    cluster_arn : string,
    cluster_endpoint : string,
    cluster_ca : string,
    cluster_oidc_issuer_url : string,
    cluster_security_group : string,
    cluster_iam_role_name : string,
    cluster_version : string,
    worker_iam_role_name : string,
    kubeconfig : string,
    worker_security_group : string,
    oidc_provider_arn : string,
  })
  description = "eks-cluster module output"
}

variable "k8s-core" {
  description = "K8s core. Typically the outputs of the remote state for the corresponding k8s-core component."
  type = object({
    default_namespace : string,
    aws_ssm_iam_role_name : string,
  })
}

variable "oidc_issuer_host" {
  description = "The OIDC issuer host for the OIDC provider to use for happy authentication"
  type        = string
  default     = "czi-prod.okta.com"
}

variable "authorized_github_repos" {
  description = "List of Github repos that are authorized to assume the created CI role"
  type        = set(string)
  default     = []
}
