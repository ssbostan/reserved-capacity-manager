# reserved-capacity-manager

### Reserved Capacity Manager for Kubernetes (Hot-Spare Pods)

This application is designed to optimize the responsiveness of Kubernetes cluster-autoscaler by proactively reserving low-priority resources on worker nodes as "hot spare" capacity. By deploying this application as a Kubernetes Deployment and scaling it proportionally with cluster-proportional-autoscaler, it ensures that each node maintains a certain number of reserved Pods. These low-priority Pods act as placeholders and are evicted when higher-priority user workloads are scheduled.

### Key Features:

  - **Hot Spare Capacity**: The application runs low-priority Pods across the cluster to reserve node resources, ensuring quicker scheduling of user workloads.
  - **Cluster Autoscaling Optimization**: When user applications are deployed and reserved Pods are evicted, the cluster-autoscaler adds new nodes for the evicted Pods, minimizing the perceived scaling time for users.
  - **Proportional Scaling**: The application dynamically adjusts its replica count using the cluster-proportional-autoscaler, based on the size of the cluster, ensuring optimal resource utilization.
  - **Efficient Resource Utilization**: Reserved Pods consume minimal resources but guarantee node availability to improve user workload scheduling.

## How It Works:

  1. The application deploys low-priority Pods across the cluster using a Deployment.

  2. The cluster-proportional-autoscaler ensures the number of reserved Pods scales dynamically based on the cluster's size and configuration.

  3. When a user deploys a normal application:
       - Reserved Pods are evicted to free up capacity.
       - If the cluster lacks sufficient resources for the evicted Pods, the cluster-autoscaler scales up the cluster by adding new nodes.

  4. By having reserved capacity readily available, the perceived cluster-autoscaler scaling time is minimized for end-users.

Copyright &copy; 2024 Saeid Bostandoust <ssbostan@yahoo.com>
