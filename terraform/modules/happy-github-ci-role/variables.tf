variable "tags" {
  description = "Standard tags. Typically generated by fogg."
  type = object({
    env : string,
    owner : string,
    project : string,
    service : string,
    managedBy : string,
  })
}

variable "dynamodb_table_arn" {
  description = "The ARN of the dynamodb table that the role should have permissions to"
  type        = string
}

variable "gh_actions_role_name" {
  description = "The name of the role to attach happy permissions to."
  type        = string
}

variable "ecrs" {
  description = "The ECRs that the role should have permissions to"
  type = map(object({
    repository_arn : string,
  }))
  default = {}
}

variable "eks_cluster_arn" {
  description = "The ARN of the EKS cluster that the role should have permissions to"
  type        = string
  default     = ""
}

variable "ecs" {
  description = "The ARN and happy app name of the ECS cluster that the role should have permissions to"
  type = object({
    arn            = string
    happy_app_name = string
  })
  default = { arn = "", happy_app_name = "" }
}

variable "aws_account_id" {
  description = "The account ID of the ECR you want to grant access to"
  type        = string
  default     = ""
}