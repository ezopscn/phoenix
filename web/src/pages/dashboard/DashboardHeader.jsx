import React from "react";
import { Avatar, Statistic } from "antd";
import { UserStates } from "../../store/Store.jsx";

// 问候语
function getHelloWord(name) {
  let arr = ["日", "一", "二", "三", "四", "五", "六"];
  let day = new Date();
  let hour = day.getHours();
  let week = day.getDay();
  let hello = name + "，今天是星期" + arr[week];
  if (hour > 22) {
    hello = hello + "，别卷了，要好好休息哦 ~";
  } else if (hour > 18) {
    hello = hello + "，适当加班，然后早点回家吧 ~";
  } else if (hour > 14) {
    hello = hello + "，如果困了，来杯咖啡提神吧 ~";
  } else if (hour > 11) {
    hello = hello + "，好好吃饭，好好休息，中午不睡下午崩溃哦 ~";
  } else if (hour > 6) {
    hello = hello + "，新的一天，也要元气满满哦 ~";
  } else if (hour > 3) {
    hello = hello + "，也太早了吧，你是还没睡吗 ~";
  } else {
    hello = hello + "，这个时候不是应该睡觉吗 ~";
  }
  return hello;
}

// 工作台 Header
const DashboardHeader = () => {
  // 问候语
  let hello = getHelloWord(
    UserStates.CurrentUserInfo?.cn_name +
      "（" +
      UserStates.CurrentUserInfo?.en_name +
      "）",
  );

  return (
    <>
      <div className="admin-left">
        <div className="admin-avatar">
          <Avatar src={UserStates.CurrentUserInfo?.avatar} size={60} />
        </div>
        <div className="admin-info">
          <div className="admin-welcome">{hello}</div>
          <div className="admin-desc">
            高级运维工程师 | 深圳运维集团 － 产品研发中心 － 运维组 － DevOPS
            团队
          </div>
        </div>
      </div>
      <div className="admin-right">
        <Statistic
          title="用户数量"
          value={1024}
          style={{
            marginRight: 30,
          }}
        />
        <Statistic
          title="入职天数"
          value={65535}
          style={{
            marginRight: 30,
          }}
        />
        <Statistic title="任务数量" value={16384} />
      </div>
    </>
  );
};

export default DashboardHeader;
