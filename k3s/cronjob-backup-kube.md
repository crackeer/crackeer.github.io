## 在k3s中使用cronjob备份集群配置

> 使用cronjob，定时备份各个namespace下的deploy、service等

```yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kube-backup-reader
rules:
- apiGroups:
  - '*'
  resources:
  - '*'
  verbs: ["get", "list"]
---

kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: kube-backup
  namespace: default
subjects:
- kind: ServiceAccount
  name: kube-backup
  namespace: kube-system
roleRef:
  kind: ClusterRole
  name: kube-backup-reader
  apiGroup: rbac.authorization.k8s.io

---

apiVersion: v1
kind: ServiceAccount
metadata:
  name: kube-backup
  namespace: kube-system

---

apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: kube-state-backup
  namespace: kube-system
  labels:
    app: kube-backup
spec:
  schedule: "* */6 * * *"
  concurrencyPolicy: Replace
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 3
  jobTemplate:
    spec:
      template:
        metadata:
          labels:
            app: kube-backup
          name: kube-backup
        spec:
          containers:
          - image: quay.io/plange/kube-backup:1.12.0-1
            imagePullPolicy: Never
            name: backup
            resources: {}
            securityContext:
              runAsUser: 1000
            env:
            - name: DRY_RUN
              value: "1"
            - name: GIT_PREFIX_PATH 
              value: "local"
            volumeMounts:
            - mountPath: /backup
              name: backup-path

          dnsPolicy: ClusterFirst
          terminationGracePeriodSeconds: 30
          serviceAccountName: kube-backup
          volumes:
          - name: backup-path
            hostPath:
              path: /data1/kube-backup
              type: DirectoryOrCreate
            
          restartPolicy: OnFailure
```