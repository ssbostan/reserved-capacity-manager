# reserved-capacity-manager

### Reserved Capacity Manager for Kubernetes

This application is designed to optimize the responsiveness of Kubernetes cluster-autoscaler by proactively reserving low-priority resources on worker nodes as "hot spare" capacity. By deploying this application as a Kubernetes Deployment and scaling it proportionally with cluster-proportional-autoscaler, it ensures that each node maintains a certain number of reserved Pods. These low-priority Pods act as placeholders and are evicted when higher-priority user workloads are scheduled.

## Key Features:

  - **Hot Spare Capacity**: The application runs low-priority Pods across the cluster to reserve node resources, ensuring quicker scheduling of user workloads.
  - **Cluster Autoscaling Optimization**: When user applications are deployed and reserved Pods are evicted, the cluster-autoscaler adds new nodes for the evicted Pods, minimizing the perceived scaling time for users.
  - **Proportional Scaling**: The application dynamically adjusts its replica count using the cluster-proportional-autoscaler, based on the size of the cluster, ensuring optimal resource utilization.
  - **Efficient Resource Utilization**: Reserved Pods consume minimal resources but guarantee node availability to improve user workload scheduling.

## How It Works:

  1. The application deploys low-priority Pods across the cluster using a Deployment.

  2. The **cluster-proportional-autoscaler** ensures the number of reserved Pods scales dynamically based on the cluster's size and configuration.

  3. When a user deploys a normal application:
       - Reserved Pods are evicted to free up capacity.
       - If the cluster lacks sufficient resources for the evicted Pods, the **cluster-autoscaler** scales up the cluster by adding new nodes.

  4. By having reserved capacity readily available, the perceived cluster-autoscaler scaling time is minimized for end-users.

## Facts used to develop this application:

  - **Kubernetes QoS (Quality of Service) classes**:
      - Kubernetes relies on this classification to make decisions about which Pods to evict when there are not enough available resources on a Node.
      - Kubernetes does this classification based on the resource requests of the Containers in that Pod, along with how those requests relate to resource limits.
      - QoS classes are used by Kubernetes to decide which Pods to evict from a Node experiencing Node Pressure.
      - The possible QoS classes are Guaranteed, Burstable, and BestEffort.
      - When a Node runs out of resources, Kubernetes will first evict BestEffort Pods running on that Node, followed by Burstable and finally Guaranteed Pods.
      - When this eviction is due to resource pressure, only Pods exceeding resource requests are candidates for eviction.
      - If a Container exceeds its resource request and the node it runs on faces resource pressure, the Pod it is in becomes a candidate for eviction.
      - The kube-scheduler does not consider QoS class when selecting which Pods to preempt. Preemption can occur when a cluster does not have enough resources to run all the Pods you defined.

  - **Node-pressure Eviction**:
      - The kubelet does not respect your configured PodDisruptionBudget or the pod's terminationGracePeriodSeconds.
      - The kubelet attempts to reclaim node-level resources before it terminates end-user pods. For example, it removes unused container images when disk resources are starved.
      - The kubelet uses the following parameters to determine the pod eviction order:
          1. Whether the pod's resource usage exceeds requests
          2. Pod Priority
          3. The pod's resource usage relative to requests
      - As a result, kubelet ranks and evicts pods in the following order:
          1. BestEffort or Burstable pods where the usage exceeds requests. These pods are evicted based on their Priority and then by how much their usage level exceeds the request.
          2. Guaranteed pods and Burstable pods where the usage is less than requests are evicted last, based on their Priority.

  - **Kubernetes Priority classes**:
      - If a Pod cannot be scheduled, the scheduler tries to preempt (evict) lower priority Pods to make scheduling of the pending Pod possible.
      - Kubernetes already ships with two PriorityClasses: system-cluster-critical and system-node-critical.
      - The cluster must have at least one global default priority class.
      - Pods with preemptionPolicy: Never will be placed in the scheduling queue ahead of lower-priority pods, but they cannot preempt other pods.
      - The global default priority class must have preemptionPolicy: Never to avoid preempting other pods.
      - The priority class created for one-off jobs like AI and ML jobs can have lower or higher priority but must be non-preempting to avoid preempting normal pods.
      - Non-preempting pods may still be preempted by other, high-priority pods.
      - When Pod priority is enabled, the scheduler orders pending Pods by their priority and a pending Pod is placed ahead of other pending Pods with lower priority in the scheduling queue.
      - When Pods are preempted, the victims get their graceful termination period. They have that much time to finish their work and exit.
      - Kubernetes supports PDB when preempting Pods, but respecting PDB is best effort.
      - The scheduler tries to find victims whose PDB are not violated by preemption, but if no such victims are found, preemption will still happen, and lower priority Pods will be removed despite their PDBs being violated.
      - A Node is considered for preemption only when the answer to this question is yes: "If all the Pods with lower priority than the pending Pod are removed from the Node, can the pending Pod be scheduled on the Node?"
      - If a pending Pod has inter-pod affinity to one or more of the lower-priority Pods on the Node, the inter-Pod affinity rule cannot be satisfied in the absence of those lower-priority Pods. In this case, the scheduler does not preempt any Pods on the Node.
      - Our recommended solution for this problem is to create inter-Pod affinity only towards equal or higher priority Pods.
      - The kubelet uses Priority to determine pod order for node-pressure eviction.
      - kubelet node-pressure eviction does not evict Pods when their usage does not exceed their requests. If a Pod with lower priority is not exceeding its requests, it won't be evicted. Another Pod with higher priority that exceeds its requests may be evicted.

Copyright &copy; 2024 Saeid Bostandoust <ssbostan@yahoo.com>
