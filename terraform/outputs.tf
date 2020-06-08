
output "ingress_host_ip" {
  value = kubernetes_ingress.grpc-cache.load_balancer_ingress.0.ip
  description = "IP mapping of ingress to host"
}
