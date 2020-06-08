provider "kubernetes" {}

# create namespace
resource "kubernetes_namespace" "grpc-cache" {
  metadata {
    name = var.kube_defaultspace
  }
}

# create deployment
resource "kubernetes_deployment" "grpc-cache" {
  metadata {
    name      = "grpc-cache"
    namespace = var.kube_defaultspace
  }

  spec {
    replicas = var.kube_deployment_replica

    selector {
      match_labels = {
        name = "grpc-cache"
      }
    }

    template {
      metadata {
        labels = {
          name = "grpc-cache"
        }
        namespace = var.kube_defaultspace
      }

      spec {
        container {
          name  = var.image_name
          image = var.image_repository
          port {
            container_port = 5001
            name           = "grpc"
          }
        }
      }
    }
  }
}

# create service

resource "kubernetes_service" "grpc-cache" {
  metadata {
    name      = "grpc-cache"
    namespace = var.kube_defaultspace
  }

  spec {
    selector = {
      name = "grpc-cache"
    }

    port {
      port        = 5001
      target_port = 5001
      name        = "grpc"
    }
  }
}

# create ingress

resource "kubernetes_ingress" "grpc-cache" {
  metadata {
    name      = "grpc-cache"
    namespace = var.kube_defaultspace
    annotations = {
      "kubernetes.io/ingress.class"                  = "nginx"
      "nginx.ingress.kubernetes.io/backend-protocol" = "GRPC"
      "nginx.ingress.kubernetes.io/ssl-redirect"     = "true"
    }
  }

  wait_for_load_balancer = true
  
  spec {
    rule {
      host = "grpc-cache.example.com"
      http {
        path {
          backend {
            service_name = "grpc-cache"
            service_port = 5001
          }
        }
      }
    }
  }
}
