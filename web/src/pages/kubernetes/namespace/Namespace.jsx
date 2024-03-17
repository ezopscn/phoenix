import React from 'react';
import CommonPageHeader from '../../../components/page-header/PageHeader.jsx';

// 提示信息
const NamespaceNotices = () => {
  return (
    <>
      <ul>
        <li>名称空间能够将同一集群中的资源划分为相互隔离的组，并能够针对各个组进行资源限制。同一名称空间内的资源名称要求唯一，但跨名称空间时则没有这个要求。</li>
        <li>名称空间作用域仅针对带有名称空间字段的对象（例如 Deployment、Service 等）生效，而对于集群范围的对象（例如 StorageClass、Node、PersistentVolume 等）则不适用。</li>
        <li>在 Kubernetes 中，有两个特殊的名称空间：default（默认名称空间）和 kube-system（集群系统组件名称空间）</li>
      </ul>
    </>
  );
};

const Namespace = () => {
  const title = '名称空间 / Namespace'; // 页面标题
  return (
    <>
      <CommonPageHeader title={title} notices={<NamespaceNotices />} />
    </>
  );
};

export default Namespace;
