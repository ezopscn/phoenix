import React from 'react';
import CommonPageHeader from '../../../components/page-header/PageHeader.jsx';

// 提示信息
const ClusterNotices = () => {
  return (
    <>
      <ul>
        <li>集群一组用于运行容器化应用的节点计算机，它至少得包含一个控制平面（Control Plane，Master），以及一个或多个节点（Worker）。</li>
        <li>集群控制平面负责维护集群的预期状态（Spec），例如运行哪个应用以及使用哪个容器镜像。节点则负责应用和工作负载的实际运行。</li>
        <li>在日常生产实践中，为了尽可能的避免因为某些人为原因导致集群瘫痪，我们建议您根据不同的运行环境将集群拆分成不同集群。例如：运行测试环境的 TEST 集群，运行生产环境的 PROD 集群等。</li>
      </ul>
    </>
  );
};

const ClusterDashboard = () => {
  const title = '集群 / Cluster'; // 页面标题
  return (
    <>
      <CommonPageHeader title={title} notices={<ClusterNotices />} />
    </>
  );
};

export default ClusterDashboard;
