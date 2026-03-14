## 在k3s中使用cronjob定时重启deploy

> 使用cronjob，直接执行kubectl操作

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: restart-sa
  namespace: business
---
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  name: restart-role
  namespace: business
rules:
  # 只允许对 apps 组下的 deployments 资源进行操作
  - apiGroups: ["apps"]
    resources: ["deployments"]
    verbs: ["get", "list", "patch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: restart-rolebinding
  namespace: business
subjects:
  - kind: ServiceAccount
    name: restart-sa
    namespace: business
roleRef:
  kind: Role
  name: restart-role
  apiGroup: rbac.authorization.k8s.io
---
apiVersion: batch/v1
kind: CronJob
metadata:
  name: restart-my-app-every-3h
  namespace: business
spec:
  # 每 1小时执行一次 (0 分钟, 0/3 小时, 每天, 每月, 每周)
  schedule: "0 */1 * * *"
  concurrencyPolicy: Forbid
  # 保留最近 1 次成功记录，0 次失败记录 (保持整洁)
  successfulJobsHistoryLimit: 1
  failedJobsHistoryLimit: 0
  jobTemplate:
    spec:
      template:
        spec:
          serviceAccountName: restart-sa
          restartPolicy: OnFailure
          containers:
            - name: kubectl
              # 使用轻量级的 kubectl 镜像，K3s 通常基于 ARM64 或 AMD64，此镜像支持多架构
              image: bitnami/kubectl:latest
              command: ["/bin/sh", "-c"]
              # 执行滚动重启命令
              args:
                - |
                  echo "Starting restart of my-app at $(date)"
                  kubectl rollout restart deployment/realsee-open-admin  -n business
                  echo "Waiting for rollout to finish..."
                  kubectl rollout status deployment/realsee-open-admin -n business
                  echo "Restart completed successfully."
```