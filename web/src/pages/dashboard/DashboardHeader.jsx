import React, { useEffect, useState } from 'react';
import { Avatar, message, Statistic } from 'antd';
import { UserStates } from '../../store/Store.jsx';
import { useSnapshot } from 'valtio';
import { CurrentUserDepartmentInfoRequest, UserCountRequest } from '../../utils/RequestAPI.jsx';

// 问候语
function getHelloWord(name) {
  let arr = ['日', '一', '二', '三', '四', '五', '六'];
  let day = new Date();
  let hour = day.getHours();
  let week = day.getDay();
  let hello = name + '，今天是星期' + arr[week];
  if (hour > 22) {
    hello = hello + '，别卷了，要好好休息哦 ~';
  } else if (hour > 18) {
    hello = hello + '，适当加班，然后早点回家吧 ~';
  } else if (hour > 14) {
    hello = hello + '，如果困了，来杯咖啡提神吧 ~';
  } else if (hour > 11) {
    hello = hello + '，好好吃饭，好好休息，中午不睡下午崩溃哦 ~';
  } else if (hour > 6) {
    hello = hello + '，新的一天，也要元气满满哦 ~';
  } else if (hour > 3) {
    hello = hello + '，也太早了吧，你是还没睡吗 ~';
  } else {
    hello = hello + '，这个时候不是应该睡觉吗 ~';
  }
  return hello;
}

// 组合部门信息
function getDepartmentNames(data) {
  if (!data.children) {
    return data.name;
  }

  const names = [];
  for (const child of data.children) {
    names.push(getDepartmentNames(child));
  }

  return data.name + ' - ' + names.join('-');
}

// 工作台 Header
const DashboardHeader = () => {
  const [departmentNames, setDepartmentNames] = useState('未知');
  const [userCount, setUserCount] = useState(0);
  const { CurrentUserInfo } = useSnapshot(UserStates);

  useEffect(() => {
    // 获取用户部门
    (async () => {
      try {
        const res = await CurrentUserDepartmentInfoRequest();
        if (res.code === 200) {
          setDepartmentNames(getDepartmentNames(res.data.info));
        } else {
          message.error(res.message);
        }
      } catch (e) {
        console.log(e);
        message.error('服务器异常，请联系管理员');
      }
    })();

    // 获取用户总数
    (async () => {
      try {
        const res = await UserCountRequest();
        if (res.code === 200) {
          setUserCount(res.data.count);
        } else {
          message.error(res.message);
        }
      } catch (e) {
        console.log(e);
        message.error('服务器异常，请联系管理员');
      }
    })();
  }, []);

  // 问候语
  let hello = getHelloWord(CurrentUserInfo?.cn_name + '（' + CurrentUserInfo?.en_name + '）');

  // 入职天数计算
  var joinTime = CurrentUserInfo?.join_time;
  const today = new Date();
  const diffTime = today.getTime() - new Date(joinTime).getTime();
  const joinDays = Math.floor(diffTime / (1000 * 60 * 60 * 24)) + 1;

  return (
    <>
      <div className="admin-left">
        <div className="admin-avatar">
          <Avatar src={CurrentUserInfo?.avatar} size={60} />
        </div>
        <div className="admin-info">
          <div className="admin-welcome">{hello}</div>
          <div className="admin-desc">
            {CurrentUserInfo?.job_name} | {departmentNames}
          </div>
        </div>
      </div>
      <div className="admin-right">
        <Statistic
          title="用户数量"
          value={userCount}
          style={{
            marginRight: 30,
          }}
        />
        <Statistic
          title="入职天数"
          value={joinDays}
          style={{
            marginRight: 30,
          }}
        />
        <Statistic title="集群数量" value={7} />
      </div>
    </>
  );
};

export default DashboardHeader;
