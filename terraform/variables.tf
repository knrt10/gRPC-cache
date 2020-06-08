variable "image_name" {
  type    = string
  default = "grpc-cache"
}

variable "image_repository" {
  type    = string
  default = "knrt10/grpc-cache"
}

variable "kube_defaultspace" {
  type    = string
  default = "grpc-cache"
}

variable "kube_deployment_replica" {
  type    = number
  default = 1
}


