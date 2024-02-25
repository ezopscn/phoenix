// 页面头部，接收以下参数
// title: 名称
// notices: 说明列表
const ButterflyPageHeader = (props) => {
  return (
    <div className="admin-common-page-header">
      <div className="admin-notice">
        <div className="admin-title">{props.title}</div>
        <div className="admin-list">{props.notices}</div>
      </div>
    </div>
  );
};

export default ButterflyPageHeader;
