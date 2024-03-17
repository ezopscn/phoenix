import React from 'react';
import ButterflyPageHeader from '../../../components/page-header/PageHeader.jsx';

// 提示信息
const NodeNotices = () => {
  return (
    <>
      <ul>
        <li>节点可以是一个虚拟机，也可以是一台物理机，但是需要注意的是，节点的名称在集群中必须唯一。</li>
        <li>
          Kubernetes 将节点（Node）分为两类：
          <ol>
            <li>部署控制平面组件（Control Plane Components）的 Master 节点，主要运行 API Server，Controller Manager，Scheduler 等集群管理核心组件。</li>
            <li>部署工作负载（Workload）的 Worker（也叫 Node）节点，主要运行容器运行时、Kubelet 以及 Kube-proxy 等系统组件和业务 Pod。</li>
          </ol>
        </li>
      </ul>
    </>
  );
};

const Node = () => {
  const title = '节点 / Node'; // 页面标题

  return (
    <>
      <ButterflyPageHeader title={title} notices={<NodeNotices />} />
    </>
  );
};

export default Node;
